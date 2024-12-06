package interfaces

type SocketConnection interface {
	Connect(url *string, connectionData *map[string]interface{}) bool
	SubscribeToChanges() *chan map[string]interface{}
	SetToken(token *string)
	Disconnect() bool
	SendData(data *map[string]interface{}) bool
}
