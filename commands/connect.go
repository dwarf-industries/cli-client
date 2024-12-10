package commands

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
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

		in := `# There are no node preferences

		Warning: Since there are no node preferences you're about to connect to the entire network for sharding
		This can lead to expensive network transfers, are you sure you want  are you sure you want to proceed

		Enter 'y' to accept or anything else to cancel the action
		`

		r, _ := glamour.NewTermRenderer( // detect background color and pick either the default dark or light theme
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(120))

		out, _ := r.Render(in)
		fmt.Print(out)

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
	passBytes := []byte(*password)
	for _, n := range currentNodes {
		socket := c.RegisterService.ConnectToNode(&n, &user, &passBytes)
		establishedConnections[n.Name] = socket
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

func (c *ConnectCommand) selectedNode(name string, nodes *[]models.Node) *models.Node {
	for _, n := range *nodes {
		if n.Name == name {
			return &n
		}
	}

	return nil
}
