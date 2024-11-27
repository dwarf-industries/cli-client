package commands

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"client/interfaces"
	"client/repositories"
)

type ConnectCommand struct {
	WalletService         interfaces.WalletService
	AuthenticationService interfaces.AuthenticationService
	UsersRepository       repositories.UsersRepository
	SocketService         interfaces.SocketConnection
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

	url := ""
	challenge, err := c.AuthenticationService.Authenticate(url, &user.Certificate)

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
		"certificate":     user.Certificate,
		"signedChallenge": signature,
		"sessionToken":    token,
	}

	connected := c.SocketService.Connect(&url, &handshake)

	if !connected {
		panic("Failed to connect to node")
	}

	os.Exit(0)
}
