package config

import (
	"encoding/json"
	"os"

	"github.com/TompaSkitfet/conf-tree/internal/domain"
)

func SaveToFile(data *domain.Node, fileData domain.FileData) error {
	jsonData := NodeToInterface(data)

	raw, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileData.Name, raw, 0644)
}

func NodeToInterface(n *domain.Node) any {
	switch n.Type {
	case domain.ValueNode:
		return n.Value
	case domain.ObjectNode:
		m := map[string]any{}
		for _, child := range n.Children {
			m[child.Key] = NodeToInterface(child)
		}
		return m
	case domain.ArrayNode:
		arr := make([]any, len(n.Children))
		for i, child := range n.Children {
			arr[i] = NodeToInterface(child)
		}
		return arr
	}
	return nil

}
