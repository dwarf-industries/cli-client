package commands

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"client/interfaces"
)

type SetupAccountCommand struct {
	WalletService   interfaces.WalletService
	PasswordManager interfaces.PasswordManager
}

func (s *SetupAccountCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "setup [private key]",
		Short: "configure your client account",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := args[0]
			s.Execute(&privateKey)
		},
	}
}

func (s *SetupAccountCommand) Execute(wallet *string) {
	password := s.PasswordManager.Input()
	privateKey, err := s.WalletService.SetWallet(wallet, password)
	if err != nil {
		fmt.Println("Failed to import wallet!")
		fmt.Println(err)
		os.Exit(1)
		return
	}
	fmt.Println("")
	publicKeyHex := s.WalletService.GetAddressForPrivateKey(&privateKey)
	privateKeyHex := privateKey.D.Text(16)

	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(header("🚀 Wallet Imported Successful!"))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Key Type", "Key Value"})
	t.AppendRow(table.Row{"Public Key", publicKeyHex})
	t.AppendRow(table.Row{"Private Key", privateKeyHex})
	t.Render()

	warning := color.New(color.FgRed, color.Bold).SprintFunc()
	fmt.Println(warning("\n⚠️  IMPORTANT: Keep a secure copy of your private key. It is required for wallet recovery and cannot be retrieved if lost."))

	walletBalance := s.WalletService.GetBalance((*common.Address)(common.FromHex(publicKeyHex)))

	if walletBalance.Int64() == 0 {
		fmt.Println("\n⚠️  ATTENTION: it appears that you wallet balance is empty, you will need to fund it in order to use it as operations wallet!")
	}

	fmt.Println("you can learn more about funding and why it's required under 'oracle network-operation-info'")
	infoHeader := color.New(color.BgCyan, color.Bold).SprintFunc()
	fmt.Println(infoHeader("for more information regarding wallets 'oracle wallets-info'"))

	os.Exit(0)
}
