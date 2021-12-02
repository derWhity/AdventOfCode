package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
	"github.com/spf13/cast"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	var depth int64
	var x int64
	var aim int64
	for _, line := range items {
		if line != "" {
			params := strings.Split(line, " ")
			command := params[0]
			amount := cast.ToInt64(params[1])
			switch command {
			case "forward":
				x += amount
				depth += aim * amount
			case "up":
				aim -= amount
			case "down":
				aim += amount
			default:
				fmt.Printf("Illegal command %#v\n", command)
				panic("ARGH!")
			}
			fmt.Printf("Aim: %d | Depth: %d | X-Position: %d | Multiplied: %d\n", aim, depth, x, depth*x)
		}
	}
	fmt.Printf("Depth: %d | X-Position: %d | Multiplied: %d\n", depth, x, depth*x)
}
