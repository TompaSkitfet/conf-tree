package tree

import (
	"fmt"
	"strings"

	"github.com/TompaSkitfet/conf-tree/internal/domain"
)

type Tree struct {
	Nodes   []*domain.Node
	Cursor  int
	Current *domain.Node
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
		switch n.Type {
		case domain.ObjectNode:
			b.WriteString(cursor + " " + n.Key + "\n")
		case domain.ArrayNode:
			b.WriteString(cursor + " " + n.Key + "\n")
		case domain.ValueNode:
			if n.Modified == true {
				b.WriteString(fmt.Sprintf("%s  %s: %v *\n", cursor, n.Key, n.Value))
			} else {
				b.WriteString(fmt.Sprintf("%s  %s: %v\n", cursor, n.Key, n.Value))
			}
		default:
			b.WriteString(cursor + "" + n.Key + "\n")
		}
	}
	return b.String()
}

func (t *Tree) MoveUp() {
	if t.Cursor > 0 {
		t.Cursor--
	}
	t.Current = t.Nodes[t.Cursor]
}

func (t *Tree) MoveDown() {
	if t.Cursor < len(t.Nodes)-1 {
		t.Cursor++
	}
	t.Current = t.Nodes[t.Cursor]
}

func (t *Tree) MoveRight() {
	if t.Selected().Children != nil {
		t.Nodes = t.Selected().Children
		t.Cursor = 0
		t.Current = t.Nodes[t.Cursor]
	}
}

func (t *Tree) MoveLeft() {
	if t.Selected().Parent.Parent != nil {
		t.Nodes = t.Selected().Parent.Parent.Children
		t.Cursor = 0
		t.Current = t.Nodes[t.Cursor]
	}
}

func (t Tree) Selected() *domain.Node {
	if len(t.Nodes) == 0 {
		return nil
	}
	return t.Nodes[t.Cursor]
}
