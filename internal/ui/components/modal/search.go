package modal

import (
	"github.com/TompaSkitfet/conf-tree/internal/domain"
	search "github.com/TompaSkitfet/conf-tree/internal/search"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SearchModal struct {
	Input  textinput.Model
	Root   *domain.Node
	Result []*domain.Node
	Done   bool
}

func NewSearchModal(r *domain.Node) SearchModal {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 30
	return SearchModal{
		Input: ti,
		Done:  false,
		Root:  r,
	}
}

func (m SearchModal) Update(msg tea.Msg) (SearchModal, tea.Cmd) {
	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.Done = true
		default:
			m.Result = search.FuzzySearch(m.Input.Value(), m.Root)
		}
	}
	return m, cmd
}

func (m SearchModal) View() string {
	return m.Input.View()
}
