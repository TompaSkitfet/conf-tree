package ui

import (
	"fmt"

	"github.com/TompaSkitfet/conf-tree/internal/domain"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

func TwoPanels(left, right, breadcrumb string) string {
	panels := lipgloss.JoinHorizontal(lipgloss.Top,
		LeftPanel.Render(left),
		RightPanel.Render(right))

	layout := lipgloss.JoinVertical(lipgloss.Left, TopBar.Render("Json explorer"), breadcrumb, panels)
	return layout

}

func BuildRightTree(selectedNode *domain.Node) string {
	t := tree.Root("")

	for _, v := range selectedNode.Children {
		if v.Children != nil {
			branch := tree.New().Root(fmt.Sprintf(" %s", v.Key))
			for _, c := range v.Children {
				if c.Value != nil {
					branch.Child(fmt.Sprintf(" %s: %v", c.Key, c.Value))
				} else {
					if c.Type == domain.ObjectNode {
						branch.Child(fmt.Sprintf(" %s", c.Key))
					} else {
						branch.Child(fmt.Sprintf(" %s", c.Key))
					}
				}
			}
			t.Child(branch)
		} else {
			t.Child(fmt.Sprintf(" %s: %v", v.Key, v.Value))
		}
	}
	return t.String()
}

func BuildBreadcrumbs(selectedNode *domain.Node) string {
	str := ""
	current := selectedNode
	for current != nil {
		str = fmt.Sprintf("%s/", current.Key) + str
		current = current.Parent
	}
	return str
}
