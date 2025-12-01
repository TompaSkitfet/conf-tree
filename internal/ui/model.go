package ui

import (
	"fmt"

	"github.com/TompaSkitfet/conf-tree/internal/config"
	"github.com/TompaSkitfet/conf-tree/internal/domain"
	"github.com/TompaSkitfet/conf-tree/internal/ui/components/modal"
	"github.com/TompaSkitfet/conf-tree/internal/ui/components/tree"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

type Model struct {
	Tree     tree.Tree
	Root     *domain.Node
	FileData domain.FileData
	width    int
	height   int

	ActiveOverlay modal.OverlayType
	SearchModal   modal.SearchModal
	InputModal    modal.InputModal
	BoolModal     modal.BoolModal
	Help          help.Model

	err error
}

func New(root *domain.Node, fileData domain.FileData) Model {
	return Model{
		Tree:       tree.New(root.Children),
		Root:       root,
		FileData:   fileData,
		InputModal: modal.InputModal{},
		Help:       help.New(),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	current := m.Tree.Selected()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		UpdatePanelWidths(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch {
		case m.err != nil && key.Matches(msg, Keys.Quit):
			m.err = nil

		case m.ActiveOverlay != modal.OverlayNone:
			switch m.ActiveOverlay {
			case modal.OverlayEditBool:
				return m.updateBoolOverlay(msg, current)

			case modal.OverlayEditInput:
				return m.updateInputOverlay(msg, current)

			case modal.OverlaySearch:
				return m.updateSearchOverlay(msg)
			}

		case key.Matches(msg, Keys.Save):
			err := config.SaveToFile(m.Root, m.FileData)
			if err != nil {
				m.err = err
				return m, nil
			}
			newData, err := config.LoadJSON(m.FileData.Name)
			if err != nil {
				m.err = err
				return m, nil
			}
			m.Root = newData
			m.Tree = tree.New(newData.Children)
			return m, nil

		case key.Matches(msg, Keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, Keys.Up):
			m.Tree.MoveUp()

		case key.Matches(msg, Keys.Down):
			m.Tree.MoveDown()

		case key.Matches(msg, Keys.Right):
			if current.Type != domain.ValueNode {
				m.Tree.MoveRight()
			} else if current.Type == domain.ValueNode {
				switch v := current.Value.(type) {
				case bool:
					m.BoolModal = modal.NewBoolModal(v)
					m.ActiveOverlay = modal.OverlayEditBool
				default:
					m.InputModal = modal.NewInputModal(fmt.Sprintf("%v", v))
					m.ActiveOverlay = modal.OverlayEditInput
				}
			}
		case key.Matches(msg, Keys.Left):
			m.Tree.MoveLeft()
		case key.Matches(msg, Keys.Search):
			m.SearchModal = modal.NewSearchModal(m.Root)
			m.ActiveOverlay = modal.OverlaySearch
		}
	}
	return m, nil
}

func (m Model) View() string {
	selected := m.Tree.Selected()
	right := "No selection"
	if selected != nil {
		right = BuildRightTree(selected)
	}

	base := TwoPanels(m.Tree.View(), right, BuildBreadcrumbs(selected))

	if m.err != nil {
		return overlay.Composite(m.err.Error(), base, overlay.Center, overlay.Center, 0, 0)
	}

	switch m.ActiveOverlay {

	case modal.OverlaySearch:
		return overlay.Composite(BuildSearchBox(m.SearchModal, m.SearchModal.Result), base, overlay.Center, overlay.Center, 0, 0)

	case modal.OverlayEditBool:
		return overlay.Composite(BuildOverlay(m.BoolModal.View()), base, overlay.Center, overlay.Center, 0, 0)
	case modal.OverlayEditInput:
		return overlay.Composite(BuildOverlay(m.InputModal.View()), base, overlay.Center, overlay.Center, 0, 0)
	}

	helpView := m.Help.View(Keys)
	return lipgloss.JoinVertical(lipgloss.Left, base, helpView)
}

func (m *Model) updateBoolOverlay(msg tea.KeyMsg, n *domain.Node) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.BoolModal, cmd = m.BoolModal.Update(msg)

	if m.BoolModal.Done {
		n.Value = m.BoolModal.Value
		n.Modified = true
		m.ActiveOverlay = modal.OverlayNone
	}
	return m, cmd
}

func (m *Model) updateInputOverlay(msg tea.KeyMsg, n *domain.Node) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.InputModal, cmd = m.InputModal.Update(msg)
	if m.InputModal.Done {
		if m.InputModal.Value != "" {
			n.Value = m.InputModal.Value
			n.Modified = true
		}
		m.ActiveOverlay = modal.OverlayNone
	}
	return m, cmd
}

func (m *Model) updateSearchOverlay(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.SearchModal, cmd = m.SearchModal.Update(msg)
	if m.SearchModal.Done {
		m.Tree.Nodes = m.SearchModal.ResultNode.Parent.Children
		m.Tree.Cursor = m.Tree.FindSelected(m.SearchModal.ResultNode)
		m.ActiveOverlay = modal.OverlayNone
	}
	return m, cmd
}
