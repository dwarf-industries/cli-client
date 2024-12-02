package repositories

import (
	"fmt"

	"client/interfaces"
	"client/models"
)

type KeysRepository struct {
	storage interfaces.Storage
}

func (k *KeysRepository) UserKeys(identity *string) (*[]models.Keys, error) {
	sql := `
		SELECT id, data, identity FROM Keys
		WHERE identity = $1
	`

	query, err := k.storage.Query(&sql, &[]interface{}{
		&identity,
	})

	if err != nil {
		fmt.Println("Failed to fetch user keys")
		return nil, err
	}

	var userKeys []models.Keys
	for query.Next() {
		var key models.Keys
		err := query.Scan(&key.Id, &key.Data, &key.Identity)

		if err != nil {
			fmt.Println("Failed to bind data, to model key aborting")
			return nil, err
		}

		userKeys = append(userKeys, key)
	}

	return &userKeys, nil
}

func (k *KeysRepository) GetKeyByIdentity(identity *string) (*string, error) {
	sql := `
		SELECT data FROM Keys
		WHERE identity = $1
	`

	querySingle := k.storage.QuerySingle(&sql, &[]interface{}{
		&identity,
	})

	var key string
	err := querySingle.Scan(&key)
	if err != nil {
		return nil, err
	}

	return &key, nil
}

func (k *KeysRepository) AddKey(identity *string, data *string) bool {
	sql := `
		INSERT INTO Keys (data, identity) VALUES ($1, $2)
	`

	err := k.storage.Exec(&sql, &[]interface{}{
		&identity, &data,
	})

	return err == nil
}

func (k *KeysRepository) DeleteUserKey(id *int) bool {
	sql := `
		DELETE FROM Keys
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
