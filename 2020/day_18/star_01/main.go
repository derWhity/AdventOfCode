package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func calc(parts []string) (int, []string) {
	// Parts are numbers, operators or parentheses
	var out int
	var op string
	for {
		if len(parts) == 0 {
			return out, parts
		}
		val := parts[0]
		parts = parts[1:] // Remove first element
		var value int
		switch val {
		case "(":
			// Sub-calculation
			value, parts = calc(parts)
		case ")":
			// Finished
			return out, parts
		case "+":
			op = "+"
			continue
		case "*":
			op = "*"
			continue
		default:
			// Read the value to operate with
			var err error
			value, err = strconv.Atoi(val)
			failOnError(err)
		}
		// If we get here, let's perform the operation and continue
		switch op {
		case "*":
			out *= value
		case "+":
			out += value
		default:
			// No operation. Just set the value
			out = value
		}
	}
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	var sum int
	for _, line := range items {
		// Split the string into single operations
		var ops []string
		for _, block := range strings.Split(line, " ") {
			for _, item := range strings.Split(block, "(") {
				if item == "" {
					ops = append(ops, "(")
				} else {
					for _, subItem := range strings.Split(item, ")") {
						if subItem == "" {
							ops = append(ops, ")")
						} else {
							ops = append(ops, subItem)
						}
					}
				}
			}
		}
		fmt.Printf("Operation list: %#v\n", ops)
		result, _ := calc(ops)
		fmt.Printf("Result: %d\n", result)
		sum += result
	}
	fmt.Printf("Sum is: %d\n", sum)
}
