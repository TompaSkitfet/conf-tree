package ui

import (
	"github.com/TompaSkitfet/conf-tree/internal/domain"
	"github.com/sahilm/fuzzy"
)

type nodeSource []*domain.Node

func (ns nodeSource) Len() int {
	return len(ns)
}

func (ns nodeSource) String(i int) string {
	return ns[i].Key
}

func FuzzySearch(s string, searchRoot *domain.Node) []*domain.Node {
	searchList := FlattenTree(searchRoot)
	src := nodeSource(searchList)
	matches := fuzzy.FindFrom(s, src)

	var result []*domain.Node
	for _, m := range matches {
		result = append(result, searchList[m.Index])
	}

	return result
}

func FlattenTree(root *domain.Node) []*domain.Node {
	var nodeList []*domain.Node
	var walk func(n *domain.Node)

	walk = func(n *domain.Node) {
		nodeList = append(nodeList, n)
		for _, c := range n.Children {
			walk(c)
		}
	}
	walk(root)

	return nodeList
}
