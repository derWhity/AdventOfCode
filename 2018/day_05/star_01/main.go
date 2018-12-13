package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
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

func exec(poly string) string {
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

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	finalPolymer := exec(strings.TrimSpace(string(input)))
	fmt.Printf("Final polymer composition is:\n\n%s\n\nComposition length is %d units\n", finalPolymer, len(finalPolymer))
}
