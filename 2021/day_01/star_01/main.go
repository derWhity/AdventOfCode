package main

import (
	"fmt"
	"path/filepath"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	items, err := input.ReadInt(filepath.Join("..", "input.txt"))
	failOnError(err)
	var amount int64
	var last int64
	for i, line := range items {
		if i > 0 && line > last {
			fmt.Printf("%d (INCREASED)\n", line)
			amount++
		} else {
			fmt.Printf("%d (-)\n", line)
		}
		last = line
	}
	fmt.Printf("Number of increases: %d\n", amount)
}
