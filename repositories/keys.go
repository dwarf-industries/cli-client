package repositories

import (
	"fmt"

	"client/interfaces"
	"client/models"
)

type KeysRepository struct {
	storage interfaces.Storage
}

func (k *KeysRepository) UserKeys(userId *int) (*models.Keys, error) {
	sql := `
		SELECT id,identity_certificate, encryption_certificate, encryption_key, priv, order_sercret FROM Keys
		WHERE user_id = $1
	`

	query := k.storage.QuerySingle(&sql, &[]interface{}{
		&userId,
	})

	var key models.Keys
	err := query.Scan(
		&key.Id,
		&key.IdentityCertifciate,
		&key.EncryptionCertificate,
		&key.EncryptionKey,
		&key.IdenitityPrivateKey,
		&key.OrderSecret,
	)

	if err != nil {
		fmt.Println("Failed to bind data, to model key aborting")
		return nil, err
	}

	return &key, nil
}

func (k *KeysRepository) GetKeyByIdentity(id *int) (*string, error) {
	sql := `
		SELECT encryption_certificate FROM Keys
		WHERE user_id = $1
	`

	querySingle := k.storage.QuerySingle(&sql, &[]interface{}{
		&id,
	})

	var key string
	err := querySingle.Scan(&key)
	if err != nil {
		return nil, err
	}

	return &key, nil
}

func (k *KeysRepository) GetEncryptionPrivateKey(id *int) (*string, error) {
	sql := `
		SELECT encryption_key FROM Keys
		WHERE user_id = $1
	`

	querySingle := k.storage.QuerySingle(&sql, &[]interface{}{
		&id,
	})

	var key string
	err := querySingle.Scan(&key)
	if err != nil {
		return nil, err
	}

	return &key, nil
}

func (k *KeysRepository) AddKey(identityCertificate *string, encryptionCertificate *string, encKey *string, priv *string, orderSecret *string, userId *int) bool {
	sql := `
	INSERT INTO Keys (identity_certificate, encryption_certificate, encryption_key, priv, order_sercret, user_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT(user_id) DO UPDATE SET
		identity_certificate = excluded.identity_certificate,
		encryption_certificate = excluded.encryption_certificate,
		encryption_key = excluded.encryption_key,
		priv = excluded.priv,
		order_sercret = excluded.order_sercret;
	`

	err := k.storage.Exec(&sql, &[]interface{}{
		&identityCertificate,
		&encryptionCertificate,
		&encKey,
		&priv,
		&orderSecret,
		&userId,
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
