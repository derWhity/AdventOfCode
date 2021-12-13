package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkLine(line string) []rune {
	opened := []rune{}
	for _, r := range line {
		switch r {
		case '(', '{', '[', '<':
			// Opening
			opened = append(opened, r)
		default:
			// Closing
			lastOpened := opened[len(opened)-1]
			required := lastOpened + 2 // All closing parentheses are 2 points after the opening one... Except...
			if lastOpened == '(' {
				required = ')'
			}
			if r != required {
				return nil
			}
			opened = opened[:len(opened)-1]
		}
	}
	return opened // The remainder
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	scores := sort.IntSlice{}
	for i, line := range items {
		if remainder := checkLine(line); len(remainder) > 0 {
			fmt.Printf("Line %d with remainder of %s\n", i, string(remainder))
			var score int
			for j := len(remainder) - 1; j >= 0; j-- {
				score *= 5
				switch remainder[j] {
				case '(':
					score += 1
				case '[':
					score += 2
				case '{':
					score += 3
				default:
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}
	scores.Sort()
	fmt.Printf("Final score: %#v\n", scores[len(scores)/2])
}
