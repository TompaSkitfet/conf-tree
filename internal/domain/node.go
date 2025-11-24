package domain

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
