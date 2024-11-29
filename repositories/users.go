package repositories

import (
	"fmt"
	"time"

	"client/interfaces"
	"client/models"
)

type UsersRepository struct {
	storage interfaces.Storage
}

func (u *UsersRepository) GetAllUsers() ([]models.User, error) {
	sql := `
		SELECT * FROM Users
	`

	query, err := u.storage.Query(&sql, &[]interface{}{})

	if err != nil {
		fmt.Println("failed to fetch users")
		return []models.User{}, err
	}

	var users []models.User
	for query.Next() {
		var user models.User
		err := query.Scan(&user.Id, &user.Certificate, &user.Name, &user.CreatedAt)
		if err != nil {
			fmt.Println("Failed to parse user, aborting!")
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UsersRepository) GetUserByName(name *string) (models.User, error) {
	sql := `
		SELECT * FROM Users
		WHERE name = $1
	`

	query := u.storage.QuerySingle(&sql, &[]interface{}{
		&name,
	})

	var user models.User
	err := query.Scan(&user.Id, &user.Certificate, &user.Name, &user.CreatedAt)
	if err != nil {
		fmt.Println("Failed to parse user, aborting!")
		return models.User{}, err
	}

	return user, nil

}

func (u *UsersRepository) GetUserById(userId int) (models.User, error) {
	sql := `
		SELECT * FROM Users
		WHERE id = $1
	`

	query := u.storage.QuerySingle(&sql, &[]interface{}{
		&userId,
	})

	var user models.User
	err := query.Scan(&user.Id, &user.Certificate, &user.Name, &user.CreatedAt)
	if err != nil {
		fmt.Println("Failed to parse user, aborting!")
		return models.User{}, err
	}

	return user, nil
}

func (u *UsersRepository) AddUser(name *string) (int, error) {
	sql := `
		INSERT INTO Users
		(name, created_at)
		VALUES ($1, $2)
	`

	return u.storage.ExecReturnID(&sql, &[]interface{}{
		name,
		time.Now().UTC(),
	})

}

func (u *UsersRepository) DeleteUser(id int) bool {
	sql := `
		DELETE FROM Users WHERE id = $1
	`

	err := u.storage.Exec(&sql, &[]interface{}{
		id,
	})

	return err == nil
}

func (u *UsersRepository) Setup(s interfaces.Storage) {
	u.storage = s
}
