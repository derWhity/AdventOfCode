package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var claimRegEx = regexp.MustCompile(`^#([0-9]+)\s*@\s*([0-9]+),([0-9]+):\s*([0-9]+)x([0-9]+).*`)

type claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
}

func unmarshalClaim(str string) claim {
	if match := claimRegEx.FindStringSubmatch(str); match != nil {
		id, _ := strconv.ParseInt(match[1], 10, strconv.IntSize)
		x, _ := strconv.ParseInt(match[2], 10, strconv.IntSize)
		y, _ := strconv.ParseInt(match[3], 10, strconv.IntSize)
		w, _ := strconv.ParseInt(match[4], 10, strconv.IntSize)
		h, _ := strconv.ParseInt(match[5], 10, strconv.IntSize)
		return claim{
			ID:     int(id),
			X:      int(x),
			Y:      int(y),
			Width:  int(w),
			Height: int(h),
		}
	}
	panic("Unparseable line")
}

func exec(input string) int {
	// X, Y, count requested - only full inches counting here
	var fabric = make(map[int]map[int]int)
	// Process the claims by the elves
	claims := strings.Split(input, "\n")
	for _, claimStr := range claims {
		claim := unmarshalClaim(claimStr)
		fmt.Printf("Claim #%d\n", claim.ID)
		// X coords
		for i := 0; i < claim.Width; i++ {
			x := claim.X + i
			if _, ok := fabric[x]; !ok {
				// Create a new column
				fabric[x] = make(map[int]int)
			}
			// Y coords
			for j := 0; j < claim.Height; j++ {
				// Write the usage
				y := claim.Y + j
				cnt, _ := fabric[x][y]
				fabric[x][y] = cnt + 1
			}
		}
	}
	counter := 0
	// Count the square inches occupied by more than one claim
	for _, col := range fabric {
		for _, rect := range col {
			if rect > 1 {
				counter++
			}
		}
	}
	return counter
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	checksum := exec(string(input))
	fmt.Printf("Multiply claimed inches: %d\n", checksum)
}
