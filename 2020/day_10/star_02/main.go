package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func main() {
	fmt.Println("-------------")
	items, err := input.ReadIntNative(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	// Sort the adapters by joltage
	iSlice := sort.IntSlice(items)
	iSlice.Sort()
	iSlice = append(iSlice, iSlice[len(iSlice)-1]+3) // Add our own device
	fmt.Printf("Slice: %+v\n", iSlice)
	pathsTo := map[int]int{0: 1}
	for _, v := range iSlice {
		pathsTo[v] = pathsTo[v-1] + pathsTo[v-2] + pathsTo[v-3]
	}
	fmt.Printf("Different possibilities %d\n", pathsTo[iSlice[len(iSlice)-1]])
}
