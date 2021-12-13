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
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	var score int
	for i, line := range items {
		opened := []rune{}
	Line:
		for j, r := range line {
			switch r {
			case '(', '{', '[', '<':
				// Opening
				opened = append(opened, r)
			default:
				// Closing
				lastOpened := opened[len(opened)-1]
				opened = opened[:len(opened)-1]
				required := lastOpened + 2 // All closing parentheses are 2 points after the opening one... Except...
				if lastOpened == '(' {
					required = ')'
				}
				if r != required {
					var s int
					switch r {
					case ')':
						s = 3
					case ']':
						s = 57
					case '}':
						s = 1197
					default:
						s = 25137
					}
					fmt.Printf("Syntax error in line %d col %d: %s required, but %s found - resulting in %5d points\n", i, j, string(required), string(r), s)
					score += s
					break Line
				}
			}
		}
	}
	fmt.Printf("Final score: %d\n", score)
}
