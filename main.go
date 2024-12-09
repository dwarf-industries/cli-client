package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"client/commands"
	"client/components"
	"client/di"
)

func main() {
	di.SetupServices()
	di.Setup()

	var rootCmd = &cobra.Command{
		Use: "client",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				logo := components.Logo{}
				logo.Init()

				_ = cmd.Help()
				os.Exit(0)
			}
		},
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
	CreateUserCommand := commands.CreateUserCommand{
		Storage:         di.DatabaseService(),
		UsersRepository: di.UsersRepository(),
	}
	importUserIdentityCommand := commands.ImportUserIdentityCommand{
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
	generateIdentityCommand := commands.GenerateIdentityCommand{
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

	connectCommand := commands.ConnectCommand{
		KeysRepository:        di.GetUserKeysRepository(),
		CertificateService:    di.GetCertificateService(),
		PasswordManager:       di.GetPasswordManager(),
		WalletService:         di.WalletService(),
		AuthenticationService: di.GetAuthenticationService(),
		UsersRepository:       di.UsersRepository(),
		NodeRepository:        di.GetNodesRepository(),
		RegisterService:       di.GetRegisterService(),
	}

	user := 2
	connectCommand.Execute(&user)

	rootCmd.AddCommand(setupAccountCommand.Executable())
	rootCmd.AddCommand(generateWalletCommand.Executable())
	rootCmd.AddCommand(rpcCommand.Executable())
	rootCmd.AddCommand(CreateUserCommand.Executable())
	rootCmd.AddCommand(importUserIdentityCommand.Executable())
	rootCmd.AddCommand(importContactsCommand.Executable())
	rootCmd.AddCommand(generateIdentityCommand.Executable())
	rootCmd.AddCommand(usersCommand.Executable())
	rootCmd.AddCommand(nodesCommand.Executable())
	rootCmd.AddCommand(connectCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
