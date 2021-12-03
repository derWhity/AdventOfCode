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

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	var counters []int
	for i, line := range items {
		if i == 0 {
			counters = make([]int, len(line))
		}
		if line != "" {
			for j, char := range strings.Split(line, "") {
				if char == "1" {
					counters[j]++
				}
			}
		}
	}
	var gamma uint64
	var epsilon uint64
	for i, count := range counters {
		if count > ((len(items) - 1) / 2) {
			gamma |= (1 << (len(counters) - i - 1))
		} else {
			epsilon |= (1 << (len(counters) - i - 1))
		}
	}
	fmt.Printf("Gamma:      %10d (%b)\n", gamma, gamma)
	fmt.Printf("Epsilon:    %10d (%b)\n", epsilon, epsilon)
	mul := gamma * epsilon
	fmt.Printf("Multiplied: %10d (%b)\n", mul, mul)
}
