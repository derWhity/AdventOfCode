package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

func calc(input uint64) uint64 {
	res := int64(math.Floor(float64(input)/3)) - 2
	if res <= 0 {
		return 0
	}
	return uint64(res)
}

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
			var res uint64
			for i := calc(uint64(weight)); i > 0; i = calc(i) {
				fmt.Printf("I: %d\n", i)
				res += i
			}
			sum += res
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
