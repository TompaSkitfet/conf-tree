package modal

import (
	"github.com/TompaSkitfet/conf-tree/internal/domain"
	search "github.com/TompaSkitfet/conf-tree/internal/search"
	"github.com/TompaSkitfet/conf-tree/internal/ui/components/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SearchModal struct {
	Input  textinput.Model
	Root   *domain.Node
	Result list.ListModel
	Done   bool
	Active bool
}

func NewSearchModal(r *domain.Node) SearchModal {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 30
	return SearchModal{
		Input:  ti,
		Done:   false,
		Root:   r,
		Active: true,
		Result: list.NewListModel(nil, 10),
	}
}

func (m SearchModal) Update(msg tea.Msg) (SearchModal, tea.Cmd) {
	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case m.Active:
			if msg.String() == "esc" {
				m.Active = false
				m.Input.Blur()
				return m, nil
			}
			if msg.String() == "down" {
				m.Active = false
				m.Input.Blur()
				m.Result.Active = true
				return m, nil
			}

			m.Result.Items = search.FuzzySearch(m.Input.Value(), m.Root)

			return m, nil
		case !m.Active:
			if msg.String() == "esc" {
				m.Done = true
			}
			updated, _ := m.Result.Update(msg)
			m.Result = updated
			return m, nil
		}
	}
	return m, cmd
}

func (m SearchModal) View() string {
	return m.Input.View()
}
