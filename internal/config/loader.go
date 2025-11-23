package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TompaSkitfet/conf-tree/internal/domain"
)

func LoadJSON() (*domain.Node, error) {
	raw, err := os.ReadFile("test.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var data any
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return buildNode("", data, nil), nil
}

func buildNode(key string, value any, parent *domain.Node) *domain.Node {
	switch v := value.(type) {
	case map[string]any:
		n := &domain.Node{
			Key:      key,
			Type:     domain.ObjectNode,
			Parent:   parent,
			Children: []*domain.Node{},
		}

		for childKey, childVal := range v {
			child := buildNode(childKey, childVal, n)
			n.Children = append(n.Children, child)
		}
		return n

	case []any:
		n := &domain.Node{
			Key:      key,
			Type:     domain.ArrayNode,
			Parent:   parent,
			Children: []*domain.Node{},
		}
		for i, childVal := range v {
			childKey := fmt.Sprintf("%s[%d]", key, i)
			child := buildNode(childKey, childVal, n)
			n.Children = append(n.Children, child)
		}
		return n
	default:
		return &domain.Node{
			Key:    key,
			Type:   domain.ValueNode,
			Value:  v,
			Parent: parent,
		}
	}
}
