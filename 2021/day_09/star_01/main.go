package main

import (
	"fmt"
	"path/filepath"

	"github.com/derWhity/AdventOfCode/lib/input"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cast"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func isLowest(x, y int, grid [][]uint) bool {
	val := grid[x][y]
	if x > 0 && grid[x-1][y] <= val {
		return false
	}
	if x < len(grid)-1 && grid[x+1][y] <= val {
		return false
	}
	if y > 0 && grid[x][y-1] <= val {
		return false
	}
	if y < len(grid[x])-1 && grid[x][y+1] <= val {
		return false
	}
	return true
}

func printGrid(grid [][]uint) []uint {
	out := []uint{}
	for x, row := range grid {
		var line string
		for y, col := range row {
			val := fmt.Sprintf("%d ", col)
			if isLowest(x, y, grid) {
				out = append(out, col)
				line += aurora.BrightGreen(val).String()
			} else {
				line += val
			}
		}
		fmt.Println(line)
	}
	return out
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	grid := [][]uint{}
	for i, row := range items {
		grid = append(grid, []uint{})
		for _, col := range row {
			grid[i] = append(grid[i], cast.ToUint(string(col)))
		}
	}
	found := printGrid(grid)
	fmt.Printf("Found: %+v\n", found)
	var sum uint
	for _, val := range found {
		sum += (val + 1)
	}
	fmt.Printf("Sum: %d\n", sum)
}
