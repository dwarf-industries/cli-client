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
	di.Setup()
	var rootCmd = &cobra.Command{
		Use: "client",
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
		Storage:            di.DatabaseService(),
		UsersRepository:    di.UsersRepository(),
		PasswordManager:    di.GetPasswordManager(),
		UserKeysRepository: di.GetUserKeysRepository(),
		UserCertificates:   di.GetUserCertificates(),
		CertificateService: di.GetCertificateService(),
		KeysService:        di.GetKeyService(),
	}
	importUserCommand := commands.ImportUsersCommand{
		UsersRepository:        di.UsersRepository(),
		PasswordManager:        di.GetPasswordManager(),
		CertificatesRepository: di.GetUserCertificates(),
		Storage:                di.DatabaseService(),
	}
	usersCommand := commands.UsersCommand{
		UsersRepository: di.UsersRepository(),
	}
	nodesCommand := commands.NodesCommand{
		Storage:         di.DatabaseService(),
		RegisterService: di.GetRegisterService(),
		NodesRepository: di.GetNodesRepository(),
	}

	nodesCommand.Execute()

	rootCmd.AddCommand(addWalletcommand.Executable())
	rootCmd.AddCommand(generateWalletCommand.Executable())
	rootCmd.AddCommand(rpcCommand.Executable())
	rootCmd.AddCommand(addUserCommand.Executable())
	rootCmd.AddCommand(usersCommand.Executable())
	rootCmd.AddCommand(importUserCommand.Executable())
	rootCmd.AddCommand(nodesCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
