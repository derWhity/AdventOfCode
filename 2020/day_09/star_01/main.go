package main

import (
	"fmt"
	"path/filepath"

	"github.com/derWhity/AdventOfCode/lib/input"
)

const (
	preambleLength = 25
)

func match(value int64, previous []int64) bool {
	for _, a := range previous {
		for _, b := range previous {
			if a == b {
				// Numbers must not be the same
				continue
			}
			if a+b == value {
				return true
			}
		}
	}
	return false
}

func main() {
	items, err := input.ReadInt(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	for i := preambleLength; i < len(items); i++ {
		item := items[i]
		previous := items[(i - preambleLength):i]
		isMatching := match(item, previous)
		fmt.Printf("Item: %d | %+v => %t\n", item, previous, isMatching)
		if !isMatching {
			break
		}
	}
}
