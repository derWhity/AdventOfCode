package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	var initRounds int
	var num string
	var lastNumber uint64
	lastSpoken := map[uint64]int{}
	// Init
	for initRounds, num = range strings.Split(items[0], ",") {
		if initRounds != 0 {
			// Only note after spoken
			lastSpoken[lastNumber] = initRounds
		}
		lastNumber, err = strconv.ParseUint(num, 10, 64)
		failOnError(err)
	}
	fmt.Printf("LastSpoken after init: %+v\n", lastSpoken)
	// Let's play the game
	var i int
	for i = initRounds + 2; i <= 30000000; i++ {
		//fmt.Printf("-- Round: %d --- Last Number: %d -------------------------------------------------\n", i, lastNumber)
		last, ok := lastSpoken[lastNumber]
		if !ok {
			// Never spoken
			lastSpoken[lastNumber] = i - 1
			lastNumber = 0
		} else {
			// Last spoken in the previous round
			num := uint64(i - 1 - last)
			lastSpoken[lastNumber] = i - 1
			lastNumber = num
		}
		//fmt.Printf("Round %d - Elf says: %d\n", i, lastNumber)
		//fmt.Printf("LastSpoken after round: %+v\n", lastSpoken)
	}
	fmt.Printf("Round %d - Elf says: %d\n", i, lastNumber)
}
