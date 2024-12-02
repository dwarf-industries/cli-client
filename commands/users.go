package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"client/interfaces"
	"client/repositories"
)

type UsersCommand struct {
	UsersRepository repositories.UsersRepository
	Storage         interfaces.Storage
}

func (u *UsersCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "users",
		Short: "Lists all users from the contact list",
		Run: func(cmd *cobra.Command, args []string) {
			u.Execute()
		},
	}
}

func (u *UsersCommand) Execute() {
	u.Storage.Open()
	defer u.Storage.Close()
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
		t.AppendRow(table.Row{user.Id, user.Name, user.CreatedAt})
	}
	t.Render()
	os.Exit(0)
}
