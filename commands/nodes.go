package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"client/interfaces"
)

type NodesCommand struct {
	RegisterService interfaces.RegisterService
}

func (n *NodesCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "nodes",
		Short: "Shows a list of available nodes",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}

func (n *NodesCommand) Execute() {
	nodes, err := n.RegisterService.Oracles()
	if err != nil {
		warning := color.New(color.FgRed, color.Bold).SprintFunc()
		fmt.Println(warning("Failed to retrive a list of nodes"))

		os.Exit(1)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	warning := color.New(color.FgGreen, color.Bold).SprintFunc()
	fmt.Println(warning("Nodes loaded."))
	t.AppendHeader(table.Row{"Node ID", "IP", "Reputation"})

	for _, node := range nodes {
		t.AppendRow(table.Row{node.Name, node.Ip, node.Reputation.Int64()})
		t.Render()
	}

	os.Exit(0)
}
