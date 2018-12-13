package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

func willReact(a, b byte) bool {
	// Large ASCII letters start @ 97; Large and small letters have a distance of 32
	return (a >= 97 && a-32 == b) || (a < 97 && a+32 == b)
}

// react only runs one reaction at a time - starting from the beginning of the string
func react(poly string) string {
	for i := 0; i < len(poly)-1; i++ {
		if willReact(poly[i], poly[i+1]) {
			if i > 0 {
				return poly[:i] + poly[i+2:]
			}
			return poly[i+2:]
		}
	}
	return poly
}

func fullReact(poly string) string {
	for {
		oldLen := len(poly)
		poly = react(poly)
		if oldLen == len(poly) {
			// No reaction possible any more
			break
		}
	}
	return poly
}

func exec(poly string) (string, string) {
	shortest := ""
	shortestRepl := ""
	for i := byte(65); i <= 90; i++ {
		reg := regexp.MustCompile(fmt.Sprintf("[%s,%s]", string(i), string(i+32)))
		str := fullReact(reg.ReplaceAllString(poly, ""))
		if shortest == "" || len(str) < len(shortest) {
			shortestRepl = fmt.Sprintf("%s/%s", string(i), string(i+32))
			shortest = str
		}
	}
	return shortest, shortestRepl
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	finalPolymer, removedChars := exec(strings.TrimSpace(string(input)))
	fmt.Printf(
		"Shortest polymer composition is obtained by removing %s:\n\n%s\n\nComposition length is %d units\n",
		removedChars,
		finalPolymer,
		len(finalPolymer),
	)
}
