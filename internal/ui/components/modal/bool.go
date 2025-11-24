package modal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type BoolModal struct {
	Value bool
	Done  bool
}

func NewBoolModal(initial bool) BoolModal {
	return BoolModal{
		Value: initial,
		Done:  false,
	}
}

func (m BoolModal) Update(msg tea.Msg) (BoolModal, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "right", " ":
			m.Value = !m.Value
		case "enter", "esc":
			m.Done = true
		}
	}
	return m, nil
}

func (m BoolModal) View() string {
	trueStyle := lipgloss.NewStyle()
	falseStyle := lipgloss.NewStyle()

	if m.Value {
		trueStyle = trueStyle.Bold(true).Underline(true)
	} else {
		falseStyle = falseStyle.Bold(true).Underline(true)
	}

	return trueStyle.Render("true") + "  " + falseStyle.Render("false")
}
