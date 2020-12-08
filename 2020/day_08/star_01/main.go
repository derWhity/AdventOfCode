package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

var (
	visitedLines = map[int]bool{}
)

type instruction struct {
	Operation string
	Argument  int64
}

func readLine(line string) instruction {
	parts := strings.Split(line, " ")
	if parts[1][0] == byte('+') {
		parts[1] = parts[1][1:]
	}
	arg, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}
	return instruction{
		Operation: strings.TrimSpace(parts[0]),
		Argument:  arg,
	}
}

func main() {
	inputLines, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	instructions := []instruction{}
	for _, line := range inputLines {
		instructions = append(instructions, readLine(line))
	}
	var accumulator int64
	var position int
	for {
		if _, ok := visitedLines[position]; ok {
			// Already visited
			break
		}
		ins := instructions[position]
		visitedLines[position] = true
		switch ins.Operation {
		case "nop":
			// Nothing to do
			position++
		case "acc":
			// Addition
			accumulator += ins.Argument
			position++
		case "jmp":
			// Jump
			position += int(ins.Argument)
		default:
			panic("Unknown operation " + ins.Operation)
		}
	}
	fmt.Printf("Visited command at %d twice | Accumulator is at %d\n", position, accumulator)
}
