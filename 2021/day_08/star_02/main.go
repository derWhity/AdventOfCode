package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
	"github.com/spf13/cast"
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
	rest := []string{}
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
		default:
			rest = append(rest, t)
		}
	}
	s.tests = rest
	//fmt.Printf("Rest: %+v\n", rest)
	rest = []string{}
	// Find 9
	for _, t := range s.tests {
		test := toBinary(t)
		if len(t) == 6 && (test&s.mapping[4] == s.mapping[4]) {
			// Found the 9
			s.mapping[9] = test
		} else {
			rest = append(rest, t)
		}
	}
	s.tests = rest
	//fmt.Printf("Rest: %+v\n", rest)
	rest = []string{}
	// Find 0 and 6
	for _, t := range s.tests {
		test := toBinary(t)
		if len(t) == 6 {
			if test&s.mapping[7] == s.mapping[7] {
				// Found the 0
				s.mapping[0] = test
			} else {
				// Found the 6
				s.mapping[6] = test
			}
		} else {
			rest = append(rest, t)
		}
	}
	s.tests = rest
	//fmt.Printf("Rest: %+v\n", rest)
	// Find the rest - all are len() = 5 now
	fiveCheck := s.mapping[1] ^ s.mapping[4]
	for _, t := range s.tests {
		test := toBinary(t)
		if test&s.mapping[1] == s.mapping[1] {
			// 3
			s.mapping[3] = test
		} else if test&fiveCheck == fiveCheck {
			// 5
			s.mapping[5] = test
		} else {
			s.mapping[2] = test
		}
	}
	// fmt.Println("Computed mapping:")
	// for i, val := range s.mapping {
	// 	fmt.Printf("[%d] => %b\n", i, val)
	// }
}

func (s *segmentDisplay) decode(val string) int {
	binVal := toBinary(val)
	for i, binMap := range s.mapping {
		test := binVal & binMap
		if test == binVal && test == binMap {
			return i
		}
	}
	return -1
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	var sum int
	for _, line := range items {
		parts := strings.Split(line, "|")
		tests := strings.Split(strings.TrimSpace(parts[0]), " ")
		values := strings.Split(strings.TrimSpace(parts[1]), " ")
		display := newDisplay(tests)
		display.compute()
		var num int
		for i, val := range values {
			decoded := display.decode(val)
			num += int(math.Pow(10, float64(3-i))) * decoded
		}
		sum += cast.ToInt(num)
		fmt.Printf("%d\n", num)
	}
	fmt.Printf("Sim is: %d\n", sum)
}
