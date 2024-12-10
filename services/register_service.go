package services

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"client/contracts"
	"client/interfaces"
	"client/models"
	"client/repositories"
)

type RegisterService struct {
	WalletService         interfaces.WalletService
	RpcService            interfaces.RpcService
	ContractAddr          string
	VerificationService   interfaces.IdentityVerificationService
	AuthenticationService interfaces.AuthenticationService
	CertificateService    interfaces.CertificateService
	PasswordManager       interfaces.PasswordManager
	KeysRepository        repositories.KeysRepository
}

func (r *RegisterService) Oracles() ([]models.Node, error) {
	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := contracts.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return []models.Node{}, err
	}

	oracleResult, err := contract.GetOracles(&bind.CallOpts{})
	if err != nil {
		return []models.Node{}, nil
	}

	var oracles []models.Node
	for _, o := range oracleResult {
		verified := r.VerificationService.Verify(o.Ip, o.Name.Hex())
		if !verified {
			continue
		}

		oracles = append(oracles, models.Node{
			Name:       o.Name.Hex(),
			Ip:         o.Ip,
			Port:       o.Port,
			Reputation: *o.Reputation,
		})
	}
	return oracles, nil
}

func (c *RegisterService) ConnectToNode(node *models.Node, user *models.User, password *[]byte) interfaces.SocketConnection {
	socketService := &SocketConnection{}
	wallet, err := c.WalletService.ActiveWallet()
	if err != nil {
		fmt.Println("No wallet set, aborting")
		os.Exit(1)
	}

	identity := c.WalletService.GetAddressForPrivateKey(wallet)
	identityBytes := []byte(identity)
	url := node.Ip
	challenge, err := c.AuthenticationService.Authenticate(url, &identityBytes)

	if err != nil {
		panic("Failed to produce challenge, can't establish link with the node")
	}

	signature, err := c.WalletService.SignMessage(*challenge)
	if err != nil {
		fmt.Printf("Failed to produce a valid signature for the given challenge: %s", *challenge)
		return nil
	}

	token, err := c.AuthenticationService.GenerateSessionToken(&url)
	if err != nil {
		panic("couldn't generate session token")
	}

	keys, err := c.KeysRepository.UserKeys(&user.Id)
	if err != nil {
		panic("user doesn't have keys")
	}

	decodedIdentity, err := hex.DecodeString(keys.IdentityCertifciate)
	if err != nil {
		panic("failed to decode identity")
	}

	cert, err := c.CertificateService.LoadCertificate(&decodedIdentity)
	if err != nil {
		panic("failed to import certificate with decoded data")
	}

	privBytes, err := hex.DecodeString(keys.IdenitityPrivateKey)
	if err != nil {
		panic("failed to decode private key bytes")
	}

	privateKeyBytes, err := c.PasswordManager.Decrypt(privBytes, *password)
	if err != nil {
		panic("failed to decrypt private key")
	}

	key := ed25519.NewKeyFromSeed(privateKeyBytes)
	s, err := key.Sign(rand.Reader, cert.RawTBSCertificate, &ed25519.Options{})
	if err != nil {
		panic("failed to produce signature for certificate")
	}

	cert.PublicKey = key.Public()

	encodeIdentity := hex.EncodeToString(user.Identity)
	signatureHex := hex.EncodeToString(signature)
	identitySignature := hex.EncodeToString(s)
	certEncoded := hex.EncodeToString(cert.Raw)
	handshake := map[string]interface{}{
		"action":          "authenticate",
		"address":         identity,
		"certificate":     encodeIdentity,
		"signedChallenge": signatureHex,
		"sessionToken":    token,
		"me":              certEncoded,
		"signature":       identitySignature,
	}

	connected := socketService.Connect(&url, &handshake, node)

	if !connected {
		fmt.Println("Failed to connect to node")
		return nil
	}

	socketService.SetToken(token)
	return socketService
}
