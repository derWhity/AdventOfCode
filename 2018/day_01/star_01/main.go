package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func freqCounter(input string) int64 {
	parts := strings.Split(input, "\n")
	var freq int64
	for _, part := range parts {
		item := strings.TrimSpace(part)
		delta, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			panic(err)
		}
		freq = freq + delta
	}
	return freq
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	freq := freqCounter(string(input))
	fmt.Printf("Final frequency: %d\n", freq)
}
