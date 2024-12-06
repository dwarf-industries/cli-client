package commands

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"client/di"
	"client/interfaces"
	"client/models"
	"client/repositories"
	"client/views"
)

type ConnectCommand struct {
	PasswordManager       interfaces.PasswordManager
	WalletService         interfaces.WalletService
	AuthenticationService interfaces.AuthenticationService
	UsersRepository       repositories.UsersRepository
	SocketService         interfaces.SocketConnection
	NodeRepository        repositories.NodesRepository
	RegisterService       interfaces.RegisterService
	KeysRepository        repositories.KeysRepository
	CertificateService    interfaces.CertificateService
	password              string
}

func (c *ConnectCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "connect [UserID]",
		Short: "Connect to a WebSocket using the provided UserID",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				log.Fatal("User ID is required")
			}
			userId := args[0]

			converted, err := strconv.Atoi(userId)
			if err != nil {
				fmt.Println("Only numbers are allowed!")
				os.Exit(0)
				return
			}
			c.Execute(&converted)
		},
	}
}

func (c *ConnectCommand) Execute(userId *int) {
	password := c.PasswordManager.Input()
	_, err := c.WalletService.GetWallet(password)
	if err != nil {
		panic("Active operation wallet not found")
	}

	c.password = *password
	di.DatabaseService().Open()
	defer di.DatabaseService().Close()

	user, err := c.UsersRepository.GetUserById(*userId)

	if err != nil {
		panic("User doesn't exist")
	}

	selectedNodes, err := c.NodeRepository.Selected()
	if err != nil || len(*selectedNodes) == 0 {
		fmt.Println("There are no node preferences")
		fmt.Println("Warning: Since there are no node preferences you're about to connect to the entire network for sharding")
		fmt.Println("This can lead to expensive network transfers, are you sure you want  are you sure you want to proceed")
		fmt.Println("Enter 'y' to accept or anything else to cancel the action")

		var answer string
		fmt.Scanf(answer)

	}

	nodes, err := c.RegisterService.Oracles()
	if err != nil {
		fmt.Println("Failed to get nodes, aborting!")
		os.Exit(1)
	}

	var currentNodes []models.Node
	for _, n := range *selectedNodes {
		selected := c.selectedNode(n, &nodes)

		if selected == nil {
			continue
		}

		currentNodes = append(currentNodes, *selected)
	}

	//Well if a man wants the full power, we give it, we don't question or ask why.
	if len(currentNodes) == 0 {
		currentNodes = nodes
	}

	establishedConnections := make(map[string]interfaces.SocketConnection)
	for _, n := range currentNodes {
		c.ConnectToNode(&n, &user)
		establishedConnections[n.Name] = c.SocketService
	}

	key, err := c.KeysRepository.GetEncryptionPrivateKey(userId)
	if err != nil {
		panic("can't get decryption key for user")
	}

	keyBytes, err := hex.DecodeString(*key)
	if err != nil {
		panic("failed to decode hex")
	}

	decryptedKey, err := c.PasswordManager.Decrypt(keyBytes, []byte(c.password))
	if err != nil {
		fmt.Printf("Failed to parse RSA private key: %v\n", err)
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(decryptedKey)
	if err != nil {
		fmt.Printf("Failed to parse RSA private key: %v\n", err)
		return
	}

	chat := views.InitChatView(
		&user,
		&establishedConnections,
		di.GetPaymentProcessor(),
		c.CertificateService,
		privateKey,
	)

	p := tea.NewProgram(chat)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func (c *ConnectCommand) ConnectToNode(node *models.Node, user *models.User) {
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
		return
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

	privateKeyBytes, err := c.PasswordManager.Decrypt(privBytes, []byte(c.password))
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

	connected := c.SocketService.Connect(&url, &handshake)

	if !connected {
		panic("Failed to connect to node")
	}

	c.SocketService.SetToken(token)

}

func (c *ConnectCommand) selectedNode(name string, nodes *[]models.Node) *models.Node {
	for _, n := range *nodes {
		if n.Name == name {
			return &n
		}
	}

	return nil
}
