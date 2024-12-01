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

	for i, choice := range l.choices {

		cursor := " "
		if l.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := l.selected[choice.Name]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}

	s += "\nControls:\n'q' to quit,\n'a' to select all\n'd' to deselect all \n"
	return s
}
