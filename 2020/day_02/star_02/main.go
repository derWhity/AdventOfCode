package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/derWhity/AdventOfCode/lib/input"
)

const (
	// FIRSTPOS is the index of the first character to match (1-indexed)
	FIRSTPOS = 1
	// SECONDPOS is the index of the second character to match (1-indexed)
	SECONDPOS = 2
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

func getCharAt(s string, i int) byte {
	i--
	// len() is number of bytes, but since we're only looking at a-z, this is okay
	if i >= 0 && i < len(s) {
		return s[i]
	}
	// Just return a rune that will never be in the password
	return byte('\n')
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	var numCorrectPasswords uint
	for _, item := range items {
		res := splitRegExp.FindStringSubmatch(item)
		first, err := strconv.ParseInt(res[FIRSTPOS], 10, 64)
		failOnError(err)
		second, err := strconv.ParseInt(res[SECONDPOS], 10, 64)
		failOnError(err)
		var numHits uint64
		if getCharAt(res[PASSWORD], int(first)) == byte(res[CHARACTER][0]) {
			numHits++
		}
		if getCharAt(res[PASSWORD], int(second)) == byte(res[CHARACTER][0]) {
			numHits++
		}
		metRequirements := numHits == 1
		if metRequirements {
			numCorrectPasswords++
		}
	}
	fmt.Printf("Number of valid passwords: %d\n", numCorrectPasswords)
}
