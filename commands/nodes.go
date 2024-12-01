package commands

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"client/interfaces"
	"client/repositories"
	"client/views"
)

type NodesCommand struct {
	Storage         interfaces.Storage
	RegisterService interfaces.RegisterService
	NodesRepository repositories.NodesRepository
}

func (n *NodesCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "nodes",
		Short: "Shows a list of available nodes",
		Run: func(cmd *cobra.Command, args []string) {
			if err := n.Execute(); err != nil {
				warning := color.New(color.FgRed, color.Bold).SprintFunc()
				fmt.Println(warning("Failed to execute nodes command:", err))
				os.Exit(1)
			}
		},
	}
}

func (n *NodesCommand) Execute() error {
	nodes, err := n.RegisterService.Oracles()
	if err != nil {
		warning := color.New(color.FgRed, color.Bold).SprintFunc()
		fmt.Println(warning("Failed to retrieve a list of nodes"))
		return err
	}

	n.Storage.Open()
	defer n.Storage.Close()

	selectedNodes, err := n.NodesRepository.Selected()

	if err != nil {
		selectedNodes = &[]string{}
		fmt.Println("Failed to get saved nodes")
	}

	if len(nodes) == 0 {
		fmt.Println("No nodes found.")
		return nil
	}
	activeNodes := make(map[string]struct{})

	for _, selected := range *selectedNodes {
		activeNodes[selected] = struct{}{}
	}

	choiceList := views.InitialModel(nodes, activeNodes)
	p := tea.NewProgram(choiceList)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	updatedNodeList := choiceList.GetSelectedNodes()

	var updateRes []string
	for k := range updatedNodeList {
		updateRes = append(updateRes, k)
		n.NodesRepository.Select(&k)
	}

	n.NodesRepository.RemoveAllExcept(updateRes)

	return nil
}
