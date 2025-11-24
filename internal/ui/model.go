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
	Width    int
	Height   int

	ShowOverlay bool
	EditingBool bool
	InputModal  modal.InputModal
	BoolModal   modal.BoolModal

	Help help.Model
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
	current := m.Tree.Current

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		UpdatePanelWidths(m.Width, m.Height)
	case tea.KeyMsg:
		switch {
		case m.ShowOverlay && m.EditingBool:
			var cmd tea.Cmd
			m.BoolModal, cmd = m.BoolModal.Update(msg)

			if m.BoolModal.Done {
				current.Value = m.BoolModal.Value
				current.Modified = true
				m.ShowOverlay = false
				m.EditingBool = false
			}
			return m, cmd

		case m.ShowOverlay && !m.EditingBool:
			var cmd tea.Cmd
			m.InputModal, cmd = m.InputModal.Update(msg)
			if m.InputModal.Done {
				if m.InputModal.Value != "" {
					current.Value = m.InputModal.Value
					current.Modified = true
				}
				m.ShowOverlay = false
			}
			return m, cmd
		case key.Matches(msg, Keys.Save):
			config.SaveToFile(m.Root, m.FileData)
			newData, err := config.LoadJSON(m.FileData.Name)
			if err != nil {
				panic(err)
			}
			m.Tree = tree.New(newData.Children)
		case key.Matches(msg, Keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, Keys.Up):
			m.Tree.MoveUp()
		case key.Matches(msg, Keys.Down):
			m.Tree.MoveDown()
		case key.Matches(msg, Keys.Right):
			if current.Type != domain.ValueNode {
				m.Tree.MoveRight()
			} else if m.Tree.Current.Type == domain.ValueNode {
				switch v := current.Value.(type) {
				case bool:
					m.BoolModal = modal.NewBoolModal(v)
					m.EditingBool = true
					m.ShowOverlay = true
				default:
					m.InputModal = modal.NewInputModal(fmt.Sprintf("%v", v))
					m.EditingBool = false
					m.ShowOverlay = true
				}
			}
		case key.Matches(msg, Keys.Left):
			m.Tree.MoveLeft()
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

	if m.ShowOverlay {
		if m.EditingBool {
			return overlay.Composite(BuildOverlay(m.BoolModal.View()), base, overlay.Center, overlay.Center, 0, 0)
		} else {
			return overlay.Composite(BuildOverlay(m.InputModal.View()), base, overlay.Center, overlay.Center, 0, 0)
		}
	}
	helpView := m.Help.View(Keys)
	return lipgloss.JoinVertical(lipgloss.Left, base, helpView)
}
