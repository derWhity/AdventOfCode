package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

func exec(input string) uint64 {
	parts := strings.Split(input, "\n")
	var sum uint64
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			weight, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				panic(err)
			}
			sum += uint64(math.Floor(float64(weight)/3)) - 2
		}
	}
	return sum
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	mass := exec(string(input))
	fmt.Printf("Fuel requirement: %d\n", mass)
}
