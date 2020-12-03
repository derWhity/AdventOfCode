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

func ride(items []string, incrementColumn, incrementRow uint) uint {
	var column uint
	var row uint
	var treeCount uint
	for {
		rowContent := items[row]
		if getPos(rowContent, column) == "#" {
			treeCount++
		}
		column += incrementColumn
		row += incrementRow
		if row >= uint(len(items)) {
			break
		}
	}
	return treeCount
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	var treeCount uint
	treeCount = ride(items, 1, 1)
	treeCount *= ride(items, 3, 1)
	treeCount *= ride(items, 5, 1)
	treeCount *= ride(items, 7, 1)
	treeCount *= ride(items, 1, 2)
	fmt.Printf("Number of trees: %d\n", treeCount)
}
