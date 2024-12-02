package commands

import (
	"fmt"
	"os"
	"strings"

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
		Use:   "import-users [directory]",
		Short: "Add a new users to the contact list located from directory path",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			files, err := os.ReadDir(path)
			if err != nil {
				fmt.Println("Path not found aborting!")
				os.Exit(1)
			}

			addUserCmd := &AddUserCommand{
				Storage:            u.Storage,
				PasswordManager:    u.PasswordManager,
				UsersRepository:    u.UsersRepository,
				UserKeysRepository: u.UserKeysRepository,
				UserCertificates:   u.UserCertificates,
				CertificateService: u.CertificateService,
				KeysService:        u.KeysService,
			}

			for _, f := range files {
				if f.IsDir() {
					continue
				}
				combine := strings.Join([]string{
					path,
					f.Name(),
				}, "/")

				addUserCmd.Execute(&combine)
			}
		},
	}
}
