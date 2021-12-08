package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func toBinary(str string) uint8 {
	var res uint8
	for _, c := range str {
		res |= (1 << (uint8(c) - 97)) // "a" is ASCII position 97
	}
	return res
}

type segmentDisplay struct {
	tests   []string
	mapping []uint8 // Binary representations of the numbers found - zero value means not found (=all segments off)
}

func newDisplay(t []string) *segmentDisplay {
	var out segmentDisplay
	out.tests = append(out.tests, t...)
	out.mapping = make([]uint8, 10)
	return &out
}

func (s *segmentDisplay) compute() {
	// Find 1, 4, 7, 8
	for _, t := range s.tests {
		switch len(t) {
		case 2:
			// 1
			s.mapping[1] = toBinary(t)
		case 3:
			// 7
			s.mapping[7] = toBinary(t)
		case 4:
			// 4
			s.mapping[4] = toBinary(t)
		case 7:
			// 8
			s.mapping[8] = toBinary(t)
		}
	}
	fmt.Println("Computed mapping:")
	for i, val := range s.mapping {
		fmt.Printf("[%d] => %b\n", i, val)
	}
}

func (s *segmentDisplay) decode(val string) *uint8 {
	binVal := toBinary(val)
	for i, binMap := range s.mapping {
		test := binVal & binMap
		if test == binVal && test == binMap {
			out := uint8(i)
			return &out
		}
	}
	return nil
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	var count uint
	for _, line := range items {
		parts := strings.Split(line, "|")
		tests := strings.Split(strings.TrimSpace(parts[0]), " ")
		values := strings.Split(strings.TrimSpace(parts[1]), " ")
		display := newDisplay(tests)
		display.compute()
		for _, val := range values {
			if decoded := display.decode(val); decoded != nil {
				fmt.Printf("Hit: %s\n", val)
				count++
			}
		}
	}
	fmt.Printf("Count is: %d\n", count)
}
