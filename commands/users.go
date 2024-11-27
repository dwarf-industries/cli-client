package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"client/repositories"
)

type UsersCommand struct {
	UsersRepository repositories.UsersRepository
}

func (u *UsersCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "add-user [certificate path] [name]",
		Short: "Add a new user to the contact list",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			certificate := args[0]
			name := args[1]

			u.Execute(&certificate, &name)
		},
	}
}

func (u *UsersCommand) Execute(certificatePath *string, name *string) {
	file, err := os.ReadFile(*certificatePath)
	if err != nil {
		fmt.Println("File doesn't exist!")
		os.Exit(1)
		return
	}
	u.UsersRepository.AddUser(&file, name)
	os.Exit(0)
}
