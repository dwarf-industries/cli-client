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

	setupAccountCommand := commands.SetupAccountCommand{
		WalletService:   di.WalletService(),
		PasswordManager: di.GetPasswordManager(),
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
	importContactsCommand := commands.ImportContactsCommand{
		Storage:            di.DatabaseService(),
		UsersRepository:    di.UsersRepository(),
		PasswordManager:    di.GetPasswordManager(),
		UserKeysRepository: di.GetUserKeysRepository(),
		UserCertificates:   di.GetUserCertificates(),
		CertificateService: di.GetCertificateService(),
		KeysService:        di.GetKeyService(),
	}
	exportContactDetails := commands.ShareContactDetails{
		Storage:            di.DatabaseService(),
		UsersRepository:    di.UsersRepository(),
		PasswordManager:    di.GetPasswordManager(),
		UserKeysRepository: di.GetUserKeysRepository(),
		UserCertificates:   di.GetUserCertificates(),
		CertificateService: di.GetCertificateService(),
		KeysService:        di.GetKeyService(),
	}
	usersCommand := commands.UsersCommand{
		UsersRepository: di.UsersRepository(),
		Storage:         di.DatabaseService(),
	}
	nodesCommand := commands.NodesCommand{
		Storage:         di.DatabaseService(),
		RegisterService: di.GetRegisterService(),
		NodesRepository: di.GetNodesRepository(),
	}

	rootCmd.AddCommand(setupAccountCommand.Executable())
	rootCmd.AddCommand(generateWalletCommand.Executable())
	rootCmd.AddCommand(rpcCommand.Executable())
	rootCmd.AddCommand(addUserCommand.Executable())
	rootCmd.AddCommand(importContactsCommand.Executable())
	rootCmd.AddCommand(exportContactDetails.Executable())
	rootCmd.AddCommand(usersCommand.Executable())
	rootCmd.AddCommand(nodesCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
