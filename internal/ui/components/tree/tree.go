package tree

import (
	"github.com/TompaSkitfet/conf-tree/internal/domain"
	"strings"
)

type Tree struct {
	Nodes  []*domain.Node
	Cursor int
}

func New(nodes []*domain.Node) Tree {
	return Tree{Nodes: nodes, Cursor: 0}
}

func (t Tree) View() string {
	var b strings.Builder

	for i, n := range t.Nodes {
		cursor := " "
		if i == t.Cursor {
			cursor = "➤ "
		}
		b.WriteString(cursor + n.Key + "\n")
	}
	return b.String()
}

func (t *Tree) MoveUp() {
	if t.Cursor > 0 {
		t.Cursor--
	}
}

func (t *Tree) MoveDown() {
	if t.Cursor < len(t.Nodes)-1 {
		t.Cursor++
	}
}

func (t *Tree) MoveRight() {
	if t.Selected().Children != nil {
		t.Nodes = t.Selected().Children
	}
}

func (t *Tree) MoveLeft() {
	if t.Selected().Parent.Parent.Children != nil {
		t.Nodes = t.Selected().Parent.Parent.Children
	}
}

func (t Tree) Selected() *domain.Node {
	if len(t.Nodes) == 0 {
		return nil
	}
	return t.Nodes[t.Cursor]
}
