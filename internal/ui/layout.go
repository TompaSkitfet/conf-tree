package ui

import (
	"fmt"

	"github.com/TompaSkitfet/conf-tree/internal/domain"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

func TwoPanels(left, right string) string {
	panels := lipgloss.JoinHorizontal(lipgloss.Top,
		LeftPanel.Render(left),
		RightPanel.Render(right))

	layout := lipgloss.JoinVertical(lipgloss.Left, TopBar.Render("Json explorer"), panels)
	return layout

}

func BuildRightTree(selectedNode *domain.Node) string {
	t := tree.Root("")

	for _, v := range selectedNode.Children {
		if v.Children != nil {
			branch := tree.New().Root(fmt.Sprintf("%s, %s", v.Key, v.Parent))
			for _, c := range v.Children {
				branch.Child(fmt.Sprintf("%s, %s", c.Key, c.Parent))
			}
			t.Child(branch)
		} else {
			t.Child(fmt.Sprintf("%s, %s", v.Key, v.Parent))
		}
	}
	return t.String()
}
