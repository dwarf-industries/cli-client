package services

import (
	"crypto/sha256"
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

func (s *DataOrderService) ReconstructChunkOrder(messageID, orderSecret string, chunkCount int) []int {
	seed := s.hashMessageIDAndSecret(messageID, orderSecret)
	rng := rand.New(rand.NewSource(seed))

	chunkOrder := make([]int, chunkCount)
	for i := 0; i < len(chunkOrder); i++ {
		chunkOrder[i] = i
	}

	for i := len(chunkOrder) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		chunkOrder[i], chunkOrder[j] = chunkOrder[j], chunkOrder[i]
	}

	return chunkOrder
}

func (s *DataOrderService) hashMessageIDAndSecret(messageID, orderSecret string) int64 {
	combined := messageID + orderSecret

	hash := sha256.New()
	hash.Write([]byte(combined))
	hashBytes := hash.Sum(nil)

	seed := int64(0)
	for i := 0; i < 8 && i < len(hashBytes); i++ {
		seed = (seed << 8) | int64(hashBytes[i])
	}
	return seed
}
