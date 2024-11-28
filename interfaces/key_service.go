package interfaces

import (
	"crypto/ed25519"
	"crypto/rsa"
)

type KeyService interface {
	IssueED25519Key() (*ed25519.PublicKey, *ed25519.PrivateKey, error)
	IssueRSAKey(bits int) (*rsa.PrivateKey, error)
}
