package interfaces

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
)

type CertificateService interface {
	LoadCertificate(certFile string) (*x509.Certificate, error)
	LoadPrivateKey(keyFile string) (*rsa.PrivateKey, error)
	IssueIdentityCertificate(pub ed25519.PublicKey, priv ed25519.PrivateKey) (*string, error)
	IssueEncryptionCertificate(priv *rsa.PrivateKey) (*string, error)
	EncryptWithCertificate(cert *x509.Certificate, plaintext []byte) ([]byte, error)
	DecryptWithPrivateKey(privKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error)
	GenerateOrderSecret(name string) (*string, error)
}
