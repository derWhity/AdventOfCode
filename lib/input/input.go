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

func ReadNonEmptyString(filename string, trim bool) ([]string, error) {
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
		if item != "" {
			out = append(out, item)
		}
	}
	return out, nil
}

// Reads blocks of string lines - a block is separated by two newlines
func ReadBlocks(filename string, trim bool) ([][]string, error) {
	input, err := ioutil.ReadFile(filepath.Join(filename))
	if err != nil {
		return nil, err
	}
	blocks := strings.Split(string(input), "\n\n")
	var out [][]string
	for _, item := range blocks {
		var block []string
		for _, line := range strings.Split(item, "\n") {
			if trim {
				line = strings.TrimSpace(line)
			}
			if item != "" {
				block = append(block, line)
			}
		}
		out = append(out, block)
	}
	return out, nil
}

// ReadInt reads the puzzle input as list of integers (int64)
func ReadInt(filename string) ([]int64, error) {
	var out []int64
	items, err := ReadString(filename, true)
	if err != nil {
		return nil, err
	}
	for i, item := range items {
		if item != "" {
			convertedItem, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse line %d: %w", i+1, err)
			}
			out = append(out, convertedItem)
		}
	}
	return out, nil
}

// ReadIntNative reads the puzzle input as list of integers (int)
func ReadIntNative(filename string) ([]int, error) {
	var out []int
	items, err := ReadString(filename, true)
	if err != nil {
		return nil, err
	}
	for i, item := range items {
		if item != "" {
			convertedItem, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse line %d: %w", i+1, err)
			}
			out = append(out, int(convertedItem))
		}
	}
	return out, nil
}
