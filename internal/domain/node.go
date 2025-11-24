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
	Modified bool
	Parent   *Node
	Children []*Node
}
