package main

import (
	"fmt"
	"path/filepath"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func main() {
	items, err := input.ReadInt(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	for i, itemA := range items {
		for j, itemB := range items {
			if itemA+itemB == 2020 {
				fmt.Printf("Entries found at [%d](%d), [%d](%d) - multiply to %d\n", i, itemA, j, itemB, itemA*itemB)
				return
			}
		}
	}
}
