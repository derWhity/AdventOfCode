package main

import (
	"fmt"
	"path/filepath"

	"github.com/derWhity/AdventOfCode/lib/input"
)

const (
	floor    = '.'
	empty    = 'L'
	occupied = '#'
	offGrid  = 0
)

type grid map[int]map[int]rune

func (g grid) print() {
	for x := 0; x < len(g); x++ {
		for y := 0; y < len(g[x]); y++ {
			fmt.Printf("%s", string(g[x][y]))
		}
		fmt.Println("")
	}
	fmt.Println("----------")
}

func (g grid) seat(x, y int) rune {
	if row, ok := g[x]; ok {
		return row[y] // zero when index does not exist
	}
	return offGrid
}

// count the first seat found in this direction
func countFirst(input grid, x, y, incX, incY int) uint {
	xPos := x
	yPos := y
	for {
		xPos += incX
		yPos += incY
		s := input.seat(xPos, yPos)
		switch s {
		case offGrid:
			return 0
		case empty:
			return 0
		case occupied:
			return 1
		}
	}
}

func numOccupiedVisible(input grid, x, y int) uint {
	var count uint
	count += countFirst(input, x, y, -1, -1)
	count += countFirst(input, x, y, 0, -1)
	count += countFirst(input, x, y, 1, -1)
	count += countFirst(input, x, y, -1, 0)
	count += countFirst(input, x, y, 1, 0)
	count += countFirst(input, x, y, -1, 1)
	count += countFirst(input, x, y, 0, 1)
	count += countFirst(input, x, y, 1, 1)
	return count
}

func round(input grid) (grid, uint, uint) {
	var numChanged uint
	var numOccupied uint
	out := grid{}
	// The order is random, but for us this should not be a problem in this quiz
	for x, row := range input {
		out[x] = map[int]rune{}
		for y, seat := range row {
			if seat == empty {
				// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
				if numOccupiedVisible(input, x, y) == 0 {
					out[x][y] = occupied
					numChanged++
				} else {
					out[x][y] = seat
				}
			} else if seat == occupied {
				// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
				if numOccupiedVisible(input, x, y) >= 5 {
					out[x][y] = empty
					numChanged++
				} else {
					out[x][y] = seat
				}
			} else {
				// Otherwise, the seat's state does not change.
				out[x][y] = seat
			}
			if out[x][y] == occupied {
				numOccupied++
			}
		}
	}
	return out, numChanged, numOccupied
}

func main() {
	seats := grid{}
	fmt.Println(">> ------------- <<")
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	// Prepare the grid
	for x, line := range items {
		seats[x] = map[int]rune{}
		for y, seat := range line {
			seats[x][y] = seat
		}
	}
	// Run until nothing changes
	var numChanged, numOccupied uint
	seats.print()
	for {
		seats, numChanged, numOccupied = round(seats)
		//seats.print()
		fmt.Printf("Round finished: %d changed | %d occupied\n", numChanged, numOccupied)
		if numChanged == 0 {
			return
		}

	}
}
