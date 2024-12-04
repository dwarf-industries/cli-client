package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"client/interfaces"
	"client/models"
	"client/repositories"
)

type ImportUserIdentityCommand struct {
	Storage            interfaces.Storage
	PasswordManager    interfaces.PasswordManager
	UsersRepository    repositories.UsersRepository
	UserKeysRepository repositories.KeysRepository
	UserCertificates   repositories.Certificate
	CertificateService interfaces.CertificateService
	KeysService        interfaces.KeyService
}

func (u *ImportUserIdentityCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "import-user-identity [user_id] [path]",
		Short: "Imports a user identity and associates it with an existing user",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			userId, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Bad input data, aborting, expected number got string: %s\n", args[0])
				os.Exit(1)
			}
			path := args[1]
			u.Execute(&userId, &path)
		},
	}
}

func (u *ImportUserIdentityCommand) Execute(userId *int, path *string) {
	password := u.PasswordManager.Input()
	ok := u.PasswordManager.Match(password)
	if !ok {
		fmt.Println("Wrong account password, aborting!")
		os.Exit(1)
	}

	u.Storage.Open()
	defer u.Storage.Close()

	file, err := os.ReadFile(*path)
	if err != nil {
		fmt.Printf("Contact doesn't exist, please make sure that %s.json is isnide the folder", *path)
		os.Exit(1)
	}

	var contactDetails models.UserContact
	if err := json.Unmarshal(file, &contactDetails); err != nil {
		fmt.Println("Failed to parse contact details, malformed structure aborting!")
		os.Exit(1)
	}

	_, err = u.UsersRepository.UpdateUser(userId, &contactDetails.Identity, &contactDetails.EncryptionCertificate, &contactDetails.OrderSecret)

	if err != nil {
		fmt.Println("Failed to create a user, aborting!")
		os.Exit(1)
	}

	fmt.Println(path)
	fmt.Print("\r")
	fmt.Println()
	fmt.Println()
	fmt.Println("Identity PEM")
	fmt.Print("\r")

	fmt.Print(contactDetails.Identity)
	fmt.Println()
	fmt.Println()
	fmt.Print("\r")
	fmt.Println("Encryption Certificate PEM")
	fmt.Print("\r")
	fmt.Print(contactDetails.EncryptionCertificate)
	os.Exit(0)
}
