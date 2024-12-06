package views

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"client/interfaces"
	"client/models"
)

type ChatViewModel struct {
	NodeConnections       *map[string]interfaces.SocketConnection
	PaymentProcessor      interfaces.PaymentProcessor
	CertifciateService    interfaces.CertificateService
	nodePayments          map[string]string
	CertificatePrivateKey *rsa.PrivateKey
	user                  *models.User
	utilizedNodes         []string
	expectedNodes         int
	currentTax            *big.Int
	encryptionCertificate *x509.Certificate
	message               string
	input                 string
}

func (c *ChatViewModel) listenForNodeMessages(connection interfaces.SocketConnection) {
	for {
		select {
		case message := <-*connection.SubscribeToChanges():
			c.handleNodeMessage(connection, message)
		case <-time.After(10 * time.Second):
		}
	}
}

func (c *ChatViewModel) ProcessInput() {
	message := strings.TrimSpace(c.input)
	if message == "/exit" {
		c.input = ""
		return
	}

	c.input = ""
	messageSize := getMeasureForDataRequest([]byte(message))
	c.currentTax = c.PaymentProcessor.CalculatePayment(messageSize)
	fmt.Print(c.currentTax)

	data := map[string]interface{}{
		"action": "pop-request",
		"size":   messageSize,
	}

	c.expectedNodes = 0
	for _, connection := range *c.NodeConnections {
		if ok := connection.SendData(&data); !ok {
			panic("failed to send request")
		}
		c.expectedNodes += 1
	}
}

func (c *ChatViewModel) handleNodeMessage(connection interfaces.SocketConnection, message map[string]interface{}) {
	var currentAction string
	var ok bool
	if currentAction, ok = message["action"].(string); !ok {
		return
	}

	switch currentAction {
	case "message":
		c.processMessage(connection, message)
	case "pop":
		c.popRequest(message)
	}
}

func (c *ChatViewModel) processMessage(connection interfaces.SocketConnection, message map[string]interface{}) {
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
	decryptedString := string(decryptedData)
	fmt.Println(decryptedString)
	connection.SendData(&map[string]interface{}{
		"action": "AC",
	})
}

func (c *ChatViewModel) popRequest(data map[string]interface{}) {
	response := data["data"].(map[string]interface{})
	paymentIdData := response["PaymentID"].(string)
	pay, _ := hex.DecodeString(paymentIdData)
	paymentId := string(pay)
	c.nodePayments[data["identifier"].(string)] = paymentId
	c.utilizedNodes = append(c.utilizedNodes, paymentId)

	paid := c.PaymentProcessor.PayNetworkTax(&c.utilizedNodes, c.currentTax)
	if !paid {
		fmt.Println("Transaction failed")
	}

	c.expectedNodes -= 1

	if c.expectedNodes == 0 {
		c.sendMessage()
	}
}

func (c *ChatViewModel) sendMessage() {
	cert, err := c.CertifciateService.LoadCertificate(&c.user.Identity)
	if err != nil {
		panic("couldn't load certificate")
	}

	encryptedMessage, err := c.CertifciateService.EncryptWithCertificate(c.encryptionCertificate, []byte(c.message))
	if err != nil {
		panic("failed to encrypt data")
	}

	certBytes := hex.EncodeToString(cert.Raw)
	for name, connection := range *c.NodeConnections {
		paymentId := c.nodePayments[name]
		encoded := hex.EncodeToString([]byte(paymentId))

		data := map[string]interface{}{
			"action":           "store",
			"encryptedMessage": encryptedMessage,
			"for":              certBytes,
			"pop":              encoded,
		}

		if ok := connection.SendData(&data); !ok {
			fmt.Println("Failed to send data")
		}

		delete(c.nodePayments, name)
	}
}

func getMeasureForDataRequest(b []byte) int {
	return len(b)
}
