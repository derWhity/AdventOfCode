package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	var highestID uint64
	for _, item := range items {
		row, err := strconv.ParseUint(
			strings.ReplaceAll(
				strings.ReplaceAll(item[:7], "F", "0"),
				"B", "1"), 2, 64)
		if err != nil {
			panic(err)
		}
		col, err := strconv.ParseUint(
			strings.ReplaceAll(
				strings.ReplaceAll(item[7:], "L", "0"),
				"R", "1"), 2, 64)
		if err != nil {
			panic(err)
		}
		seatID := (row * 8) + col
		if seatID > highestID {
			highestID = seatID
		}
		fmt.Printf("Row: [%d] | Col: [%d] => %d\n", row, col, seatID)
	}
	fmt.Println("============================================")
	fmt.Printf("Highest seat ID is: %d\n", highestID)
}
