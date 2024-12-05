package commands

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/spf13/cobra"

	"client/interfaces"
	"client/models"
	"client/repositories"
)

type GenerateIdentityCommand struct {
	Storage            interfaces.Storage
	PasswordManager    interfaces.PasswordManager
	UsersRepository    repositories.UsersRepository
	UserKeysRepository repositories.KeysRepository
	UserCertificates   repositories.Certificate
	CertificateService interfaces.CertificateService
	KeysService        interfaces.KeyService
}

func (u *GenerateIdentityCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "generate-identity [name] [user_id]",
		Short: "Generates contact details that can be exported and shared, allowing another client to import them and establish communication with the current user.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			userData, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Printf("Bad input data, expected a number got %v\n", args[1])
				return
			}
			u.Execute(&name, &userData)
		},
	}
}

func (u *GenerateIdentityCommand) Execute(name *string, userId *int) {
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

	user, err := u.UsersRepository.GetUserById(*userId)
	if err != nil {
		fmt.Println("User doesn't exist, aborting")
		os.Exit(1)
	}

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
	identityPrivBytes := priv.Seed()
	identityPrivateKey, err := u.PasswordManager.Encrypt(identityPrivBytes, []byte(*password))
	if err != nil {
		fmt.Println("Failed to save the encryption private key aborting!")
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

	encryptionCertificate, err := u.CertificateService.IssueEncryptionCertificate(pk)
	if err != nil {
		fmt.Println("Failed to create encryption certificate aborting!")
		os.Exit(1)
	}

	orderSecret, err := u.CertificateService.GenerateOrderSecret(*name)
	if err != nil {
		fmt.Println("Failed to generate order secret, aborting")
		os.Exit(1)
	}

	orderSecretEncrypted, err := u.PasswordManager.Encrypt([]byte(*orderSecret), []byte(*password))
	if err != nil {
		fmt.Println("Failed to save the encryption private key aborting!")
		os.Exit(1)
	}

	encryptionIdentityHex := hex.EncodeToString(*identityPrivateKey)
	encryptionPkHex := hex.EncodeToString(*encryptionPk)
	encryptedOrderSecretHex := hex.EncodeToString(*orderSecretEncrypted)
	keySaved := u.UserKeysRepository.AddKey(
		cert,
		encryptionCertificate,
		&encryptionPkHex,
		&encryptionIdentityHex,
		&encryptedOrderSecretHex,
		&user.Id,
	)

	if !keySaved {
		fmt.Println("Failed to save key to the database, aborting!")
		os.Exit(1)
	}

	file, err := os.Create(fmt.Sprintf("%v.json", *name))
	if err != nil {
		fmt.Println("Failed to create file for user contact")
		os.Exit(1)
	}
	defer file.Close()

	data := models.UserContact{
		OrderSecret:           *orderSecret,
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
