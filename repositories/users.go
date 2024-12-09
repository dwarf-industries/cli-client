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
		var identityEnc *string
		var certificateEnc *string
		var orderSecret *string
		var createdTime string
		var user models.User
		err := query.Scan(&user.Id, &user.Name, &identityEnc, &certificateEnc, &orderSecret, &createdTime)
		if err != nil {
			fmt.Println("Failed to parse user, aborting!")
			return []models.User{}, err
		}

		if certificateEnc != nil {
			user.Certificate, _ = hex.DecodeString(*certificateEnc)
		}

		if identityEnc != nil {
			user.Identity, _ = hex.DecodeString(*identityEnc)

		}
		if orderSecret != nil {
			user.OrderSecret, _ = hex.DecodeString(*orderSecret)
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
	var identityEnc *string
	var certificateEnc *string
	var orderSecret *string
	var createdTime string
	var user models.User
	err := query.Scan(&user.Id, &user.Name, &identityEnc, &certificateEnc, &orderSecret, &createdTime)
	if err != nil {
		fmt.Println("Failed to parse user, aborting!")
		return models.User{}, err
	}

	if certificateEnc != nil {
		user.Certificate, _ = hex.DecodeString(*certificateEnc)
	}

	if identityEnc != nil {
		user.Identity, _ = hex.DecodeString(*identityEnc)

	}
	if orderSecret != nil {
		user.OrderSecret = []byte(*orderSecret)
	}

	user.CreatedAt, _ = time.Parse("", createdTime)

	return user, nil
}

func (u *UsersRepository) CreateUser(name *string) (int, error) {
	sql := `
		INSERT INTO Users
		(name)
		VALUES ($1)
	`

	return u.storage.ExecReturnID(&sql, &[]interface{}{
		&name,
	})
}

func (u *UsersRepository) UpdateUser(id *int, identity *string, certificate *string, orderSecret *string) (int, error) {
	user, err := u.GetUserById(*id)

	if err != nil {
		fmt.Println("User doesn't exist")
		return 0, err
	}

	fmt.Println(user.Id, user.Name)

	sql := `
		UPDATE Users SET
			identity_contract=$1,
			encryption_certificate=$2,
			order_secret=$3
		WHERE id=$4
	`

	err = u.storage.Exec(&sql, &[]interface{}{
		identity,
		certificate,
		orderSecret,
		id,
	})

	return 0, err
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
