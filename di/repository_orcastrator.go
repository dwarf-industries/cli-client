package di

import (
	"client/repositories"
)

var users repositories.UsersRepository
var userKeys repositories.KeysRepository
var userCertificates repositories.Certificate
var nodesRepository repositories.NodesRepository

func Setup() {
	users = repositories.UsersRepository{}
	users.Setup(DatabaseService())
	userKeys = repositories.KeysRepository{}
	userKeys.Init(DatabaseService())
	userCertificates = repositories.Certificate{}
	userCertificates.Init(DatabaseService())
	nodesRepository = repositories.NodesRepository{}
	nodesRepository.Init(DatabaseService())
}

func UsersRepository() repositories.UsersRepository {
	return users
}

func GetUserKeysRepository() repositories.KeysRepository {
	return userKeys
}

func GetUserCertificates() repositories.Certificate {
	return userCertificates
}

func GetNodesRepository() repositories.NodesRepository {
	return nodesRepository
}
