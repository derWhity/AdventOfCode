// Package input is a centralized reader for puzzle inputs
package input

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// ReadString reads the puzzle input as list of (optionally trimmed) strings
func ReadString(filename string, trim bool) ([]string, error) {
	input, err := ioutil.ReadFile(filepath.Join(filename))
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(input), "\n")
	var out []string
	for _, item := range parts {
		if trim {
			item = strings.TrimSpace(item)
		}
		out = append(out, item)
	}
	return out, nil
}

// ReadInt reads the puzzle input as list of integers
func ReadInt(filename string) ([]int64, error) {
	var out []int64
	items, err := ReadString(filename, true)
	if err != nil {
		return nil, err
	}
	for i, item := range items {
		convertedItem, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse line %d: %w", i+1, err)
		}
		out = append(out, convertedItem)
	}
	return out, nil
}
