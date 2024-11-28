package repositories

import (
	"fmt"

	"client/interfaces"
	"client/models"
)

type Certificate struct {
	storage interfaces.Storage
}

func (c *Certificate) UserCertificates(userId *int) *[]models.UserCertificate {
	sql := `
		SELECT id, user_id, certificate_type, certificate_data from User_Certificates
		WHERE user_id = $1
	`

	query, err := c.storage.Query(&sql, &[]interface{}{
		&userId,
	})

	if err != nil {
		return nil
	}

	var certificates []models.UserCertificate
	for query.Next() {
		var certificate models.UserCertificate
		err := query.Scan(
			&certificate.Id,
			&certificate.UserId,
			&certificate.CertificateType,
			&certificate.CertificateData,
			&certificate.CreatedAt,
		)

		if err != nil {
			fmt.Println("failed to parse certificate data, aborting!")
			return nil
		}

		certificates = append(certificates, certificate)
	}

	return &certificates
}

func (c *Certificate) AddCertificate(userId *int, certificateType *int, data *string) bool {
	sql := `
		INSERT INTO User_Certificates (user_id,certificate_type,certificate_data)
		VALUES ($1,$2,$3)
	`

	err := c.storage.Exec(&sql, &[]interface{}{
		&userId,
		&certificateType,
		&data,
	})

	return err == nil
}

func (c *Certificate) DeleteUserCertificate(id *int) bool {
	sql := `
		DEELTE FROM User_Certificates
		WHERE id = $1
	`

	err := c.storage.Exec(&sql, &[]interface{}{
		&id,
	})

	return err == nil
}

func (c *Certificate) Init(storage interfaces.Storage) {
	c.storage = storage
}
