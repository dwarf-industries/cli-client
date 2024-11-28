package models

import "time"

type UserCertificate struct {
	Id              int
	UserId          int
	CertificateData string
	CertificateType int
	CreatedAt       time.Time
}
