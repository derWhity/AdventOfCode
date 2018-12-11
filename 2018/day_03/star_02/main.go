package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/dhconnelly/rtreego"
)

var claimRegEx = regexp.MustCompile(`^#([0-9]+)\s*@\s*([0-9]+),([0-9]+):\s*([0-9]+)x([0-9]+).*`)

type claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
}

func (c *claim) Bounds() *rtreego.Rect {
	rect, _ := rtreego.NewRect(
		rtreego.Point{float64(c.X), float64(c.Y)},
		[]float64{float64(c.Width), float64(c.Height)},
	)
	return rect
}

func unmarshalClaim(str string) *claim {
	if match := claimRegEx.FindStringSubmatch(str); match != nil {
		id, _ := strconv.ParseInt(match[1], 10, strconv.IntSize)
		x, _ := strconv.ParseInt(match[2], 10, strconv.IntSize)
		y, _ := strconv.ParseInt(match[3], 10, strconv.IntSize)
		w, _ := strconv.ParseInt(match[4], 10, strconv.IntSize)
		h, _ := strconv.ParseInt(match[5], 10, strconv.IntSize)
		return &claim{
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
	claims := strings.Split(input, "\n")
	// We'll use an R-tree this time, since it makes everything easier
	var fabric = rtreego.NewTree(2, 10, len(claims))
	// Process the claims by the elves
	claimoObjs := make([]*claim, 0)
	for _, claimStr := range claims {
		claim := unmarshalClaim(claimStr)
		claimoObjs = append(claimoObjs, claim)
		fabric.Insert(claim)
	}
	// Check all the claims and return the first without intersections
	for _, c := range claimoObjs {
		if len(fabric.SearchIntersect(c.Bounds())) == 1 {
			return c.ID
		}
	}
	return 0
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	checksum := exec(string(input))
	fmt.Printf("Non intersecting claim found: %d\n", checksum)
}
