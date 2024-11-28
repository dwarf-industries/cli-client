package interfaces

import (
	"crypto/rsa"
	"crypto/x509"
)

type CertificateService interface {
	LoadCertificate(certFile string) (*x509.Certificate, error)
	LoadPrivateKey(keyFile string) (*rsa.PrivateKey, error)
	IssueCertificate(certFile, keyFile string, isEncryption bool) (*string, error)
	EncryptWithCertificate(cert *x509.Certificate, plaintext []byte) ([]byte, error)
	DecryptWithPrivateKey(privKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error)
}
