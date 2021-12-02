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

	for _, line := range items {
		if line != "" {
			params := strings.Split(line, " ")
			command := params[0]
			amount := cast.ToInt64(params[1])
			switch command {
			case "forward":
				x += amount
			case "up":
				depth -= amount
			case "down":
				depth += amount
			default:
				fmt.Printf("Illegal command %#v\n", command)
				panic("ARGH!")
			}
		}
	}
	fmt.Printf("Depth: %d | X-Position: %d | Multiplied: %d\n", depth, x, depth*x)
}
