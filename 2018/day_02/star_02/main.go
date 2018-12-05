package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func numDiff(candidate, contestant string) (int, string) {
	diff := 0
	var sameStr string
	for i := 0; i < len(candidate); i++ {
		if candidate[i] != contestant[i] {
			diff++
		} else {
			sameStr += string(candidate[i])
		}
	}
	return diff, sameStr
}

func exec(input string) string {
	parts := strings.Split(input, "\n")
	for _, candidate := range parts {
		candidate = strings.TrimSpace(candidate)
		for _, contestant := range parts {
			contestant = strings.TrimSpace(contestant)
			if num, str := numDiff(candidate, contestant); num == 1 {
				return str
			}
		}
	}
	return ""
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	res := exec(string(input))
	fmt.Printf("Same letters: %s\n", res)
}
