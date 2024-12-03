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
		var identityEnc string
		var certificateEnc string
		var orderSecret string
		var createdTime string
		var user models.User
		err := query.Scan(&user.Id, &user.Name, &identityEnc, &certificateEnc, &orderSecret, &createdTime)
		if err != nil {
			fmt.Println("Failed to parse user, aborting!")
			return []models.User{}, err
		}
		user.Certificate, _ = hex.DecodeString(certificateEnc)
		user.Identity, _ = hex.DecodeString(identityEnc)
		user.CreatedAt, _ = time.Parse("", createdTime)
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

	var identityEnc string
	var certificateEnc string
	var orderSecret string
	var user models.User
	err := query.Scan(&user.Id, &user.Name, &identityEnc, &certificateEnc, &orderSecret, &user.CreatedAt)

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
	var identityEnc string
	var certificateEnc string
	var orderSecret string
	var createdTime string
	var user models.User
	err := query.Scan(&user.Id, &user.Name, &identityEnc, &certificateEnc, &orderSecret, &createdTime)
	if err != nil {
		fmt.Println("Failed to parse user, aborting!")
		return models.User{}, err
	}

	user.Certificate, _ = hex.DecodeString(certificateEnc)
	user.Identity, _ = hex.DecodeString(identityEnc)
	user.CreatedAt, _ = time.Parse("", createdTime)

	return user, nil
}

func (u *UsersRepository) AddUser(name *string, identity *string, certificate *string, orderSecret *string) (int, error) {
	sql := `
		INSERT INTO Users
		(name, identity,encryptionCertificate,orderSecret, created_at)
		VALUES ($1,$2,$3,$4,$5)
	`

	return u.storage.ExecReturnID(&sql, &[]interface{}{
		&name,
		&identity,
		&certificate,
		&orderSecret,
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
