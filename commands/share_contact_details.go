package commands

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/spf13/cobra"

	"client/interfaces"
	"client/models"
	"client/repositories"
)

type ShareContactDetails struct {
	Storage            interfaces.Storage
	PasswordManager    interfaces.PasswordManager
	UsersRepository    repositories.UsersRepository
	UserKeysRepository repositories.KeysRepository
	UserCertificates   repositories.Certificate
	CertificateService interfaces.CertificateService
	KeysService        interfaces.KeyService
}

func (u *ShareContactDetails) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "generate-identity [name]",
		Short: "Generates contact details that can be exported and shared, allowing another client to import them and establish communication with the current user.",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			u.Execute(&name)
		},
	}
}

func (u *ShareContactDetails) Execute(name *string) {
	password := u.PasswordManager.Input()
	ok := u.PasswordManager.Match(password)
	if !ok {
		fmt.Println("Wrong account password, aborting!")
		os.Exit(1)
	}

	if name == nil {
		gofakeit.Seed(0)
		generateName := gofakeit.Name()
		name = &generateName
	}

	u.Storage.Open()
	defer u.Storage.Close()

	pub, priv, err := u.KeysService.IssueED25519Key()
	if err != nil {
		fmt.Println("failed to generate ED25519 key aborting!")
		os.Exit(1)
	}

	cert, err := u.CertificateService.IssueIdentityCertificate(*pub, *priv)
	if err != nil {
		fmt.Println("Failed to generate identity certificate with the given keys, aborting!")
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

	keySaved := u.UserKeysRepository.AddKey(cert, &encryptionPkHex)

	if !keySaved {
		fmt.Println("Failed to save key to the database, aborting!")
		os.Exit(1)
	}

	encryptionCertificate, err := u.CertificateService.IssueEncryptionCertificate(pk)
	if err != nil {
		fmt.Println("Failed to create encryption certificate aborting!")
		os.Exit(1)
	}

	file, err := os.Create(fmt.Sprintf("%v.json", *name))
	if err != nil {
		fmt.Println("Failed to create file for user contact")
		os.Exit(1)
	}
	defer file.Close()

	data := models.UserContact{
		OrderSecret:           "123",
		Identity:              *cert,
		EncryptionCertificate: *encryptionCertificate,
	}

	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("failed to parse data")
	}

	_, err = file.Write(json)
	if err != nil {
		fmt.Println("failed to write data to file")
	}
	fmt.Println(name)
	fmt.Print("\r")
	fmt.Println()
	fmt.Println()
	fmt.Println("Identity PEM")
	fmt.Print("\r")

	fmt.Print(cert)
	fmt.Println()
	fmt.Println()
	fmt.Print("\r")
	fmt.Println("Encryption Certificate PEM")
	fmt.Print("\r")
	fmt.Print(encryptionCertificate)

	fmt.Printf("%s.json exported", *name)
	os.Exit(0)
}
