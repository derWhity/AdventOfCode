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

func printField(data [][]int) {
	for _, line := range data {
		for _, val := range line {
			switch {
			case val == 0: // Just flashed
				fmt.Print(aurora.BrightGreen(fmt.Sprintf("%2d ", val)).String())
			case val < 0: // Border
				fmt.Print(aurora.Gray(5, "🁢🁢🁢"))
			case val == 9:
				fmt.Print(aurora.BrightYellow(fmt.Sprintf("%2d ", val)).String())
			case val == 10:
				fmt.Print(aurora.BrightCyan(fmt.Sprintf("%2d ", val)).String())
			case val > 10:
				fmt.Print(aurora.BrightRed(fmt.Sprintf("%2d ", val)).String())
			default:
				fmt.Printf("%2d ", val)
			}
		}
		fmt.Println("")
	}
}

func inc(data [][]int) {
	for i, line := range data {
		for j, val := range line {
			data[i][j] = val + 1
		}
	}
}

func flash(data [][]int) int {
	flashed := 0
	for i, line := range data {
		for j, val := range line {
			if val >= 10 {
				data[i+1][j] = data[i+1][j] + 1
				data[i][j+1] = data[i][j+1] + 1
				data[i-1][j] = data[i-1][j] + 1
				data[i][j-1] = data[i][j-1] + 1
				data[i+1][j-1] = data[i+1][j-1] + 1
				data[i+1][j+1] = data[i+1][j+1] + 1
				data[i-1][j-1] = data[i-1][j-1] + 1
				data[i-1][j+1] = data[i-1][j+1] + 1
				data[i][j] = -500
				flashed++
			}
		}
	}
	return flashed
}

func reset(data [][]int) {
	for i, line := range data {
		for j, val := range line {
			switch {
			case val < 0 && val >= -500:
				data[i][j] = 0
			case val < 0 && val < -500:
				data[i][j] = -9999
			}
		}
	}
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	// 12x12 to have one element wide border block around the field
	octopi := make([][]int, 12)
	for i := range octopi {
		octopi[i] = []int{-9999, -9999, -9999, -9999, -9999, -9999, -9999, -9999, -9999, -9999, -9999, -9999}
	}
	// Fill the field
	for i, line := range items {
		for j, c := range line {
			octopi[i+1][j+1] = cast.ToInt(string(c))
		}
	}
	printField(octopi)
	numFlashed := 0
	for i := 0; i < 100; i++ {
		fmt.Printf("== Round %3d =====================\n", i)
		inc(octopi)
		for {
			flashed := flash(octopi)
			numFlashed += flashed
			if flashed == 0 {
				break
			}
		}
		reset(octopi)
		printField(octopi)
		fmt.Printf("Flashed: %d\n", numFlashed)
	}
}
