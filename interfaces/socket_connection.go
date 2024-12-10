package interfaces

import "client/models"

type SocketConnection interface {
	Connect(url *string, connectionData *map[string]interface{}, password *models.Node) bool
	Get(url *string) *models.Node
	SubscribeToChanges() *chan map[string]interface{}
	SetToken(token *string)
	Disconnect() bool
	SendData(data *map[string]interface{}) bool
}
