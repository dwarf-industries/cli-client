package interfaces

type DataOrderService interface {
	GenerateMessageID() string
	ShuffleNodes(orderSecret string, nodeConnections map[string]SocketConnection) []string
	HashOrderSecret(orderSecret string) int64
	ReconstructChunkOrder(messageID, orderSecret string, chunkCount int) []int
}
