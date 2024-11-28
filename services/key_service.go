package services

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
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
