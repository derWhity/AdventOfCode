package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

const (
	// MIN is the index of the minimum occurrence in the regex result
	MIN = 1
	// MAX is the index of the maximum occurrence in the regex result
	MAX = 2
	// CHARACTER is the index of the character to check in the regex result
	CHARACTER = 3
	// PASSWORD is the index of the checked password in the regex result
	PASSWORD = 4
)

var splitRegExp = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	var numCorrectPasswords uint
	for _, item := range items {
		res := splitRegExp.FindStringSubmatch(item)
		min, err := strconv.ParseUint(res[MIN], 10, 64)
		failOnError(err)
		max, err := strconv.ParseUint(res[MAX], 10, 64)
		failOnError(err)
		charCount := strings.Count(res[PASSWORD], res[CHARACTER])
		metRequirements := min <= uint64(charCount) && max >= uint64(charCount)
		if metRequirements {
			numCorrectPasswords++
		}
		fmt.Printf(
			"Password '%s' contains the char '%s' %d times (%d - %d allowed): %t\n",
			res[PASSWORD],
			res[CHARACTER],
			charCount,
			min,
			max,
			metRequirements,
		)
	}
	fmt.Printf("Number of valid passwords: %d\n", numCorrectPasswords)
}
