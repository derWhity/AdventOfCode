package main

import (
	"fmt"
	"math"
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
	// Every seat in the row is one bit
	var plane [128]uint8
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
		// Write our find into the plane map
		plane[row] |= (1 << col)
	}
	// Find our row - the first full row starts the search
	started := false
	for row, seats := range plane {
		if !started && seats == 255 { // 255 = 11111111 = full row
			started = true
		} else if started && seats != 255 {
			// Found our row
			col := int(math.Round(math.Log(float64(255^seats)) / math.Log(2))) // 2^n=x => n = log(x) / log(2)
			fmt.Printf("Found our seat in row %d, col %d => %d\n", row, col, row*8+col)
			break
		}
	}
}
