package commands

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"client/interfaces"
	"client/models"
	"client/repositories"
)

type ImportUsersCommand struct {
	UsersRepository        repositories.UsersRepository
	CertificatesRepository repositories.Certificate
	PasswordManager        interfaces.PasswordManager
	password               string
}

func (i *ImportUsersCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "import-users [path]",
		Args:  cobra.ExactArgs(0),
		Short: "Imports users based on a valid contract list",
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			i.Execute(&path)
		},
	}
}

func (i *ImportUsersCommand) Execute(contactDir *string) {
	fmt.Println("Please enter your password!")
	fmt.Scanf(i.password)

	files, err := os.ReadDir(*contactDir)
	if err != nil {
		panic("Folder doesn't exist, aborting!")
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileType := file.Type()

		name := file.Name()
		if fileType.String() != "json" {
			continue
		}

		filePath := filepath.Join(*contactDir, name)
		content, _ := os.ReadFile(filePath)

		var user models.UserContract
		if err := json.Unmarshal(content, &user); err != nil {
			fmt.Println("Failed to parse contact info, skipping, malformed json, please check your data!")
			continue
		}

		_, err := i.UsersRepository.GetUserByName(&user.Name)

		if err == nil {
			fmt.Printf("Contact with name %v already exists!", user.Name)
			continue
		}

		i.importUser(user)
	}
}

func (i *ImportUsersCommand) importUser(user models.UserContract) {
	created, err := i.UsersRepository.AddUser(&user.Name)

	if err != nil {
		fmt.Println("Failed to save new user with name: ", user.Name)
	}

	identityPublic, err := i.PasswordManager.Encrypt([]byte(user.Identity), []byte(i.password))

	if err != nil {
		fmt.Println("Failed to encrypt identity certificate!")
		return
	}

	identityCertificateType := 2
	identityCert := hex.EncodeToString(*identityPublic)
	saveCertificate := i.CertificatesRepository.AddCertificate(&created, &identityCertificateType, &identityCert)

	if !saveCertificate {
		fmt.Println("Failed to save identity certificate aborting!")
		return
	}

	encryptionCertificate, err := i.PasswordManager.Encrypt([]byte(user.EncryptionCertificate), []byte(i.password))

	if err != nil {
		fmt.Println("Failed to encrypt encryption certificate!")
		return
	}
	encryptCert := hex.EncodeToString(*encryptionCertificate)
	encryptionCertificateType := 1
	saveCertificate = i.CertificatesRepository.AddCertificate(&created, &encryptionCertificateType, &encryptCert)

	if !saveCertificate {
		fmt.Println("Failed to save encryption certificate aborting!")
		return
	}

	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(header("User Created"))

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Name", "Identity PEM", "Encryption PEM"})
	t.AppendRow(table.Row{"Id", created})
	t.AppendRow(table.Row{"Name", user.Name})
	t.AppendRow(table.Row{"Identity PEM", user.Identity})
	t.AppendRow(table.Row{"Encryption PEM", user.EncryptionCertificate})
	t.Render()
}
