package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"client/repositories"
)

type DeleteUserCommand struct {
	UsersRepository repositories.UsersRepository
}

func (d *DeleteUserCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "delete-user [ID]",
		Short: "Deletes a user",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			idString := args[1]
			id, err := strconv.Atoi(idString)
			if err != nil {
				fmt.Println("Wrong value format, only numbers allowed!")
				os.Exit(0)
				return

			}

			d.Execute(&id)
		},
	}
}

func (d *DeleteUserCommand) Execute(id *int) {

	deleted := d.UsersRepository.DeleteUser(*id)
	if !deleted {
		warning := color.New(color.FgRed, color.Bold).SprintFunc()
		fmt.Println(warning("\n⚠️  Either user doesn't exist, or failed to delete the user with the given ID."))
		os.Exit(0)
		return
	}

	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(header("User deleted!"))
	os.Exit(0)
}
