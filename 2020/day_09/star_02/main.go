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

func sum(values ...int64) int64 {
	var out int64
	for _, v := range values {
		out += v
	}
	return out
}

func min(values ...int64) int64 {
	out := values[0]
	for _, v := range values {
		if v < out {
			out = v
		}
	}
	return out
}

func max(values ...int64) int64 {
	out := values[0]
	for _, v := range values {
		if v > out {
			out = v
		}
	}
	return out
}

func main() {
	items, err := input.ReadInt(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	var item int64
	for i := preambleLength; i < len(items); i++ {
		item = items[i]
		previous := items[(i - preambleLength):i]
		isMatching := match(item, previous)
		if !isMatching {
			break
		}
	}
	fmt.Printf("Item found: %d\n", item)
	// Now find the list of contiguous numbers summing up to this one
	var startPos int
	for {
		if startPos > len(items)-2 {
			fmt.Println("Reached end of the list - nothing found")
			break
		}
		for length := 2; startPos+length < len(items); length++ {
			values := items[startPos : startPos+length]
			s := sum(values...)
			if s > item {
				// No need to continue - we're already larger
				break
			}
			if s == item {
				fmt.Printf("The list %#v sums up to %d\n", values, item)
				minVal := min(values...)
				maxVal := max(values...)
				fmt.Printf("Sum of %d and %d is: %d\n", minVal, maxVal, (minVal + maxVal))
				return
			}
		}
		startPos++
	}
}
