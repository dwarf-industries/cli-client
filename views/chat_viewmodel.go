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
	updateCallback        func(msg tea.Msg)
	password              *[]byte
}

func (c *ChatViewModel) init() {
	c.receivedChunks = make(map[string]map[int]string)
	for _, connection := range *c.NodeConnections {
		go c.listenForNodeMessages(connection)
	}
}

func (c *ChatViewModel) listenForNodeMessages(connection interfaces.SocketConnection) {
	for {
		select {
		case message := <-*connection.SubscribeToChanges():
			c.handleNodeMessage(message)
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

	messageSize := len([]byte(message))

	numNodes := len(*c.NodeConnections)
	if numNodes == 0 {
		panic("no nodes available")
	}

	baseChunkSize := messageSize / numNodes

	remainder := messageSize % numNodes

	var chunks [][]byte
	startIndex := 0
	for i := 0; i < numNodes; i++ {
		chunkSize := baseChunkSize
		if i < remainder {
			chunkSize++
		}

		endIndex := startIndex + chunkSize
		chunks = append(chunks, []byte(message)[startIndex:endIndex])
		startIndex = endIndex
	}

	c.chunkBytes = chunks
	c.expectedNodes = numNodes

	c.nodeChunks = numNodes
	c.currentTax = c.PaymentProcessor.CalculatePayment(messageSize)
	var cIndex int
	for _, connection := range *c.NodeConnections {
		data := map[string]interface{}{
			"action": "pop-request",
			"size":   len(c.chunkBytes[cIndex]),
		}

		if ok := connection.SendData(&data); !ok {
			panic("failed to send request")
		}
		cIndex++
	}

	c.mu.Unlock()
}

func (c *ChatViewModel) handleNodeMessage(message map[string]interface{}) {
	var currentAction string
	var ok bool
	if currentAction, ok = message["action"].(string); !ok {
		return
	}

	switch currentAction {
	case "disconnected":
		c.reconnect(message)
	case "message":
		c.processMessage(message)
	case "pop":
		c.popRequest(message)
	}
}

func (c *ChatViewModel) reconnect(message map[string]interface{}) {
	c.mu.Lock()
	data, ok := message["node"].(string)
	if !ok {
		log.Println("Node URL missing: expected valid URL string")
		return
	}

	disconnectedNodeName, disconnectedNode := c.getSelectedNode(&data)
	nodeConnections := *c.NodeConnections
	_, ok = nodeConnections[*disconnectedNodeName]
	if !ok {
		log.Println("Node not found in connections")
		return
	}
	c.mu.Lock()
	newConnection := di.GetRegisterService().ConnectToNode(disconnectedNode, c.user, c.password)
	nodeConnections[*disconnectedNodeName] = newConnection
	go newConnection.SubscribeToChanges()
	c.NodeConnections = &nodeConnections
	c.mu.Unlock()
}

func (c *ChatViewModel) getSelectedNode(url *string) (*string, *models.Node) {
	var selectedNode string
	var node models.Node
	for name, n := range *c.NodeConnections {
		disconnectedNode := n.Get(url)
		if disconnectedNode == nil {
			continue
		}
		selectedNode = name
		node = *disconnectedNode
	}

	return &selectedNode, &node
}

func (c *ChatViewModel) processMessage(message map[string]interface{}) {
	c.mu.Lock()
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

		var reassembledMessage string
		for _, index := range chunkOrder {
			reassembledMessage += c.receivedChunks[messageID][index]
		}

		c.messages = append(c.messages, Message{
			Sender:  false,
			Content: reassembledMessage,
		})

		if c.updateCallback != nil {
			c.updateCallback(Message{
				Sender:  false,
				Content: reassembledMessage,
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
	c.expectedNodes -= 1
	if c.expectedNodes == 0 {

		paid := c.PaymentProcessor.PayNetworkTax(&c.utilizedNodes, c.currentTax)
		if !paid {
			fmt.Println("Transaction failed")
		}
		c.input = ""

		c.sendMessage()
	}
}

func (c *ChatViewModel) sendMessage() {
	cert, err := c.CertifciateService.LoadCertificate(&c.user.Identity)
	if err != nil {
		panic("couldn't load certificate")
	}

	certBytes := hex.EncodeToString(cert.Raw)

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
