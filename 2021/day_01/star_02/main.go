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
	for i := 0; i < len(items)-2; i++ {
		window := []int64{items[i], items[i+1], items[i+2]}
		sum := window[0] + window[1] + window[2]
		fmt.Printf("Window: %+v => %d", window, sum)
		if i > 0 && last < sum {
			amount++
			fmt.Println(" (increased)")
		} else {
			fmt.Println(" (-)")
		}
		last = sum
	}
	fmt.Printf("Number of increases: %d\n", amount)
}
