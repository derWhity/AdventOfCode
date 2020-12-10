package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func main() {
	items, err := input.ReadIntNative(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	// Sort the adapters by joltage
	iSlice := sort.IntSlice(items)
	iSlice.Sort()
	fmt.Printf("Slice: %+v\n", iSlice)
	counter := 0
	differences := map[int]int{}
	for _, adapter := range iSlice {
		fmt.Printf("From %d to %d = %d\n", counter, adapter, adapter-counter)
		differences[adapter-counter]++
		counter = adapter
	}
	// Add our own device
	differences[3]++
	fmt.Printf("Differences: %#v\n", differences)
	fmt.Printf("1 Jolts * 3 Jolts = %d\n", differences[1]*differences[3])
}
