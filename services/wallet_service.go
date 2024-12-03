package services

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"client/interfaces"
)

type WalletService struct {
	Storage         interfaces.Storage
	PasswordManager interfaces.PasswordManager
	RpcService      interfaces.RpcService
	activeWallet    *ecdsa.PrivateKey
}

func (w *WalletService) NewWallet() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(crypto.S256(), crypto.NewKeccakState())
}

func (w *WalletService) SetWallet(wallet *string, password *string) (ecdsa.PrivateKey, error) {
	w.Storage.Open()
	defer w.Storage.Close()
	ok := w.PasswordManager.Match(password)
	if !ok {
		fmt.Println("Account password setup, please provide the correct password in order to change it")
		fmt.Println("Wrong account password aborting!")
		os.Exit(1)
	}

	configured := w.PasswordManager.SetupPassword(password)

	if !configured {
		return ecdsa.PrivateKey{}, errors.New("failed to setup password, aborting")
	}

	privateKey, err := crypto.HexToECDSA(*wallet)
	if err != nil {
		return ecdsa.PrivateKey{}, err
	}

	byteSlice := []byte(*wallet)
	ciphertext, err := w.PasswordManager.Encrypt(byteSlice, []byte(*password))
	if ciphertext == nil || err != nil {
		return ecdsa.PrivateKey{}, err
	}
	savePassword := `
		UPDATE Accounts
		SET key = $1
	`
	ciphertextHex := hex.EncodeToString(*ciphertext)
	err = w.Storage.Exec(&savePassword, &[]interface{}{
		ciphertextHex,
	})

	if err != nil {
		return ecdsa.PrivateKey{}, fmt.Errorf("failed to write encrypted data to file: %w", err)
	}

	return *privateKey, nil
}

func (w *WalletService) GetWallet(password *string) (*ecdsa.PrivateKey, error) {
	w.Storage.Open()
	defer w.Storage.Close()

	sql := `
		SELECT key FROM Accounts
	`
	row := w.Storage.QuerySingle(&sql, &[]interface{}{})

	var wallet string
	err := row.Scan(&wallet)
	if err != nil {
		return nil, err
	}

	walletBytes, err := hex.DecodeString(wallet)
	if err != nil {
		return nil, err
	}

	privateKeyData, err := w.PasswordManager.Decrypt(walletBytes, []byte(*password))
	if err != nil {
		fmt.Println("failed to decrypt private key")
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(string(privateKeyData))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	w.activeWallet = privateKey
	return privateKey, nil
}

func (w *WalletService) ActiveWallet() (*ecdsa.PrivateKey, error) {
	if w.activeWallet == nil {
		return nil, fmt.Errorf("wallet is not unlocked")
	}

	return w.activeWallet, nil
}

func (w *WalletService) SignMessage(message []byte) ([]byte, error) {
	privateKey, err := w.ActiveWallet()
	if err != nil {
		return nil, fmt.Errorf("failed to access active wallet: %v", err)
	}

	addressHex := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Println("Wallet address:", addressHex)

	hash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %v", err)
	}

	return signature, nil
}

func (w *WalletService) GetAddressForPrivateKey(key *ecdsa.PrivateKey) string {
	addressHex := crypto.PubkeyToAddress(key.PublicKey).Hex()
	return addressHex
}

func (w *WalletService) VerifySignature(message []byte, signature []byte, expectedAddress string) (bool, error) {
	hash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	publicKey, err := crypto.SigToPub(hash.Bytes(), signature[:65])
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %v", err)
	}

	recoveredAddress := crypto.PubkeyToAddress(*publicKey).Hex()
	return recoveredAddress == expectedAddress, nil
}

func (w *WalletService) GetBalance(wallet *common.Address) big.Int {
	client := w.RpcService.GetClient()

	latestBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to get latest block, aborting!")
		return *big.NewInt(0)
	}
	walletBalance, err := client.BalanceAt(context.Background(), *wallet, big.NewInt(int64(latestBlock)))

	if err != nil {
		fmt.Println("Failed to get latest block, aborting!")
		return *big.NewInt(0)
	}

	return *walletBalance
}

func (r *WalletService) NewTransactor(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainID, err := r.RpcService.GetClient().ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	auth.GasLimit = uint64(300000)

	auth.GasPrice, err = r.RpcService.GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		auth.GasPrice = big.NewInt(20000000000)
		fmt.Println("Warning: Failed to suggest gas price, using fallback value")
	}

	block, err := r.RpcService.GetClient().BlockNumber(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to retrive current block: %v", err)
	}
	nonce, err := r.RpcService.GetClient().NonceAt(context.Background(), auth.From, big.NewInt(int64(block)))
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	fmt.Println("Transaction details:")
	fmt.Println("ChainID:", chainID)
	fmt.Println("GasLimit:", auth.GasLimit)
	fmt.Println("GasPrice:", auth.GasPrice)
	fmt.Println("Nonce:", auth.Nonce)
	fmt.Println("Block:", block)

	return auth, nil
}
