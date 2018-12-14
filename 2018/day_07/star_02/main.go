package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

var lineRegEx = regexp.MustCompile("([A-Z]) must be finished before step ([A-Z])")

const (
	statusWaiting = iota
	statusRunning
	statusDone
	numWorkers = 5
)

type node struct {
	status       int
	name         string
	requirements []*node
	timeLeft     int
}

type nodeMap map[string]*node

//-- Node methods ------------------------------------------------------------------------------------------------------

func (n *node) String() string {
	if n != nil {
		return fmt.Sprintf("%s(%3d)", n.name, n.timeLeft)
	}
	return "  --  "

}

func (n *node) IsExecutable() bool {
	if n.status != statusWaiting {
		return false
	}
	for _, nd := range n.requirements {
		if nd.status != statusDone {
			return false
		}
	}
	return true
}

func (n *node) ExecTime() int {
	return int(n.name[0]) - 4 // 64 chars to get a "1" - since it's 60 + letterNumber, it's 4 here
}

func (n *node) Tick() bool {
	if n.status != statusRunning {
		return true
	}
	n.timeLeft--
	if n.timeLeft == 0 {
		n.status = statusDone
		return true
	}
	return false
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
			status:       statusWaiting,
		}
		nd.timeLeft = nd.ExecTime()
		nm[name] = nd
		fmt.Printf("New node: %s\n", name)
	}
	return nd
}

// Next returns the next node that needs to be executed
func (nm nodeMap) Next() *node {
	var ret *node
	for _, nd := range nm {
		if (ret == nil || nd.name[0] < ret.name[0]) && nd.IsExecutable() {
			ret = nd
		}
	}
	return ret
}

func (nm nodeMap) allDone() bool {
	for _, nd := range nm {
		if nd.status != statusDone {
			return false
		}
	}
	return true
}

//-- Misc funcs --------------------------------------------------------------------------------------------------------

func exec(input string) int {
	fmt.Println("Starting...")
	lines := strings.Split(input, "\n")
	nMap := make(nodeMap)
	nMap.Load(lines)
	// Write a status header
	str := "Second   |"
	for i := 0; i < numWorkers; i++ {
		str += fmt.Sprintf("    %2d    |", i)
	}
	str += fmt.Sprintf(" Done %s|", strings.Repeat(" ", len(nMap)-4))
	fmt.Println(str)
	fmt.Println(strings.Repeat("-", len(str)))
	// Let's go
	var second int
	var done string
	workers := make([]*node, numWorkers)
	for {
		// Continue working
		for i, nd := range workers {
			if nd != nil && nd.Tick() {
				workers[i] = nil
				done += nd.name
			}
		}
		// Fill the idle workers with new things to do
		for i, nd := range workers {
			if nd == nil {
				next := nMap.Next()
				if next == nil {
					break
				}
				next.status = statusRunning
				workers[i] = next
			}
		}
		// Display the status
		str := fmt.Sprintf("%8d |", second)
		for _, nd := range workers {
			str += fmt.Sprintf("  %s  |", nd)
		}
		fmt.Printf("%s %s\n", str, done)
		if nMap.allDone() {
			return second
		}
		second++
	}
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	secs := exec(strings.TrimSpace(string(input)))
	fmt.Printf(
		"Seconds taken: %d\n",
		secs,
	)
}
