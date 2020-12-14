package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

var (
	indexRegex = regexp.MustCompile(`^mem\[([0-9]+)\]$`)
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lines, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	mem := map[int]uint64{} // Our memory block
	andMask := ^uint64(0)   // Let all pass
	orMask := uint64(0)     // Leave all untouched
	for _, line := range lines {
		cmdSplit := strings.Split(line, " = ")
		if cmdSplit[0] == "mask" {
			// Set the mask
			tmp, err := strconv.ParseUint(strings.ReplaceAll(cmdSplit[1], "X", "0"), 2, 64)
			failOnError(err)
			orMask = tmp
			tmp, err = strconv.ParseUint(strings.ReplaceAll(cmdSplit[1], "X", "1"), 2, 64)
			failOnError(err)
			andMask = tmp
			fmt.Printf("&mask: %b\n|mask: %b\n", andMask, orMask)
		} else {
			// Write to mem
			value, err := strconv.ParseUint(cmdSplit[1], 10, 64)
			failOnError(err)
			res := indexRegex.FindStringSubmatch(cmdSplit[0])
			var pos int64
			pos, err = strconv.ParseInt(res[1], 10, 64)
			failOnError(err)
			value = value & andMask
			value = value | orMask
			mem[int(pos)] = value
			fmt.Printf("Written at [%d]: %b (%d)\n", pos, value, value)
		}
	}
	// All written - now for the sum
	var sum uint64
	for _, val := range mem {
		sum += val
	}
	fmt.Printf("Sum of all memory values: %d\n", sum)
}
