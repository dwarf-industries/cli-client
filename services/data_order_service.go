package services

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"client/interfaces"
)

type DataOrderService struct {
}

func (d *DataOrderService) GenerateMessageID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (d *DataOrderService) ShuffleNodes(orderSecret string, nodeConnections map[string]interfaces.SocketConnection) []string {
	nodeNames := make([]string, 0, len(nodeConnections))
	for name := range nodeConnections {
		nodeNames = append(nodeNames, name)
	}

	seed := d.HashOrderSecret(orderSecret)
	rng := rand.New(rand.NewSource(seed))

	sort.Slice(nodeNames, func(i, j int) bool {
		return rng.Int63()%2 == 0
	})

	return nodeNames
}

func (d *DataOrderService) HashOrderSecret(orderSecret string) int64 {
	hashValue := int64(0)
	for i := 0; i < len(orderSecret); i++ {
		hashValue = (hashValue * 31) + int64(orderSecret[i])
	}
	return hashValue
}
func (c *DataOrderService) ShuffleChunks(orderSecret string, chunkIndices []int) []int {
	// Use the same seed derived from the orderSecret to shuffle the chunk indices
	seed := c.HashOrderSecret(orderSecret)

	rand.Seed(seed)
	rand.Shuffle(len(chunkIndices), func(i, j int) {
		chunkIndices[i], chunkIndices[j] = chunkIndices[j], chunkIndices[i]
	})

	return chunkIndices
}
