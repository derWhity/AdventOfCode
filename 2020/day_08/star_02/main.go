package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
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

func execute(instructions []*instruction, changePos int) (int64, error) {
	var accumulator int64
	var position int
	visitedLines := map[int]bool{}
	for {
		if _, ok := visitedLines[position]; ok {
			return 0, fmt.Errorf("Line %d visited twice", position)
		}
		if position > len(instructions) {
			// Abnormal termination - we're too far after the end
			return 0, fmt.Errorf("Jumped outside of range: %d of %d", position, len(instructions))
		}
		if position == len(instructions) {
			// Index after the last instruction - terminate normally
			return accumulator, nil
		}
		ins := instructions[position]
		visitedLines[position] = true
		// Perform a switch?
		if ins.Operation == "jmp" || ins.Operation == "nop" {
			if changePos == 0 {
				// Do it
				if ins.Operation == "jmp" {
					ins.Operation = "nop"
				} else {
					ins.Operation = "jmp"
				}
			}
			changePos-- // Will get below zero when we do the change - never executing the change again
		}
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
}

func copy(instructions []*instruction) []*instruction {
	out := []*instruction{}
	for _, ins := range instructions {
		insCpy := *ins
		out = append(out, &insCpy)
	}
	return out
}

func main() {
	inputLines, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	instructions := []*instruction{}
	for _, line := range inputLines {
		ins := readLine(line)
		instructions = append(instructions, &ins)
	}
	for changePos := 0; changePos < len(instructions); changePos++ {
		fmt.Printf("Iteration %d: ", changePos)
		accumulator, err := execute(copy(instructions), changePos)
		if err != nil {
			fmt.Printf("[Error] %s\n", err.Error())
		} else {
			fmt.Printf("[Success] Accumulator: %d\n", accumulator)
			return
		}
	}
}
