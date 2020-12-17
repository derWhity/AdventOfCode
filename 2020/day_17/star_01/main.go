package main

import (
	"fmt"
	"path/filepath"

	"github.com/derWhity/AdventOfCode/lib/input"
)

// grid represents a 3-dimensional grid of cube states
type grid map[int]map[int]map[int]bool

func (g grid) set(x, y, z int, val bool) {
	plane, ok := g[x]
	if !ok {
		plane = map[int]map[int]bool{}
		g[x] = plane
	}
	row, ok := plane[y]
	if !ok {
		row = map[int]bool{}
		plane[y] = row
	}
	row[z] = val
}

func (g grid) get(x, y, z int) bool {
	plane, ok := g[x]
	if !ok {
		return false
	}
	row, ok := plane[y]
	if !ok {
		return false
	}
	return row[z]
}

// determineNewState calculates the new state the given cube will have
func (g grid) determineNewStateAt(xPos, yPos, zPos int) bool {
	var activeNeighbours uint
	for x := xPos - 1; x <= xPos+1; x++ {
		for y := yPos - 1; y <= yPos+1; y++ {
			for z := zPos - 1; z <= zPos+1; z++ {
				if (x != xPos || y != yPos || z != zPos) && g.get(x, y, z) {
					activeNeighbours++
				}
			}
		}
	}
	selfActive := g.get(xPos, yPos, zPos)
	if selfActive {
		return activeNeighbours == 2 || activeNeighbours == 3
	}
	return activeNeighbours == 3
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	pocketDimension := grid{}
	// Three vectors
	scanCube := [][]int{
		{-1, 1}, // X - from, to
		{-1, 1}, // Y - from, to
		{-1, 1}, // Z - from, to
	}
	for x, line := range items {
		for y, value := range line {
			pocketDimension.set(x, y, 0, value == '#')
			scanCube[1][1]++
		}
		scanCube[0][1]++
	}

	for i := 0; i < 6; i++ {
		newDimentionState := grid{}
		// Iterate through the rounds
		for x := scanCube[0][0]; x <= scanCube[1][1]; x++ {
			for y := scanCube[1][0]; y <= scanCube[1][1]; y++ {
				for z := scanCube[2][0]; z <= scanCube[2][1]; z++ {
					newDimentionState.set(x, y, z, pocketDimension.determineNewStateAt(x, y, z))
				}
			}
		}
		pocketDimension = newDimentionState
		// Grow the scan cube by 1 in each direction
		scanCube[0][0]--
		scanCube[1][0]--
		scanCube[2][0]--
		scanCube[0][1]++
		scanCube[1][1]++
		scanCube[2][1]++
		// Finally, count the active cubes
		var sumCubes uint
		for _, yGrid := range pocketDimension {
			for _, zGrid := range yGrid {
				for _, val := range zGrid {
					if val {
						sumCubes++
					}
				}
			}
		}
		fmt.Printf("Iteration #%d: Number of active cubes: %d\n", i, sumCubes)
	}

}
