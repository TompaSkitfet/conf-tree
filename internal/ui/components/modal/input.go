package modal

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputModal struct {
	Input textinput.Model
	Done  bool
	Value string
}

func NewInputModal(initial string) InputModal {
	ti := textinput.New()
	ti.SetValue(initial)
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 30
	return InputModal{
		Input: ti,
		Done:  false,
		Value: "",
	}
}

func (m InputModal) Update(msg tea.Msg) (InputModal, tea.Cmd) {
	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.Done = true
			m.Value = ""
		case "enter":
			m.Done = true
			m.Value = m.Input.Value()
		}
	}
	return m, cmd
}

func (m InputModal) View() string {
	return m.Input.View()
}
