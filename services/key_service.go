package services

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

type KeyService struct {
}

func (k *KeyService) IssueED25519Key() (*ed25519.PublicKey, *ed25519.PrivateKey, error) {
	priv, pub, err := ed25519.GenerateKey(rand.Reader)
	return &priv, &pub, err
}

func (k *KeyService) IssueRSAKey(bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

func (k *KeyService) GenerateRandomPassword() (*string, error) {
	passwordBytes := make([]byte, 32)

	_, err := rand.Read(passwordBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random password: %v", err)
	}

	password := base64.URLEncoding.EncodeToString(passwordBytes)

	return &password, nil
}
