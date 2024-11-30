package interfaces

import "client/models"

type RegisterService interface {
	Oracles() ([]models.Node, error)
}
