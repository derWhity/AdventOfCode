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

	currentGroup := map[rune]int{}
	groups := []uint{}
	numPeople := 0

	for i, line := range inputLines {
		if line != "" {
			// Line containing answers
			numPeople++
			taken := map[rune]bool{}
			for _, c := range line {
				if _, ok := taken[c]; !ok {
					currentGroup[c]++
				}
				taken[c] = true
			}
		}
		if line == "" || i == len(inputLines)-1 {
			var count uint
			// Count the fully answered ones
			for _, num := range currentGroup {
				if num == numPeople {
					count++
				}
			}
			groups = append(groups, count)

			// Reset
			currentGroup = map[rune]int{}
			numPeople = 0
		}

	}
	var sum uint
	fmt.Printf("Group counts: %+v\n", groups)
	for _, groupCount := range groups {
		sum += groupCount
	}
	fmt.Printf("Sum: %d\n", sum)
}
