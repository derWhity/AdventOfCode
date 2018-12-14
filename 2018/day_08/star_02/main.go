package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type node struct {
	children []*node
	meta     []int
}

func (n *node) walk() int {
	var ret int
	if len(n.children) == 0 {
		// Metadata sum
		for _, meta := range n.meta {
			ret += meta
		}
	} else {
		// Sum of the child node refs
		for _, meta := range n.meta {
			if len(n.children) >= meta {
				ret += n.children[meta-1].walk()
			}
		}
	}
	return ret
}

//-- Misc funcs --------------------------------------------------------------------------------------------------------

func buildTree(numbers []string) (*node, []string) {
	nd := &node{}
	// We'll ditch error handling for now
	numChilden, _ := strconv.ParseInt(numbers[0], 10, strconv.IntSize)
	numMeta, _ := strconv.ParseInt(numbers[1], 10, strconv.IntSize)
	numbers = numbers[2:]
	// Load the children
	for i := 0; i < int(numChilden); i++ {
		var child *node
		child, numbers = buildTree(numbers)
		nd.children = append(nd.children, child)
	}
	// Load the metadata
	for i := 0; i < int(numMeta); i++ {
		meta, _ := strconv.ParseInt(numbers[0], 10, strconv.IntSize)
		numbers = numbers[1:]
		nd.meta = append(nd.meta, int(meta))
	}
	return nd, numbers
}

func exec(input string) int {
	numbers := strings.Split(input, " ")
	root, _ := buildTree(numbers) // We don't need the remainder slice here (it's empty nonetheless)
	return root.walk()
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	sum := exec(strings.TrimSpace(string(input)))
	fmt.Printf(
		"Node sum: %d\n",
		sum,
	)
}
