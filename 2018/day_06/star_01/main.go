package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

type coordinate struct {
	x    int
	y    int
	area int
}

func loadCoords(lines []string) ([]*coordinate, int, int, int, int) {
	minX := -1
	minY := -1
	maxX := -1
	maxY := -1
	coords := []*coordinate{}
	for _, line := range lines {
		c := coordinate{}
		vals := strings.Split(line, ",")
		tmp, _ := strconv.ParseInt(strings.TrimSpace(vals[0]), 10, strconv.IntSize)
		c.x = int(tmp)
		tmp, _ = strconv.ParseInt(strings.TrimSpace(vals[1]), 10, strconv.IntSize)
		c.y = int(tmp)
		fmt.Printf("(%d, %d)\n", c.x, c.y)
		if minX < 0 || c.x < minX {
			minX = c.x
		}
		if minY < 0 || c.y < minY {
			minY = c.y
		}
		if maxX < 0 || c.x > maxX {
			maxX = c.x
		}
		if maxY < 0 || c.y > maxY {
			maxY = c.y
		}
		coords = append(coords, &c)
	}
	fmt.Printf("%d coords found with rect of %d, %d, %d, %d\n", len(coords), minX, minY, maxX, maxY)
	return coords, minX, minY, maxX, maxY
}

func winner(coords []*coordinate, x, y int) *coordinate {
	minDist := float64(-1)
	var minCoord *coordinate
	minCount := 0
	for _, coord := range coords {
		dist := math.Abs(float64(coord.x-x)) + math.Abs(float64(coord.y-y))
		if minDist == -1 || dist < minDist {
			minDist = dist
			minCoord = coord
			minCount = 1
		} else if dist == minDist {
			minCount++
		}

	}
	if minCount == 1 {
		return minCoord
	}
	return nil
}

func exec(input string) *coordinate {
	fmt.Println("Starting...")
	lines := strings.Split(input, "\n")
	coords, minX, minY, maxX, maxY := loadCoords(lines)
	// Calculate the "winner" for each coord
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			minCoord := winner(coords, x, y)

			if minCoord != nil && minCoord.area >= 0 {
				// The ones on the borders have infinite areas - so let's drop them as candidates
				if x == minX || y == minY || x == maxX || y == maxY {
					minCoord.area = -1
				} else {
					minCoord.area++
				}
			}
		}
	}
	// And now finally find the one with the largest area
	max := 0
	var found *coordinate
	for _, coord := range coords {
		fmt.Printf("%+v\n", coord)
		if coord.area > max {
			max = coord.area
			found = coord
		}
	}
	return found
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	coord := exec(strings.TrimSpace(string(input)))
	fmt.Printf(
		"Size of largest non-infinite area (%d, %d): %d\n",
		coord.x,
		coord.y,
		coord.area,
	)
}
