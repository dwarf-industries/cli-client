package services

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type SocketConnection struct {
	connection *websocket.Conn
	token      *string
	messageCh  chan map[string]interface{}
	url        *string
	handshake  *map[string]interface{}
	mu         sync.Mutex
}

func (s *SocketConnection) Connect(url *string, handshake *map[string]interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	socketUrl := fmt.Sprintf("wss://%s/v1/rlt/ws", strings.TrimPrefix(*url, "https://"))
	headers := http.Header{}

	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, headers)
	if err != nil {
		log.Printf("Failed to connect to WebSocket: %v", err)
		return false
	}

	if err := conn.WriteJSON(handshake); err != nil {
		log.Printf("Failed to send handshake: %v", err)
		return false
	}

	var authResponse map[string]interface{}
	if err := conn.ReadJSON(&authResponse); err != nil {
		log.Printf("Failed to read handshake response: %v", err)
		return false
	}

	if authResponse["State"] != "Authenticated" {
		log.Printf("Authentication failed: %v", authResponse)
		return false
	}

	s.connection = conn
	s.messageCh = make(chan map[string]interface{})
	s.url = url
	s.handshake = handshake

	go s.handlePing()
	go s.listenForMessages()

	return true
}

func (s *SocketConnection) reconnect() {
	for {
		log.Println("Attempting to reconnect...")
		if s.Connect(s.url, s.handshake) {
			log.Println("Reconnected successfully")
			return
		}
		time.Sleep(5 * time.Second)
	}
}

func (s *SocketConnection) handlePing() {
	for {
		s.mu.Lock()
		err := s.connection.WriteMessage(websocket.PingMessage, nil)
		s.mu.Unlock()
		if err != nil {
			log.Printf("Ping failed, WebSocket might be disconnected: %v", err)
			close(s.messageCh)
			go s.reconnect()
			break
		}
		time.Sleep(30 * time.Second)
	}
}

func (s *SocketConnection) listenForMessages() {
	for {
		var message map[string]interface{}
		err := s.connection.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Printf("WebSocket closed unexpectedly: %v", err)
			} else {
				log.Printf("ReadJSON error: %v", err)
			}
			close(s.messageCh)
			go s.reconnect()
			return
		}
		s.messageCh <- message
	}
}

func (s *SocketConnection) SubscribeToChanges() *chan map[string]interface{} {
	return &s.messageCh
}

func (s *SocketConnection) SetToken(token *string) {
	s.token = token
}

func (s *SocketConnection) SendData(data *map[string]interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	requestData := *data
	requestData["sessionToken"] = *s.token
	if err := s.connection.WriteJSON(data); err != nil {
		fmt.Printf("Failed to send message: %v", err)
		return false
	}

	return true
}

func (s *SocketConnection) Disconnect() bool {
	err := s.connection.Close()
	return err != nil
}
