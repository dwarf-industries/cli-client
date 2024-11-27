package interfaces

type SocketConnection interface {
	Connect(url *string, connectionData *map[string]interface{}) bool
	Disconnect() bool
	SendData(data *map[string]interface{}) *map[string]interface{}
}
