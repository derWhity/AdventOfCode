package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/derWhity/AdventOfCode/lib/input"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cast"
)

const (
	typeLowest = 1
	typeBasin  = 2
	typeWall   = 3
)

type entry struct {
	value uint
	typ   uint
}

func newEntry(val uint) *entry {
	return &entry{
		value: val,
	}
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func isLowest(x, y int, grid [][]*entry) bool {
	val := grid[x][y].value
	if x > 0 && grid[x-1][y].value <= val {
		return false
	}
	if x < len(grid)-1 && grid[x+1][y].value <= val {
		return false
	}
	if y > 0 && grid[x][y-1].value <= val {
		return false
	}
	if y < len(grid[x])-1 && grid[x][y+1].value <= val {
		return false
	}
	grid[x][y].typ = typeLowest
	return true
}

func fillBasin(x, y int, grid [][]*entry) int {
	var out int
	// Left
	if x > 0 && grid[x-1][y].typ == 0 {
		if grid[x-1][y].value != 9 {
			grid[x-1][y].typ = typeBasin
			out += 1 + fillBasin(x-1, y, grid)
		} else {
			grid[x-1][y].typ = typeWall
		}
	}
	// Right
	if x < len(grid)-1 && grid[x+1][y].typ == 0 {
		if grid[x+1][y].value != 9 {
			grid[x+1][y].typ = typeBasin
			out += 1 + fillBasin(x+1, y, grid)
		} else {
			grid[x+1][y].typ = typeWall
		}
	}
	// Above
	if y > 0 && grid[x][y-1].typ == 0 {
		if grid[x][y-1].value != 9 {
			grid[x][y-1].typ = typeBasin
			out += 1 + fillBasin(x, y-1, grid)
		} else {
			grid[x][y-1].typ = typeWall
		}
	}
	// Below
	if y < len(grid[x])-1 && grid[x][y+1].typ == 0 {
		if grid[x][y+1].value != 9 {
			grid[x][y+1].typ = typeBasin
			out += 1 + fillBasin(x, y+1, grid)
		} else {
			grid[x][y+1].typ = typeWall
		}
	}
	return out
}

func printGrid(grid [][]*entry) sort.IntSlice {
	basinSizes := sort.IntSlice{}
	for x, row := range grid {
		for y, col := range row {
			if isLowest(x, y, grid) {
				col.typ = typeLowest
				basinSizes = append(basinSizes, 1+fillBasin(x, y, grid))
			}
		}
	}
	for _, row := range grid {
		for _, col := range row {
			val := fmt.Sprintf("%d ", col.value)
			switch col.typ {
			case typeBasin:
				val = aurora.Yellow(val).String()
			case typeLowest:
				val = aurora.BrightGreen(val).String()
			case typeWall:
				val = aurora.BgGray(6, val).String()
			}
			fmt.Printf("%s", val)
		}
		fmt.Println("")
	}
	return basinSizes
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	grid := [][]*entry{}
	for i, row := range items {
		grid = append(grid, []*entry{})
		for _, col := range row {
			grid[i] = append(grid[i], newEntry(cast.ToUint(string(col))))
		}
	}
	basins := printGrid(grid)
	fmt.Printf("Sizes:\n%+v\n", basins)
	basins.Sort()
	fmt.Printf("Sizes (sorted):\n%+v\n", basins)
	result := basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
	fmt.Printf("Result: %d\n", result)
}
