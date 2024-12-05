package views

import (
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"client/interfaces"
	"client/models"
)

type ChatView struct {
	NodeConnections       *map[string]interfaces.SocketConnection
	PaymentProcessor      interfaces.PaymentProcessor
	CertifciateService    interfaces.CertificateService
	nodePayments          map[string]string
	user                  *models.User
	CertificatePrivateKey *rsa.PrivateKey
}

func (c *ChatView) Init(user *models.User) {
	c.user = user
	c.nodePayments = make(map[string]string)
	fmt.Printf("\n--- Chat Mode ---\n")
	fmt.Printf("Connected nodes: %d\n", len(*c.NodeConnections))
	fmt.Println("Type a message and press Enter to send to all connected nodes.")
	fmt.Println("Type '/exit' to leave chat mode.")

	for _, connection := range *c.NodeConnections {
		go c.listenForNodeMessages(connection)
	}

	for {
		fmt.Print("Message: ")
		var message string
		_, err := fmt.Scanln(&message)
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		if message == "/exit" {
			fmt.Println("Exiting chat mode.")
			break
		}
		var utilizedNodes []string
		messageSize := getMeasureForDataRequest([]byte(message))
		data := map[string]interface{}{
			"action": "pop-request",
			"size":   messageSize,
		}
		for name, connection := range *c.NodeConnections {
			response := *connection.SendData(&data)
			paymentIdData := response["PaymentID"].(string)
			fmt.Println(paymentIdData)
			pay, _ := hex.DecodeString(paymentIdData)
			fmt.Println(len(pay))
			paymentId := string(pay)
			c.nodePayments[name] = paymentId
			amount := response["Amount"]

			utilizedNodes = append(utilizedNodes, paymentId)
			fmt.Println(paymentId, amount)
		}

		tax := c.PaymentProcessor.CalculatePayment(messageSize)
		paid := c.PaymentProcessor.PayNetworkTax(&utilizedNodes, tax)
		if !paid {
			fmt.Println("Payment failed, message has been declined")
			continue
		}

		cert, err := c.CertifciateService.LoadCertificate(&user.Identity)
		if err != nil {
			panic("couldn't load certificate")
		}

		certBytes := hex.EncodeToString(cert.Raw)
		for name, connection := range *c.NodeConnections {
			paymentId := c.nodePayments[name]
			encoded := hex.EncodeToString([]byte(paymentId))

			data := map[string]interface{}{
				"action":           "store",
				"encryptedMessage": []byte(message),
				"for":              certBytes,
				"pop":              encoded,
			}

			response := *connection.SendData(&data)
			fmt.Println(response)
			delete(c.nodePayments, name)
		}
	}
}

func (c *ChatView) listenForNodeMessages(connection interfaces.SocketConnection) {
	for {
		select {
		case message := <-*connection.SubscribeToChanges():
			c.handleNodeMessage(connection, message)
		case <-time.After(10 * time.Second):
			fmt.Println("Timeout waiting for message from node")
		}
	}
}

func (c *ChatView) handleNodeMessage(connection interfaces.SocketConnection, message map[string]interface{}) {
	fmt.Printf("Received message from node: %v\n", message)

	if action, ok := message["action"]; ok && action == "message" {
		data, ok := message["data"].(map[string]interface{})
		if !ok {
			log.Println("Invalid data format: expected a map")
			return
		}

		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Printf("Failed to marshal data: %v\n", err)
			return
		}

		var content models.Data
		err = json.Unmarshal(dataBytes, &content)
		if err != nil {
			log.Printf("Failed to decode content: %v\n", err)
			return
		}

		decryptedData, err := c.CertifciateService.DecryptWithPrivateKey(c.CertificatePrivateKey, []byte(content.Content))
		if err != nil {
			panic("failed to decrypt data aborting")
		}

		fmt.Println(string(decryptedData))
		connection.SendData(&map[string]interface{}{
			"action": "AC",
		})
	}
}

func getMeasureForDataRequest(b []byte) int {
	return len(b)
}
