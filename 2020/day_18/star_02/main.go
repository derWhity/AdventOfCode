package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

var operatorOrder = []string{"+", "*"}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func isOperator(token string) bool {
	for _, op := range operatorOrder {
		if token == op {
			return true
		}
	}
	return false
}

func indexOf(arr []string, val string) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

func hasHigherOrEqualPrecedence(op1, op2 string) bool {
	return indexOf(operatorOrder, op1) <= indexOf(operatorOrder, op2)
}

func toUPN(parts []string) []string {
	var out []string
	var stack []string
	for _, token := range parts {
		if isOperator(token) {
			for {
				if len(stack) == 0 {
					break
				}
				top := stack[len(stack)-1]
				if !isOperator(top) ||
					top == "(" ||
					!hasHigherOrEqualPrecedence(top, token) {
					break
				}
				out = append(out, top)
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for {
				top := stack[len(stack)-1] // Okay since we don't expect wrong formulas
				stack = stack[:len(stack)-1]
				if top == "(" {
					break
				} else {
					out = append(out, top)
				}
			}
		} else {
			// Number in our case
			out = append(out, token)
		}
	}
	// Read the stack onto the output
	for i := len(stack) - 1; i >= 0; i-- {
		out = append(out, stack[i])
	}
	return out
}

func calc(ops []string) int {
	var stack []int
	var val, left, right int
	var err error
	for _, op := range ops {
		if isOperator(op) {
			right = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch op {
			case "+":
				val = left + right
			case "*":
				val = left * right
			}
		} else {
			val, err = strconv.Atoi(op)
			failOnError(err)
		}
		stack = append(stack, val)
	}
	return stack[len(stack)-1]
}

func main() {
	items, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	var sum int
	for _, line := range items {
		// Split the string into single operations
		var ops []string
		for _, block := range strings.Split(line, " ") {
			for _, item := range strings.Split(block, "(") {
				if item == "" {
					ops = append(ops, "(")
				} else {
					for _, subItem := range strings.Split(item, ")") {
						if subItem == "" {
							ops = append(ops, ")")
						} else {
							ops = append(ops, subItem)
						}
					}
				}
			}
		}
		ops = toUPN(ops)
		result := calc(ops)
		fmt.Printf("Result: %d\n", result)
		sum += result
	}
	fmt.Printf("Sum is: %d\n", sum)
}
