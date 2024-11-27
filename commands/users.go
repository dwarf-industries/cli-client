package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
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
		Run: func(cmd *cobra.Command, args []string) {
			u.Execute()
		},
	}
}

func (u *UsersCommand) Execute() {
	users, err := u.UsersRepository.GetAllUsers()

	if err != nil {
		os.Exit(1)
		return
	}

	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(header("Contact List"))

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Name", "Created At"})

	for _, user := range users {
		t.AppendRow(table.Row{"Id", user.Id})
		t.AppendRow(table.Row{"Name", user.Name})
		t.AppendRow(table.Row{"Name", user.CreatedAt})
		t.Render()

	}
}
