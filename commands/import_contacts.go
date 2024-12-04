package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/spf13/cobra"

	"client/interfaces"
	"client/repositories"
)

type ImportContactsCommand struct {
	Storage            interfaces.Storage
	PasswordManager    interfaces.PasswordManager
	UsersRepository    repositories.UsersRepository
	UserKeysRepository repositories.KeysRepository
	UserCertificates   repositories.Certificate
	CertificateService interfaces.CertificateService
	KeysService        interfaces.KeyService
}

func (u *ImportContactsCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "import-users [users ([names])] [directory]",
		Short: "Add a new users to the contact list located from directory path and generates users, if [names] is empty it will generate the names ",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			names := args[0]
			namesSplit := strings.Split(names, ",")
			path := args[1]
			files, err := os.ReadDir(path)
			if err != nil {
				fmt.Println("Path not found aborting!")
				os.Exit(1)
			}

			importUserIdentityCommand := &ImportUserIdentityCommand{
				Storage:            u.Storage,
				PasswordManager:    u.PasswordManager,
				UsersRepository:    u.UsersRepository,
				UserKeysRepository: u.UserKeysRepository,
				UserCertificates:   u.UserCertificates,
				CertificateService: u.CertificateService,
				KeysService:        u.KeysService,
			}

			if len(namesSplit) < len(files) {
				gofakeit.Seed(0)

				for i := len(namesSplit); i < len(files); i++ {
					generateName := gofakeit.Name()
					namesSplit[i] = generateName
				}
			}

			for i, f := range files {
				userName := namesSplit[i]
				if f.IsDir() {
					continue
				}
				combine := strings.Join([]string{
					path,
					f.Name(),
				}, "/")

				created, err := u.UsersRepository.CreateUser(&userName)
				if err != nil {
					fmt.Println("Aborting, failed to create user")
					os.Exit(1)
				}
				importUserIdentityCommand.Execute(&created, &combine)
			}
		},
	}
}
