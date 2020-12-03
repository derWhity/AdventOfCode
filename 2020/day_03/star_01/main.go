package main

import (
	"fmt"
	"path/filepath"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func getPos(row string, column uint) string {
	// Again, only ascii - so the number of bytes is okay
	localCol := column % uint(len(row))
	return string(row[localCol])
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	var column uint
	var treeCount uint
	for _, row := range items {
		if getPos(row, column) == "#" {
			treeCount++
		}
		column += 3
	}
	fmt.Printf("Number of trees: %d\n", treeCount)
}
