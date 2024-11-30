package repositories

import (
	"fmt"

	"client/interfaces"
	"client/models"
)

type KeysRepository struct {
	storage interfaces.Storage
}

func (k *KeysRepository) UserKeys(userId *int) (*[]models.UserKey, error) {
	sql := `
		SELECT id, user_id, key_data, created_at FROM User_Keys
		WHERE user_id = $1
	`

	query, err := k.storage.Query(&sql, &[]interface{}{
		&userId,
	})

	if err != nil {
		fmt.Println("Failed to fetch user keys")
		return nil, err
	}

	var userKeys []models.UserKey
	for query.Next() {
		var key models.UserKey
		err := query.Scan(&key.Id, &key.UserId, &key.KeyData, &key.CreatedAt)

		if err != nil {
			fmt.Println("Failed to bind data, to model key aborting")
			return nil, err
		}

		userKeys = append(userKeys, key)
	}

	return &userKeys, nil
}

// KeyType 1 Encryption Private Key, 2 Identity Public Key, 3 Identity Private Key
func (k *KeysRepository) AddKey(userId *int, keyType int, data *string) bool {
	sql := `
		INSERT INTO User_Keys (key_data, user_id, key_type) VALUES ($1, $2, $3)
	`

	err := k.storage.Exec(&sql, &[]interface{}{
		&userId, &keyType, &data,
	})

	return err == nil
}

func (k *KeysRepository) DeleteUserKey(id *int) bool {
	sql := `
		DELETE FROM User_keys
		WHERE id = $1
	`

	err := k.storage.Exec(&sql, &[]interface{}{
		&id,
	})

	return err == nil
}

func (k *KeysRepository) Init(storage interfaces.Storage) {
	k.storage = storage
}
