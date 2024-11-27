package di

import (
	"client/repositories"
)

var users repositories.UsersRepository

func Setup() {
	users := repositories.UsersRepository{}
	users.Setup(DatabaseService())
}

func UsersRepository() repositories.UsersRepository {
	return users
}
