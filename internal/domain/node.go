package domain

import "fmt"

type NodeType int

const (
	ObjectNode NodeType = iota
	ArrayNode
	ValueNode
)

type Node struct {
	Key      string
	Type     NodeType
	Value    any
	Parent   *Node
	Children []*Node
}

func (n *Node) String() string {
	switch n.Type {
	case ObjectNode:
		return fmt.Sprintf("ObjectNode(%s)", n.Key)
	case ArrayNode:
		return fmt.Sprintf("ArrayNode(%s)", n.Key)
	case ValueNode:
		return fmt.Sprintf("ValueNode(%s: %v)", n.Key, n.Value)
	}
	return "UnknownNode"
}
