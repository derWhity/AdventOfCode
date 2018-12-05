package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func exec(input string) int {
	parts := strings.Split(input, "\n")
	var threes int
	var twos int
	for _, part := range parts {
		runeCounter := make(map[rune]int)
		part := strings.TrimSpace(part)
		for _, rn := range part {
			c, _ := runeCounter[rn]
			runeCounter[rn] = c + 1
		}
		var twosDone, threesDone bool
		for _, cnt := range runeCounter {
			if !twosDone && cnt == 2 {
				twos++
				twosDone = true
			}
			if !threesDone && cnt == 3 {
				threes++
				threesDone = true
			}
			if threesDone && twosDone {
				break
			}
		}
	}
	return threes * twos
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	checksum := exec(string(input))
	fmt.Printf("Final checksum: %d\n", checksum)
}
