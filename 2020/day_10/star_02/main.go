package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func countPossibilities(items []int, collected []int) uint64 {
	var out uint64
	coll := []int{}
	coll = append(coll, collected...)
	coll = append(coll, items[0])
	if len(items) == 1 {
		// We are at the last item - so we found a match
		//fmt.Printf("Match: %+v\n", collected)
		return 1
	}
	if len(items) >= 2 {
		if items[1]-items[0] <= 3 {
			out += countPossibilities(items[1:], coll)
		}
	}
	if len(items) >= 3 {
		if items[2]-items[0] <= 3 {
			out += countPossibilities(items[2:], coll)
		}
	}
	if len(items) >= 4 {
		if items[3]-items[0] <= 3 {
			out += countPossibilities(items[3:], coll)
		}
	}
	return out
}

func main() {
	fmt.Println("-------------")
	items, err := input.ReadIntNative(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	// Sort the adapters by joltage
	items = append(items, 0) // Add the seat outlet
	iSlice := sort.IntSlice(items)
	iSlice.Sort()
	iSlice = append(iSlice, iSlice[len(iSlice)-1]+3) // Add our own device
	fmt.Printf("Slice: %+v\n", iSlice)
	// Counting recursively will take a looooooong time - time to grab some coffee
	count := countPossibilities(iSlice, []int{})
	fmt.Printf("Different possibilities %d\n", count)
}
