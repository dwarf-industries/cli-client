package interfaces

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type WalletService interface {
	NewWallet() (*ecdsa.PrivateKey, error)
	GetWallet(password *string) (*ecdsa.PrivateKey, error)
	SetWallet(wallet *string, password *string) (ecdsa.PrivateKey, error)
	ActiveWallet() (*ecdsa.PrivateKey, error)
	GetAddressForPrivateKey(key *ecdsa.PrivateKey) string
	SignMessage(message []byte) ([]byte, error)
	VerifySignature(message []byte, signature []byte, expectedAddress string) (bool, error)
	GetBalance(wallet *common.Address) big.Int
	NewTransactor(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error)
}
