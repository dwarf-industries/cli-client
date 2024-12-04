package views

import (
	"encoding/hex"
	"fmt"

	"client/interfaces"
	"client/models"
)

type ChatView struct {
	NodeConnections  *map[string]interfaces.SocketConnection
	PaymentProcessor interfaces.PaymentProcessor
	nodePayments     map[string]string
	user             *models.User
}

func (c *ChatView) Init(user *models.User) {
	c.user = user
	c.nodePayments = make(map[string]string)
	fmt.Printf("\n--- Chat Mode ---\n")
	fmt.Printf("Connected nodes: %d\n", len(*c.NodeConnections))
	fmt.Println("Type a message and press Enter to send to all connected nodes.")
	fmt.Println("Type '/exit' to leave chat mode.")

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

		for name, connection := range *c.NodeConnections {
			paymentId := c.nodePayments[name]
			encoded := hex.EncodeToString([]byte(paymentId))
			data := map[string]interface{}{
				"action":           "store",
				"encryptedMessage": []byte(message),
				"for":              user.Identity,
				"pop":              encoded,
			}
			response := *connection.SendData(&data)
			fmt.Println(response)
			delete(c.nodePayments, name)
		}
	}
}

func getMeasureForDataRequest(b []byte) int {
	return len(b)
}
