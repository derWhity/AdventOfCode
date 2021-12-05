package main

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/derWhity/AdventOfCode/lib/input"
	"github.com/spf13/cast"
)

var (
	lineRegex = regexp.MustCompile(`^([0-9]+),([0-9]+) \-> ([0-9]+),([0-9]+)$`)
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

type point struct {
	x uint
	y uint
}

type line struct {
	a point
	b point
}

func max(a uint, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func min(a uint, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func readLines(items []string) ([]line, uint, uint) {
	var out []line
	var maxX, maxY uint
	for _, lineStr := range items {
		coords := lineRegex.FindStringSubmatch(lineStr)
		x1 := cast.ToUint(coords[1])
		maxX = max(maxX, x1)
		y1 := cast.ToUint(coords[2])
		maxY = max(maxY, y1)
		x2 := cast.ToUint(coords[3])
		maxX = max(maxX, x2)
		y2 := cast.ToUint(coords[4])
		maxY = max(maxY, y2)
		l := line{
			a: point{x: x1, y: y1},
			b: point{x: x2, y: y2},
		}
		out = append(out, l)
	}
	return out, maxX, maxY
}

func drawLine(coords [][]uint, l line) {
	if l.a.x == l.b.x {
		// vertical
		for y := min(l.a.y, l.b.y); y <= max(l.a.y, l.b.y); y++ {
			coords[l.a.x][y]++
		}
	} else if l.a.y == l.b.y {
		// horizontal
		for x := min(l.a.x, l.b.x); x <= max(l.a.x, l.b.x); x++ {
			coords[x][l.a.y]++
		}
	}
}

func countAreasGt(coords [][]uint, max uint) uint {
	var sum uint
	for _, row := range coords {
		for _, val := range row {
			if val > max {
				sum++
			}
		}
	}
	return sum
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	lineDefs, maxX, maxY := readLines(items)
	fmt.Printf("Read %d lines with [maxX: %d, maxY: %d]\n", len(lineDefs), maxX, maxY)
	coords := make([][]uint, maxX+1)
	for i := range coords {
		coords[i] = make([]uint, maxY+1)
	}
	for _, lDef := range lineDefs {
		drawLine(coords, lDef)
	}
	for y := 0; y <= int(maxY); y++ {
		for x := 0; x <= int(maxY); x++ {
			fmt.Printf("%2d ", coords[x][y])
		}
		fmt.Println("")
	}
	fmt.Printf("Coordinates with density > 1: %d\n", countAreasGt(coords, 1))

}
