package main

import (
	"fmt"
	"path/filepath"

	"github.com/derWhity/AdventOfCode/lib/input"
)

// grid represents a 3-dimensional grid of cube states
type grid map[int]map[int]map[int]map[int]bool

func (g grid) set(x, y, z, w int, val bool) {
	yGrid, ok := g[x]
	if !ok {
		yGrid = map[int]map[int]map[int]bool{}
		g[x] = yGrid
	}
	zGrid, ok := yGrid[y]
	if !ok {
		zGrid = map[int]map[int]bool{}
		yGrid[y] = zGrid
	}
	wGrid, ok := zGrid[z]
	if !ok {
		wGrid = map[int]bool{}
		zGrid[z] = wGrid
	}
	wGrid[w] = val
}

func (g grid) get(x, y, z, w int) bool {
	yGrid, ok := g[x]
	if !ok {
		return false
	}
	zGrid, ok := yGrid[y]
	if !ok {
		return false
	}
	wGrid, ok := zGrid[z]
	if !ok {
		return false
	}
	return wGrid[w]
}

// determineNewState calculates the new state the given cube will have
func (g grid) determineNewStateAt(xPos, yPos, zPos, wPos int) bool {
	var activeNeighbours uint
	for x := xPos - 1; x <= xPos+1; x++ {
		for y := yPos - 1; y <= yPos+1; y++ {
			for z := zPos - 1; z <= zPos+1; z++ {
				for w := wPos - 1; w <= wPos+1; w++ {
					if (x != xPos || y != yPos || z != zPos || w != wPos) && g.get(x, y, z, w) {
						activeNeighbours++
					}
				}
			}
		}
	}
	selfActive := g.get(xPos, yPos, zPos, wPos)
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
		{-1, 1}, // W - from, to
	}
	for x, line := range items {
		for y, value := range line {
			pocketDimension.set(x, y, 0, 0, value == '#')
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
					for w := scanCube[3][0]; w <= scanCube[3][1]; w++ {
						newDimentionState.set(x, y, z, w, pocketDimension.determineNewStateAt(x, y, z, w))
					}
				}
			}
		}
		pocketDimension = newDimentionState
		// Grow the scan cube by 1 in each direction
		scanCube[0][0]--
		scanCube[1][0]--
		scanCube[2][0]--
		scanCube[3][0]--
		scanCube[0][1]++
		scanCube[1][1]++
		scanCube[2][1]++
		scanCube[3][1]++
		// Finally, count the active cubes
		var sumCubes uint
		for _, yGrid := range pocketDimension {
			for _, zGrid := range yGrid {
				for _, wGrid := range zGrid {
					for _, val := range wGrid {
						if val {
							sumCubes++
						}
					}
				}
			}
		}
		fmt.Printf("Iteration #%d: Number of active cubes: %d\n", i, sumCubes)
	}

}
