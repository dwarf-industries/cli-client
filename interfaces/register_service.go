package interfaces

import (
	"client/models"
)

type RegisterService interface {
	Oracles() ([]models.Node, error)
	ConnectToNode(node *models.Node, user *models.User, password *[]byte) SocketConnection
}
