package models

type Keys struct {
	Id                    int
	EncryptionCertificate string
	IdentityCertifciate   string
	EncryptionKey         string
	IdenitityPrivateKey   string
	OrderSecret           string
}
