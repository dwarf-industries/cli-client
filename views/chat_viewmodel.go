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
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"client/di"
	"client/interfaces"
	"client/models"
)

type ChatViewModel struct {
	messages              []Message
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
	nodeChunks            int
	chunkBytes            [][]byte
	receivedChunks        map[string]map[int]string
	mu                    sync.Mutex
	updateCallback        func(msg tea.Msg) // Callback to send updates to Bubble Tea
}

func (c *ChatViewModel) init() {
	for _, connection := range *c.NodeConnections {
		go c.listenForNodeMessages(connection)
	}
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

	c.mu.Lock()
	c.messages = append(c.messages, Message{
		Sender:  true,
		Content: message,
	})
	c.message = message
	messageSize := getMeasureForDataRequest([]byte(message))
	c.currentTax = c.PaymentProcessor.CalculatePayment(messageSize)

	c.nodeChunks = len([]byte(message)) / len(*c.NodeConnections)
	c.chunkBytes = splitIntoChunks([]byte(c.message), c.nodeChunks)

	c.expectedNodes = 0
	var cIndex int
	fmt.Println(c.NodeConnections)
	for _, connection := range *c.NodeConnections {

		data := map[string]interface{}{
			"action": "pop-request",
			"size":   len(c.chunkBytes[cIndex]),
		}

		if ok := connection.SendData(&data); !ok {
			panic("failed to send request")
		}
		c.expectedNodes += 1
		cIndex++
	}

	cIndex = 0
	c.mu.Unlock()
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

	// Parse the metadata from the chunk
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

	c.mu.Lock()
	messageID := content.MessageID
	chunkIndex := content.ChunkIndex
	totalChunks := content.TotalChunks

	if c.receivedChunks[messageID] == nil {
		c.receivedChunks[messageID] = make(map[int]string)
	}

	c.receivedChunks[messageID][chunkIndex] = string(decryptedData)

	if len(c.receivedChunks[messageID]) == totalChunks {
		chunkOrder := make([]int, totalChunks)
		for i := 0; i < totalChunks; i++ {
			chunkOrder[i] = i
		}

		chunkOrder = di.GetDataOrderService().ReconstructChunkOrder(messageID, string(c.user.OrderSecret), totalChunks)

		var reassembledMessage string
		for _, index := range chunkOrder {
			reassembledMessage += c.receivedChunks[messageID][index]
		}

		decryptedMessage, err := c.CertifciateService.DecryptWithPrivateKey(c.CertificatePrivateKey, []byte(reassembledMessage))
		if err != nil {
			log.Printf("Failed to decrypt full message: %v\n", err)
			return
		}

		c.mu.Lock()
		c.messages = append(c.messages, Message{
			Sender:  false,
			Content: string(decryptedMessage),
		})
		c.mu.Unlock()

		if c.updateCallback != nil {
			c.updateCallback(Message{
				Sender:  false,
				Content: string(decryptedMessage),
			})
		}

		delete(c.receivedChunks, messageID)
	}
	c.mu.Unlock()
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
	c.input = ""

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

	certBytes := hex.EncodeToString(cert.Raw)

	// Generate a unique message ID (could be a timestamp or random UUID)
	messageID := di.GetDataOrderService().GenerateMessageID()

	shuffledNodes := di.GetDataOrderService().ShuffleNodes(string(c.user.OrderSecret), *c.NodeConnections)

	var cIndex int
	for _, name := range shuffledNodes {
		connection := (*c.NodeConnections)[name]
		encryptedMessage, err := c.CertifciateService.EncryptWithCertificate(c.encryptionCertificate, c.chunkBytes[cIndex])
		if err != nil {
			panic("failed to encrypt data")
		}

		paymentId := c.nodePayments[name]
		encoded := hex.EncodeToString([]byte(paymentId))

		data := map[string]interface{}{
			"action":           "store",
			"encryptedMessage": encryptedMessage,
			"messageID":        messageID,
			"for":              certBytes,
			"pop":              encoded,
			"chunkIndex":       cIndex,
			"totalChunks":      len(shuffledNodes),
		}

		if ok := connection.SendData(&data); !ok {
			fmt.Println("Failed to send data")
		}

		delete(c.nodePayments, name)
		cIndex++
	}

	cIndex = 0
}

func splitIntoChunks(data []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}
	return chunks
}

func getMeasureForDataRequest(b []byte) int {
	return len(b)
}
