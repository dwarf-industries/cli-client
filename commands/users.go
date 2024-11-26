package commands

import (
	"os"

	"github.com/spf13/cobra"
)

type UsersCommand struct {
}

func (u *UsersCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "new-wallet [password]",
		Short: "Generates a new empty wallet",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}

func (u *UsersCommand) Execute(password *string) {

	os.Exit(0)
}
