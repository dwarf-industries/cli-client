package commands

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"client/interfaces"
	"client/models"
	"client/repositories"
)

type ConnectCommand struct {
	WalletService         interfaces.WalletService
	AuthenticationService interfaces.AuthenticationService
	UsersRepository       repositories.UsersRepository
	SocketService         interfaces.SocketConnection
	NodeRepository        repositories.NodesRepository
	RegisterService       interfaces.RegisterService
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
	_, err := c.WalletService.ActiveWallet()
	if err != nil {
		panic("Active operation wallet not found")
	}

	user, err := c.UsersRepository.GetUserById(*userId)

	if err != nil {
		panic("User doesn't exist")
	}

	selectedNodes, err := c.NodeRepository.Selected()
	if err != nil {
		fmt.Println("There are no node preferences")
		fmt.Println("Warning: Since there are no node preferences you're about to connect to the entire network for sharding")
		fmt.Println("This can lead to expensive network transfers, are you sure you want  are you sure you want to proceed")
		fmt.Println("Enter 'y' to accept or anything else to cancel the action")

		var answer string
		fmt.Scanf(answer)

		if answer != "y" {
			os.Exit(1)
		}
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

	for _, n := range currentNodes {
		go c.ConnectToNode(&n, &user)

	}

	os.Exit(0)
}

func (c *ConnectCommand) ConnectToNode(node *models.Node, user *models.User) {

	url := ""
	challenge, err := c.AuthenticationService.Authenticate(url, &user.Identity)

	if err != nil {
		panic("Failed to produce challenge, can't establish link with the node")
	}
	challengeBytes, err := hex.DecodeString(*challenge)
	if err != nil {
		panic("Not a valid encode string")
	}

	signature, err := c.WalletService.SignMessage(challengeBytes)
	if err != nil {
		fmt.Printf("Failed to produce a valid signature for the given challenge: %s", *challenge)
		return
	}

	token, err := c.AuthenticationService.GenerateSessionToken(&url)
	if err != nil {
		panic("couldn't generate session token")
	}

	handshake := map[string]interface{}{
		"action":          "verifyChallenge",
		"certificate":     user.Identity,
		"signedChallenge": signature,
		"sessionToken":    token,
	}

	connected := c.SocketService.Connect(&url, &handshake)

	if !connected {
		panic("Failed to connect to node")
	}

}

func (c *ConnectCommand) selectedNode(name string, nodes *[]models.Node) *models.Node {
	for _, n := range *nodes {
		if n.Name == name {
			return &n
		}
	}

	return nil
}
