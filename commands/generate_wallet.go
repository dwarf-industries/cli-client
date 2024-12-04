package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"client/interfaces"
)

type GenerateWalletCommand struct {
	WalletService interfaces.WalletService
}

func (g *GenerateWalletCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "new-wallet [password]",
		Short: "Generates a new empty wallet",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			password := args[0]
			g.Execute(&password)
		},
	}
}

func (g *GenerateWalletCommand) Execute(password *string) {
	wallet, err := g.WalletService.NewWallet()
	if err != nil {
		fmt.Println("Failed to create new wallet, please submit an issue or try again!")
		fmt.Println(err)
	}

	fmt.Println("")
	publicKeyHex := g.WalletService.GetAddressForPrivateKey(wallet)
	privateKeyHex := wallet.D.Text(16)

	_, err = g.WalletService.SetWallet(&privateKeyHex, password)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create new wallet.")
		os.Exit(1)
		return
	}

	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(header("🚀 Wallet Creation Successful!"))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Key Type", "Key Value"})
	t.AppendRow(table.Row{"Public Key", publicKeyHex})
	t.AppendRow(table.Row{"Private Key", privateKeyHex})
	t.Render()

	warning := color.New(color.FgRed, color.Bold).SprintFunc()
	fmt.Println(warning("\n⚠️  IMPORTANT: Keep a secure copy of your private key. It is required for wallet recovery and cannot be retrieved if lost."))
	fmt.Println("\n⚠️  IMPORTANT: new wallets have an empty balance, you should use the public key for the newly created wallet to fund any operations such as registering it as a valid oracle")
	fmt.Println("you can learn more about funding and why it's required under 'client network-operation-info'")
	infoHeader := color.New(color.BgCyan, color.Bold).SprintFunc()
	fmt.Println(infoHeader("for more information regarding wallets 'client wallets-info'"))

	os.Exit(0)
}
