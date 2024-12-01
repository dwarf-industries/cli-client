package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"client/models"
)

type ListView struct {
	choices  []models.Node
	cursor   int
	selected map[string]struct{}
}

func InitialModel(nodes []models.Node, activeNodes map[string]struct{}) ListView {
	return ListView{
		choices:  nodes,
		selected: activeNodes,
	}
}

func (m ListView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "a":
			for _, n := range m.choices {
				if _, ok := m.selected[n.Name]; !ok {
					m.selected[n.Name] = struct{}{}
				}
			}
		case "d":
			m.selected = make(map[string]struct{})
		case "enter", " ":
			_, ok := m.selected[m.choices[m.cursor].Name]
			if ok {
				delete(m.selected, m.choices[m.cursor].Name)
			} else {
				m.selected[m.choices[m.cursor].Name] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m ListView) Init() tea.Cmd {

	return nil
}

func (m ListView) View() string {
	s := "Please select the list of active nodes use for interaction with the network\n\n"

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[choice.Name]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}

	s += "\nControls:\n'q' to quit,\n'a' to select all\n'd' to deselect all \n"
	return s
}
