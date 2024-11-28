package di

import (
	"client/repositories"
)

var users repositories.UsersRepository
var userKeys repositories.KeysRepository

func Setup() {
	users := repositories.UsersRepository{}
	users.Setup(DatabaseService())
	userKeys := repositories.KeysRepository{}
	userKeys.Init(DatabaseService())

}

func UsersRepository() repositories.UsersRepository {
	return users
}

func GetUserKeysRepository() repositories.KeysRepository {
	return userKeys
}
