package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"client/interfaces"
)

type SetRpcCommand struct {
	RpcService interfaces.RpcService
}

func (s *SetRpcCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "use-rpc [rpc url]",
		Short: "configure custom RPC to connect to the blockchain",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rpc := args[0]
			s.Execute(&rpc)
		},
	}
}

func (s *SetRpcCommand) Execute(rpc *string) {
	err := s.RpcService.OverrideClient(rpc)
	if err != nil {
		fmt.Printf("Failed to override existing client, error: %v", err)
		return
	}
	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(header("🚀 Rpc Updated!"))
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Key Value"})
	t.AppendRow(table.Row{rpc})
	t.Render()

	os.Exit(0)
}
