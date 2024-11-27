package services

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type SocketConnection struct {
	connection *websocket.Conn
}

func (s *SocketConnection) Connect(url *string, handshake *map[string]interface{}) bool {
	conn, _, err := websocket.DefaultDialer.Dial(*url, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connected to WebSocket")

	if err := conn.WriteJSON(handshake); err != nil {
		log.Fatalf("Failed to send handshake: %v", err)
	}

	var authResponse map[string]interface{}
	if err := conn.ReadJSON(&authResponse); err != nil {
		log.Fatalf("Failed to read handshake response: %v", err)
	}
	fmt.Printf("Handshake response: %v\n", authResponse)

	if authResponse["State"] != "Authenticated" {
		log.Fatalf("Authentication failed: %v", authResponse)
	}

	return false
}

func (s *SocketConnection) SendData(data *map[string]interface{}) *map[string]interface{} {

	if err := s.connection.WriteJSON(data); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	var response map[string]interface{}
	if err := s.connection.ReadJSON(&response); err != nil {
		log.Fatalf("Failed to read handshake response: %v", err)
	}
	return &response
}

func (s *SocketConnection) Disconnect() bool {
	err := s.connection.Close()
	return err != nil
}
