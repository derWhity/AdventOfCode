package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

var lineRegEx = regexp.MustCompile("([A-Z]) must be finished before step ([A-Z])")

type node struct {
	executed     bool
	name         string
	requirements []*node
}

type nodeMap map[string]*node

//-- Node methods ------------------------------------------------------------------------------------------------------

func (n *node) String() string {
	return fmt.Sprintf("%s(%t)", n.name, n.executed)
}

func (n *node) IsExecutable() bool {
	if n.executed {
		return false
	}
	for _, nd := range n.requirements {
		if !nd.executed {
			return false
		}
	}
	return true
}

//-- NodeMap methods ---------------------------------------------------------------------------------------------------

func (nm nodeMap) Load(lines []string) {
	for _, line := range lines {
		matches := lineRegEx.FindStringSubmatch(line)
		before := nm.Get(matches[1])
		after := nm.Get(matches[2])
		after.requirements = append(after.requirements, before)
	}
}

func (nm nodeMap) Get(name string) *node {
	nd, ok := nm[name]
	if !ok {
		nd = &node{
			name:         name,
			requirements: make([]*node, 0),
		}
		nm[name] = nd
		fmt.Printf("New node: %s\n", name)
	}
	return nd
}

// Next returns the next node that needs to be executed
func (nm nodeMap) Next() string {
	var ret *node
	for _, nd := range nm {
		if (ret == nil || nd.name[0] < ret.name[0]) && nd.IsExecutable() {
			ret = nd
		}
	}
	if ret != nil {
		return ret.name
	}
	return ""
}

//-- Misc funcs --------------------------------------------------------------------------------------------------------

func exec(input string) string {
	fmt.Println("Starting...")
	lines := strings.Split(input, "\n")
	nMap := make(nodeMap)
	nMap.Load(lines)
	out := ""
	next := nMap.Next()
	for next != "" {
		out += next
		nMap[next].executed = true
		next = nMap.Next()
	}
	return out
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	order := exec(strings.TrimSpace(string(input)))
	fmt.Printf(
		"The ordered steps are:\n%s\n",
		order,
	)
}
