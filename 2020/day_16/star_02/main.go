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

func contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

type candidateMap map[string][]int

// update updates the map of possible candidates with the current ticket
func (m candidateMap) update(ticket []uint64, rules ruleSet) {
	indexMap := map[int][]string{}
	for idx, value := range ticket {
		validFields := rules.validRules(value)
		if len(validFields) == 0 {
			// Ignore the whole ticket
			fmt.Printf("Ignoring ticket %+v\n", ticket)
			return
		}
		indexMap[idx] = validFields
	}
	// Now remove each index from the fields not listed in its candidate list
	for idx, fieldList := range indexMap {
		for fieldName, candidates := range m {
			if !contains(fieldList, fieldName) {
				// Remove our index from the candidate list
				var c []int
				for _, index := range candidates {
					if index != idx {
						c = append(c, index)
					}
				}
				m[fieldName] = c
			}
		}
	}
}

func (m candidateMap) print() {
	fmt.Printf("Candidates: %+v\n", m)
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	mode := modeRules
	rules := ruleSet{}
	var ownTicket []uint64
	candidates := candidateMap{}
	for _, line := range items {
		if line != "" {
			switch mode {
			case modeRules:
				if line == "your ticket:" {
					fmt.Println("Mode change to own ticket")
					mode = modeOwn
					// Compile the candidates
					var candidateList []int
					for x := 0; x < len(rules); x++ {
						candidateList = append(candidateList, x)
					}
					for fieldName := range rules {
						candidates[fieldName] = candidateList
					}
					candidates.print()
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
				ownTicket = newTicket(line)
				candidates.update(ownTicket, rules)
			case modeNearby:
				ticket := newTicket(line)
				candidates.update(ticket, rules)
			}
		}
	}
	candidates.print()
	// Now find the fields based on the candidate list
	fields := map[string]int{}
	for {
		if len(candidates) == 0 {
			break
		}
		var found []int
		for fieldName, fieldCandidates := range candidates {
			if len(fieldCandidates) == 1 {
				found = append(found, fieldCandidates[0])
				fields[fieldName] = fieldCandidates[0]
			}
			if len(fieldCandidates) == 0 {
				delete(candidates, fieldName)
			}
		}
		// Remove the found items from all fields
		for _, fInt := range found {
			for fieldName, fieldCandidates := range candidates {
				var newCandidates []int
				for _, val := range fieldCandidates {
					if val != fInt {
						newCandidates = append(newCandidates, val)
					}
				}
				candidates[fieldName] = newCandidates
			}
		}
	}
	fmt.Printf("Fields: %+v\n", fields)
	// Find the field names beginning with "departure"
	product := uint64(1)
	for name, idx := range fields {
		if strings.HasPrefix(name, "departure") {
			fmt.Printf("Field '%s' has value %d\n", name, ownTicket[idx])
			product *= ownTicket[idx]
		}
	}
	fmt.Printf("All multiplied are: %d\n", product)
}
