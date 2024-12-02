package commands

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v6"
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
		Short: "Add a new user to the contact list",
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

	gofakeit.Seed(0)
	generateName := gofakeit.Name()

	if !ok {
		fmt.Println("Wrong account password, aborting!")
		os.Exit(1)
	}

	u.Storage.Open()
	defer u.Storage.Close()

	created, err := u.UsersRepository.AddUser(name)

	if err != nil {
		fmt.Println("Failed to create user, can't proceed")
		os.Exit(1)
	}

	pub, priv, err := u.KeysService.IssueED25519Key()
	if err != nil {
		fmt.Println("failed to generate ED25519 key aborting!")
		os.Exit(1)
	}

	encryptedHexBytes, err := u.PasswordManager.Encrypt(*pub, []byte(*password))
	if err != nil {
		fmt.Println("Failed to encrypt public key hex bytes")
		os.Exit(1)
	}

	encryptedHexBytesHex := hex.EncodeToString(*encryptedHexBytes)
	u.UserKeysRepository.AddKey(&created, 2, &encryptedHexBytesHex)

	encryptPrivKey, err := u.PasswordManager.Encrypt(*priv, []byte(*password))
	if err != nil {
		fmt.Println("Failed to encrypt private key hex bytes")
		os.Exit(1)
	}
	encryptPrivKeyHex := hex.EncodeToString(*encryptPrivKey)
	u.UserKeysRepository.AddKey(&created, 3, &encryptPrivKeyHex)

	cert, err := u.CertificateService.IssueIdentityCertificate(*pub, *priv)

	if err != nil {
		fmt.Println("Failed to generate identity certificate with the given keys, aborting!")
		os.Exit(1)
	}

	addUserCertificate := u.UserCertificates.AddCertificate(&created, 2, cert)

	if !addUserCertificate {
		fmt.Println("Failed to save user certificate aborting!")
		os.Exit(1)
	}

	pk, err := u.KeysService.IssueRSAKey(4096)

	if err != nil {
		fmt.Println("failed to generate encryption key aborting!")
		os.Exit(1)
	}

	privBytes := x509.MarshalPKCS1PrivateKey(pk)
	encryptionPk, err := u.PasswordManager.Encrypt(privBytes, []byte(*password))
	if err != nil {
		fmt.Println("Failed to save the encryption private key aborting!")
		os.Exit(1)
	}
	encryptionPkHex := hex.EncodeToString(*encryptionPk)
	keySaved := u.UserKeysRepository.AddKey(&created, 1, &encryptionPkHex)

	if !keySaved {
		fmt.Println("Failed to save key to the database, aborting!")
		os.Exit(1)
	}

	encryptionCertificate, err := u.CertificateService.IssueEncryptionCertificate(pk)

	if err != nil {
		fmt.Println("Failed to create encryption certificate aborting!")
		os.Exit(1)

	}
	certificateCreated := u.UserCertificates.AddCertificate(&created, 1, encryptionCertificate)
	if !certificateCreated {
		fmt.Println("Failed to save encryption certificate aborting!")
		os.Exit(1)
	}

	userContact := models.UserContact{
		Name:                  generateName,
		Identity:              *cert,
		EncryptionCertificate: *encryptionCertificate,
	}

	file, err := os.Create(fmt.Sprintf("%v_user.json", created))
	if err != nil {
		fmt.Println("Failed to create file for user contact")
		os.Exit(1)
	}
	defer file.Close()

	fmt.Println(userContact.Name)
	fmt.Print("\r")
	fmt.Println()
	fmt.Println()
	fmt.Println("Identity PEM")
	fmt.Print("\r")

	fmt.Print(userContact.Identity)
	fmt.Println()
	fmt.Println()
	fmt.Print("\r")
	fmt.Println("Encryption Certificate PEM")
	fmt.Print("\r")
	fmt.Print(userContact.EncryptionCertificate)
	os.Exit(0)
}
