package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"client/interfaces"
	"client/models"
	"client/repositories"
)

type AddUserCommand struct {
	Storage            interfaces.Storage
	PasswordManager    interfaces.PasswordManager
	UsersRepository    repositories.UsersRepository
	UserKeysRepository repositories.KeysRepository
	UserCertificates   repositories.Certificate
	CertificateService interfaces.CertificateService
	KeysService        interfaces.KeyService
}

func (u *AddUserCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "add-user [path]",
		Short: "Add a new user to the contact list located from path",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			u.Execute(&path)
		},
	}
}

func (u *AddUserCommand) Execute(path *string) {
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

	_, err = u.UsersRepository.AddUser(path, &contactDetails.Identity, &contactDetails.EncryptionCertificate, &contactDetails.OrderSecret)

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
