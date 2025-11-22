package main

import (
	"fmt"

	"github.com/TompaSkitfet/conf-tree/internal/config"
)

func main() {
	data, err := config.LoadJSON()
	if err != nil {
		panic(err)
	}

	for _, child := range data.Children {
		for _, grandChild := range child.Children {
			fmt.Printf("%s: %v\n", grandChild.Key, grandChild.Value)
		}
	}
}
