package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func findFreq(input string) int64 {
	parts := strings.Split(input, "\n")
	var freq int64
	foundFreqs := make(map[int64]bool)
	for {
		for _, part := range parts {
			item := strings.TrimSpace(part)
			delta, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				panic(err)
			}
			freq = freq + delta
			if _, ok := foundFreqs[freq]; ok {
				return freq
			}
			foundFreqs[freq] = true
		}
	}
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	freq := findFreq(string(input))
	fmt.Printf("Final frequency: %d\n", freq)
}
