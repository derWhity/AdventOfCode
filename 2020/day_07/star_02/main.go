package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

type bag struct {
	Color    string
	Children map[string]*bagEntry
}

type bagEntry struct {
	Bag   *bag
	Count uint64
}

var (
	bags     = map[string]*bag{}
	splitReg = regexp.MustCompile(`(\d+) ([a-z ]+) bags?`)
)

func newBag(color string) *bag {
	return &bag{
		Color:    color,
		Children: map[string]*bagEntry{},
	}
}

func readLine(line string) {
	parts := strings.Split(line, "bags contain")
	color := strings.TrimSpace(parts[0])
	b, ok := bags[color]
	if !ok {
		b = newBag(color)
		bags[color] = b
	}

	if strings.TrimSpace(parts[1]) == "no other bags." {
		// No children
		return
	}
	parts = strings.Split(parts[1], ",")
	for _, part := range parts {
		res := splitReg.FindStringSubmatch(part)
		entry := bagEntry{}
		var err error
		entry.Count, err = strconv.ParseUint(res[1], 10, 64)
		if err != nil {
			panic(err)
		}
		color = res[2]
		bChild, ok := bags[color]
		if !ok {
			bChild = newBag(color)
			bags[color] = bChild
		}
		entry.Bag = bChild
		b.Children[color] = &entry
	}
}

func countChildren(b *bag) uint64 {
	var count uint64
	count++
	for _, child := range b.Children {
		count += child.Count * countChildren(child.Bag)
	}
	return count
}

func main() {
	inputLines, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	for _, line := range inputLines {
		readLine(line)
	}
	count := countChildren(bags["shiny gold"])
	fmt.Printf("Count: %d\n", count-1) // Don't count the outermost bag

}
