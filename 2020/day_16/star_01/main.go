package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

const (
	modeRules  = 0
	modeOwn    = 1
	modeNearby = 2
)

type validRange struct {
	lower uint64
	upper uint64
}

type ranges []validRange

// valid checks if the given number matches any of the valid ranges
func (r ranges) valid(num uint64) bool {
	for _, currentRange := range r {
		if num <= currentRange.upper && num >= currentRange.lower {
			return true
		}
	}
	return false
}

type ruleSet map[string]ranges

// validRules determines the names of the field rule candidates that may be valid for the given number
func (r ruleSet) validRules(num uint64) []string {
	var out []string
	for name, r := range r {
		if r.valid(num) {
			out = append(out, name)
		}
	}
	return out
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// newTicket reads a ticket from the given string and returns it as list of fields
func newTicket(input string) []uint64 {
	var out []uint64
	for _, numStr := range strings.Split(input, ",") {
		num, err := strconv.ParseUint(numStr, 10, 64)
		failOnError(err)
		out = append(out, num)
	}
	return out
}

// newRanges extracts the ranges for a field name and returns both
func newRanges(input string) (string, ranges) {
	tmp := strings.Split(input, ": ")
	out := ranges{}
	fieldName := tmp[0]
	for _, ele := range strings.Split(tmp[1], " or ") {
		r := validRange{}
		var err error
		values := strings.Split(ele, "-")
		r.lower, err = strconv.ParseUint(values[0], 10, 64)
		failOnError(err)
		r.upper, err = strconv.ParseUint(values[1], 10, 64)
		failOnError(err)
		out = append(out, r)
	}
	return fieldName, out
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	mode := modeRules
	var sumErrorValues uint64
	rules := ruleSet{}
	for i, line := range items {
		if line != "" {
			switch mode {
			case modeRules:
				if line == "your ticket:" {
					fmt.Println("Mode change to own ticket")
					mode = modeOwn
					continue
				}
				// Parse the next rule
				fld, r := newRanges(line)
				rules[fld] = r
			case modeOwn:
				if line == "nearby tickets:" {
					fmt.Println("Mode change to nearby tickets")
					mode = modeNearby
					continue
				}
				// Parse your own ticket - but since it's not needed in this solution, let's do nothing
			case modeNearby:
				ticket := newTicket(line)
				fmt.Printf("Ticket: %+v\n", ticket)
				for j, fieldValue := range ticket {
					if len(rules.validRules(fieldValue)) == 0 {
						// None match - so add this field to the sum
						fmt.Printf("Ticket in line %d, field %d with value %d does not match any criteria\n", i, j, fieldValue)
						sumErrorValues += fieldValue
					}
				}
			}
		}
	}
	fmt.Printf("Ticket scanning error rate is: %d\n", sumErrorValues)
}
