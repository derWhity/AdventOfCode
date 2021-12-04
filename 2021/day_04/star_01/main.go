package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cast"
)

var (
	splitRegex = regexp.MustCompile(`\s+`)
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

type entry struct {
	num    int
	marked bool
}

type board struct {
	entries [][]*entry
}

func newBoard() board {
	return board{
		entries: [][]*entry{},
	}
}

func (b *board) addRow(rowStr string) {
	elements := splitRegex.Split(rowStr, -1)
	row := []*entry{}
	for _, str := range elements {
		el := entry{
			num: cast.ToInt(str),
		}
		row = append(row, &el)
	}
	b.entries = append(b.entries, row)
}

func (b *board) numberDrawn(num int) bool {
	elementsInRow := len(b.entries)
	elementsInCol := len(b.entries[0])
	rowCounter := make([]uint, elementsInRow)
	colCounter := make([]uint, elementsInCol)
	for row, rowElements := range b.entries {
		for col, e := range rowElements {
			if e.num == num {
				e.marked = true
			}
			if e.marked {
				rowCounter[row]++
				colCounter[col]++
			}
		}
	}
	for i, sum := range rowCounter {
		if sum == uint(elementsInRow) {
			fmt.Println(aurora.BgRed(aurora.White(fmt.Sprintf("Match in row %d", i))).String())
			return true
		}
	}
	for i, sum := range colCounter {
		if sum == uint(elementsInCol) {
			fmt.Println(aurora.BgRed(aurora.White(fmt.Sprintf("Match in column %d", i))).String())
			return true
		}
	}
	return false
}

func (b *board) String() string {
	var out string
	for _, row := range b.entries {
		for _, col := range row {
			var val string
			if col.marked {
				val = aurora.BrightGreen(fmt.Sprintf("%2d ", col.num)).String()
			} else {
				val = aurora.BrightBlue(fmt.Sprintf("%2d ", col.num)).String()
			}
			out += val
		}
		out += "\n"
	}
	return out
}

func (b *board) getSumOfUnmarkedNumbers() int {
	var sum int
	for _, row := range b.entries {
		for _, e := range row {
			if !e.marked {
				sum += e.num
			}
		}
	}
	return sum
}

func main() {
	items, err := input.ReadBlocks(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	boardDefs := items[1:]
	drawnNumbers := strings.Split(items[0][0], ",")
	fmt.Printf("Drawn numbers: %v\n", drawnNumbers)
	// Create the boards
	boards := []*board{}
	for _, lines := range boardDefs {
		board := newBoard()
		for _, line := range lines {
			if line != "" {
				board.addRow(line)
			}
		}
		boards = append(boards, &board)
		fmt.Println(board.String())
	}
	// Let's play
	for round, numStr := range drawnNumbers {
		num := cast.ToInt(numStr)
		fmt.Printf("-- Round %3d => %2d ------------------\n", round, num)
		var boardWithBingo *board
		for _, b := range boards {
			if res := b.numberDrawn(num); res && boardWithBingo == nil {
				boardWithBingo = b
			}
			fmt.Println(b.String())
		}
		if boardWithBingo != nil {
			score := boardWithBingo.getSumOfUnmarkedNumbers() * num
			fmt.Printf("Bingo!!! - Board has %s points\n", aurora.Bold(aurora.BrightRed(fmt.Sprintf("%d", score))))
			break
		}

	}
}
