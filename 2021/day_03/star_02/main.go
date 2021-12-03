package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/derWhity/AdventOfCode/lib/input"
	"github.com/spf13/cast"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func mostCommon(lines []string, pos int) bool {
	var count uint
	for _, line := range lines {
		if line != "" && line[pos] == '1' {
			count++
		}
	}
	return float64(count)/float64(len(lines)) >= 0.5
}

func getRating(lines []string, pos int, leastCommon bool) int64 {
	searchOnes := mostCommon(lines, pos) // We use booleans here instead of 1 and 0
	if leastCommon {
		searchOnes = !searchOnes
	}
	searchChar := cast.ToString(cast.ToInt(searchOnes))
	result := []string{}
	for _, line := range lines {
		if line != "" && line[pos] == searchChar[0] {
			result = append(result, line)
		}
	}
	fmt.Printf("Rest: %v\n", result)
	if len(result) == 1 {
		// We're done
		val, err := strconv.ParseInt(result[0], 2, 64)
		failOnError(err)
		return val
	}
	return getRating(result, pos+1, leastCommon)
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	oxygenGeneratorRating := getRating(items, 0, false)
	fmt.Println("------------------------")
	co2ScrubberRating := getRating(items, 0, true)
	fmt.Printf(
		"O2:  %10d\nCO2: %10d\nLSR: %10d\n",
		oxygenGeneratorRating,
		co2ScrubberRating,
		oxygenGeneratorRating*co2ScrubberRating,
	)
}
