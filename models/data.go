package models

type Data struct {
	ID          string
	Content     []byte
	MessageID   string
	ChunkIndex  int
	TotalChunks int
}
