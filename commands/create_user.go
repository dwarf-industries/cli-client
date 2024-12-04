package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"client/interfaces"
	"client/repositories"
)

type CreateUserCommand struct {
	Storage         interfaces.Storage
	UsersRepository repositories.UsersRepository
}

func (u *CreateUserCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "create-user [name]",
		Short: "Creates a new empty user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			name := args[0]
			u.Execute(&name)
		},
	}
}

func (u *CreateUserCommand) Execute(name *string) {

	u.Storage.Open()
	defer u.Storage.Close()

	id, err := u.UsersRepository.CreateUser(name)
	if err != nil {
		fmt.Println("Failed to create a user, aborting!")
		os.Exit(1)
	}

	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(header(fmt.Printf("🚀 User created with name: %s and id: %v!", *name, id)))
	os.Exit(0)
}
