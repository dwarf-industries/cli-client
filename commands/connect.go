package commands

import (
	"os"

	"github.com/spf13/cobra"
)

type ConnectCommand struct {
}

func (c *ConnectCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "connect [ID]",
		Short: "Connects to a specific user based on it's CertificateID, to list users 'client users'",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func (c *ConnectCommand) Execute(password *string) {

	os.Exit(0)
}
