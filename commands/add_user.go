package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

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
		Use:   "add-user [name]",
		Short: "Add a new user to the contact list located from the contacts directory",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			u.Execute(&name)
		},
	}
}

func (u *AddUserCommand) Execute(name *string) {
	password := u.PasswordManager.Input()
	ok := u.PasswordManager.Match(password)
	if !ok {
		fmt.Println("Wrong account password, aborting!")
		os.Exit(1)
	}

	u.Storage.Open()
	defer u.Storage.Close()

	contact := strings.Join([]string{
		"./contacts/",
		*name,
	}, "")

	file, err := os.ReadFile(contact)
	if err != nil {
		fmt.Printf("Contact doesn't exist, please make sure that %s.json is isnide the folder", *name)
		os.Exit(1)
	}

	var contactDetails models.UserContact
	if err := json.Unmarshal(file, &contactDetails); err != nil {
		fmt.Println("Failed to parse contact details, malformed structure aborting!")
		os.Exit(1)
	}

	_, err = u.UsersRepository.AddUser(name, &contactDetails.Identity, &contactDetails.EncryptionCertificate, &contactDetails.OrderSecret)

	if err != nil {
		fmt.Println("Failed to create a user, aborting!")
		os.Exit(1)
	}

	fmt.Println(name)
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
