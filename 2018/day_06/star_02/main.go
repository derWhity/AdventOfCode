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
	x int
	y int
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

func distToAllCoords(coords []*coordinate, x, y int) int {
	dist := 0
	for _, coord := range coords {
		dist += int(math.Abs(float64(coord.x-x)) + math.Abs(float64(coord.y-y)))
	}
	return dist
}

func exec(input string) int {
	fmt.Println("Starting...")
	lines := strings.Split(input, "\n")
	coords, minX, minY, maxX, maxY := loadCoords(lines)
	area := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			distSum := distToAllCoords(coords, x, y)
			if distSum < 10000 {
				// Within the safe area
				area++
			}
		}
	}
	return area
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	safeAreaSize := exec(strings.TrimSpace(string(input)))
	fmt.Printf(
		"Size of the safe area is: %d\n",
		safeAreaSize,
	)
}
