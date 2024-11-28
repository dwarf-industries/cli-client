package services

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

type CertificateService struct {
}

func (c *CertificateService) LoadCertificate(certFile string) (*x509.Certificate, error) {
	certPEM, err := os.ReadFile(certFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %v", err)
	}
	block, _ := pem.Decode(certPEM)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, errors.New("failed to decode PEM block containing certificate")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %v", err)
	}
	return cert, nil
}

func (c *CertificateService) LoadPrivateKey(keyFile string) (*rsa.PrivateKey, error) {
	keyPEM, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %v", err)
	}
	block, _ := pem.Decode(keyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}
	return privKey, nil
}

func (c *CertificateService) IssueIdentityCertificate(priv *ed25519.PublicKey, pub *ed25519.PrivateKey) (*string, error) {

	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject: pkix.Name{
			Organization:  []string{"Shadowkeep"},
			Country:       []string{"II"},
			Province:      []string{"World Wide"},
			Locality:      []string{"Who knows"},
			StreetAddress: []string{"Easy Street"},
			PostalCode:    []string{"1"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, pub, priv)
	if err != nil {
		log.Fatalf("Failed to create identity certificate: %v", err)
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	log.Printf("Identity certificate and key saved to ")
	pemHex := hex.EncodeToString(certPEM)
	return &pemHex, nil
}

func (c *CertificateService) IssueEncryptionCertificate(priv *rsa.PrivateKey) (*string, error) {
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject: pkix.Name{
			Organization:  []string{"Shadowkeep"},
			Country:       []string{"II"},
			Province:      []string{"World Wide"},
			Locality:      []string{"Who knows"},
			StreetAddress: []string{"Easy Street"},
			PostalCode:    []string{"1"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		BasicConstraintsValid: true,
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		log.Fatalf("Failed to create encryption certificate: %v", err)
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	log.Printf("Encryption certificate and key saved to")
	pemHex := hex.EncodeToString(certPEM)
	return &pemHex, nil
}

func (c *CertificateService) EncryptWithCertificate(cert *x509.Certificate, plaintext []byte) ([]byte, error) {
	pubKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("certificate does not contain an RSA public key")
	}

	hash := crypto.SHA256

	ciphertext, err := rsa.EncryptOAEP(hash.New(), rand.Reader, pubKey, plaintext, nil)
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %v", err)
	}
	return ciphertext, nil
}

func (c *CertificateService) DecryptWithPrivateKey(privKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	hash := crypto.SHA256

	plaintext, err := rsa.DecryptOAEP(hash.New(), rand.Reader, privKey, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %v", err)
	}
	return plaintext, nil
}

func saveToFile(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filename, err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to file %s: %v", filename, err)
	}
	log.Printf("Saved %s", filename)
}
