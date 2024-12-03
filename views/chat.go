package views

import (
	"fmt"

	"client/interfaces"
)

type ChatView struct {
	NodeConnections  *map[string]interfaces.SocketConnection
	PaymentProcessor interfaces.PaymentProcessor
}

func (c *ChatView) Init() {
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
		var utilizedNodes [][32]byte
		messageSize := getMeasureForDataRequest([]byte(message))
		data := map[string]interface{}{
			"action": "pop-request",
			"size":   messageSize,
		}
		for _, connection := range *c.NodeConnections {

			response := *connection.SendData(&data)
			paymentId := response["PaymentID"].([32]byte)
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

		for _, connection := range *c.NodeConnections {
			response := *connection.SendData(&data)
			paymentId := response["PaymentID"].([32]byte)
			amount := response["Amount"]
			utilizedNodes = append(utilizedNodes, paymentId)
			fmt.Println(paymentId, amount)
		}

	}

}

func getMeasureForDataRequest(b []byte) int {
	return len(b)
}
