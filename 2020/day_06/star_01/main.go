package main

import (
	"fmt"
	"path/filepath"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func main() {
	inputLines, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}

	groups := []map[rune]uint{
		{}, // Initialize the first group already
	}
	groupNo := 0

	for _, line := range inputLines {
		if line != "" {
			// Line containing answers
			for _, c := range line {
				groups[groupNo][c]++
			}
		}
		if line == "" {
			groupNo++
			groups = append(groups, map[rune]uint{})
		}

	}
	var sum int
	fmt.Printf("Groups: %+v\n", groups)
	for i, group := range groups {
		fmt.Printf("Group %d: %d\n", i, len(group))
		sum += len(group)
	}
	fmt.Printf("Sum: %d\n", sum)
}
