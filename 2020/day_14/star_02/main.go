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

func computeMasks(input string) []uint64 {
	var out []uint64
	var max uint64
	for _, c := range input {
		if c == 'X' {
			max = max << 1
			max++
		}
	}
	fmt.Printf("Max: %d (%b)\n", max, max)
	for i := uint64(0); i <= max; i++ {
		tmp := fmt.Sprintf("%%0%db", len(fmt.Sprintf("%b", max)))
		fmt.Printf("Tmp: %s\n", tmp)
		val := fmt.Sprintf(tmp, i)
		fmt.Printf("Val: %s\n", val)
		var str string
		for _, c := range input {
			if c == 'X' {
				if len(val) > 0 {
					str += string(val[0]) // Safe because we're not outside the ASCII range
					val = val[1:]
				} else {
					str += "0"
				}
			} else {
				str += string(c)
			}
		}
		fmt.Printf("Replaced: %s\n", str)
		mask, err := strconv.ParseUint(str, 2, 64)
		failOnError(err)
		out = append(out, mask)
	}
	return out
}

func main() {
	lines, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	mem := map[uint64]uint64{} // Our memory block
	var masks []uint64
	var andMask uint64
	for _, line := range lines {
		cmdSplit := strings.Split(line, " = ")
		if cmdSplit[0] == "mask" {
			// Set the mask
			masks = computeMasks(cmdSplit[1])
			tmp, err := strconv.ParseUint(
				strings.ReplaceAll(strings.ReplaceAll(cmdSplit[1], "0", "1"), "X", "0"),
				2,
				64,
			)
			failOnError(err)
			andMask = tmp
			fmt.Println("Masks:")
			for _, m := range masks {
				fmt.Printf("> %06b (%d)\n", m, m)
			}
		} else {
			// Write to mem
			value, err := strconv.ParseUint(cmdSplit[1], 10, 64)
			failOnError(err)
			res := indexRegex.FindStringSubmatch(cmdSplit[0])
			var pos uint64
			pos, err = strconv.ParseUint(res[1], 10, 64)
			failOnError(err)
			for _, m := range masks {
				position := pos&andMask | m
				//fmt.Printf("Mask: %06b\n      %06b\n      %06b\n====> %06b (%d)\n", pos, andMask, m, position, position)
				mem[position] = value
				//fmt.Printf("Written at [%d / %b]: %b (%d)\n", position, position, value, value)
			}
		}
	}
	// All written - now for the sum
	var sum uint64
	for _, val := range mem {
		sum += val
	}
	fmt.Printf("Sum of all memory values: %d\n", sum)
}
