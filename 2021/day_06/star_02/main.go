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

func sumFishes(counts []uint64) uint64 {
	var sum uint64
	for _, val := range counts {
		sum += val
	}
	return sum
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	parts := strings.Split(items[0], ",")
	counts := make([]uint64, 9)
	for _, str := range parts {
		counts[cast.ToInt(str)]++
	}
	fmt.Printf("Initial counts: %v\n", counts)
	for i := 0; i < 256; i++ {
		breeding := counts[0]
		counts = counts[1:]
		counts[6] += breeding
		counts = append(counts, breeding)
		fmt.Printf("Fishes after day %2d: %v (%d)\n", i+1, counts, sumFishes(counts))
	}
}
