package views

import (
	"fmt"
	"math/big"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"client/contracts"
	"client/converters"
	"client/di"
	"client/models"
)

type ListView struct {
	choices   []models.Node
	cursor    int
	selected  map[string]struct{}
	feeSetter *contracts.FeeSetter
}

func InitialModel(nodes []models.Node, activeNodes map[string]struct{}) ListView {
	contractAddress := common.HexToAddress(os.Getenv("FEE_SETTER"))
	feeContract, err := contracts.NewFeeSetter(contractAddress, di.RpcService().GetClient())
	if err != nil {
		panic("Failed to initialize fee contract")
	}
	return ListView{
		choices:   nodes,
		selected:  activeNodes,
		feeSetter: feeContract,
	}
}

func (l *ListView) GetSelectedNodes() map[string]struct{} {
	return l.selected
}

func (l ListView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return l, tea.Quit
		case "up", "k":
			if l.cursor > 0 {
				l.cursor--
			}
		case "down", "j":
			if l.cursor < len(l.choices)-1 {
				l.cursor++
			}
		case "a":
			for _, n := range l.choices {
				if _, ok := l.selected[n.Name]; !ok {
					l.selected[n.Name] = struct{}{}
				}
			}
		case "d":
			l.selected = make(map[string]struct{})
		case "enter", " ":
			_, ok := l.selected[l.choices[l.cursor].Name]
			if ok {
				delete(l.selected, l.choices[l.cursor].Name)
			} else {
				l.selected[l.choices[l.cursor].Name] = struct{}{}
			}
		}
	}

	return l, nil
}

func (l ListView) Init() tea.Cmd {

	return nil
}

func (l ListView) View() string {
	s := "Please select the list of active nodes use for interaction with the network\n\n"
	var num int
	for i, choice := range l.choices {
		cursor := " "
		if l.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := l.selected[choice.Name]; ok {
			num++
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}
	cost, err := l.feeSetter.GetCostPerKilobyte(&bind.CallOpts{})
	if err != nil {
		fmt.Println("failed to get cost per kylobyte fee from the network!")
	}

	estimatedPrice := big.NewInt(1).Mul(cost, big.NewInt(int64(num)))
	formated := converters.WeiToEth(estimatedPrice)
	priceFormat := fmt.Sprintf("%sEstimated cost for transaction on the network based on the amount of nodes used: %s ETH\n", "\nControls:\n'q' to quit,\n'a' to select all\n'd' to deselect all\n", formated.String())
	s += priceFormat
	return s
}
