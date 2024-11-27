package repositories

import (
	"encoding/hex"
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

func (u *UsersRepository) GetUserById(userId int) (models.User, error) {
	sql := `
		SELECT * FROM Users
		WHERE id = $1
	`

	query := u.storage.QuerySingle(&sql, &[]interface{}{})

	var user models.User
	err := query.Scan(&user.Id, &user.Certificate, &user.Name, &user.CreatedAt)
	if err != nil {
		fmt.Println("Failed to parse user, aborting!")
		return models.User{}, err
	}

	return user, nil
}

func (u *UsersRepository) AddUser(certificate *[]byte, name *string) bool {
	sql := `
		INSERT INTO Users
		(certificate, name, created_at)
		VALUES ($1, $2, $3)
	`

	err := u.storage.Exec(&sql, &[]interface{}{
		hex.EncodeToString(*certificate),
		name,
		time.Now().UTC(),
	})

	return err == nil
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
