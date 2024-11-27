package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"client/commands"
	"client/di"
)

func main() {
	di.SetupServices()
	var rootCmd = &cobra.Command{
		Use: "oracle",
	}

	addWalletcommand := commands.AddWalletCommand{
		WalletService: di.WalletService(),
	}
	generateWalletCommand := commands.GenerateWalletCommand{
		WalletService: di.WalletService(),
	}
	rpcCommand := commands.SetRpcCommand{
		RpcService: di.RpcService(),
	}
	addUserCommand := commands.AddUserCommand{
		UsersRepository: di.UsersRepository(),
	}

	rootCmd.AddCommand(addWalletcommand.Executable())
	rootCmd.AddCommand(generateWalletCommand.Executable())
	rootCmd.AddCommand(rpcCommand.Executable())
	rootCmd.AddCommand(addUserCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
