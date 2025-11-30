package list

import (
	"github.com/TompaSkitfet/conf-tree/internal/domain"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListModel struct {
	Items        []*domain.Node
	Active       bool
	Cursor       int
	Height       int
	SelectedItem *domain.Node
}

func NewListModel(items []*domain.Node, height int) ListModel {
	return ListModel{
		Items:  items,
		Height: 10,
		Active: false,
	}
}

func (m ListModel) Init() tea.Cmd { return nil }

func (m ListModel) Update(msg tea.Msg) (ListModel, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Items)-1 {
				m.Cursor++
			}
		case "enter":
			if len(m.Items) > 0 {
				m.SelectedItem = m.Items[m.Cursor]
			}
			return m, nil
		}

	}
	return m, nil
}

func (m ListModel) View() string {
	var out string

	max := 10
	if max > len(m.Items) {
		max = len(m.Items)
	}

	for i := 0; i < max; i++ {
		item := m.Items[i]

		if i == m.Cursor && m.Active {
			out += highlightStyle.Render("> "+item.Key) + "\n"
		} else {
			out += normalStyle.Render(" "+item.Key) + "\n"
		}
	}
	return out
}

var highlightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

var normalStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
