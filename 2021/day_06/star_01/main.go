package main

import (
	"fmt"
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

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	parts := strings.Split(items[0], ",")
	fishes := []uint8{}
	for _, str := range parts {
		fishes = append(fishes, cast.ToUint8(str))
	}
	fmt.Printf("Initial fishes: %v\n", fishes)
	for i := 0; i < 80; i++ {
		toAppend := []uint8{}
		for j, fishTimer := range fishes {
			if fishTimer == 0 {
				// New lanternfish
				toAppend = append(toAppend, 8)
				fishes[j] = 6
			} else {
				fishes[j]--
			}
		}
		// Add the new fishies
		fishes = append(fishes, toAppend...)
		fmt.Printf("Fishes after day %2d: %d \n", i+1, len(fishes))
	}
}
