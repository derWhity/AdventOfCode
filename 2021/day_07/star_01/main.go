package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
	"github.com/spf13/cast"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func median(arr sort.IntSlice) int {
	arr.Sort()
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}
	if len(arr)%2 == 0 {
		return int((arr[len(arr)/2] + arr[(len(arr)/2)-1]) / 2)
	}
	return arr[len(arr)/2]
}

func main() {
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	parts := strings.Split(items[0], ",")
	numbers := sort.IntSlice{}
	for _, str := range parts {
		numbers = append(numbers, cast.ToInt(str))
	}
	med := median(numbers)
	fmt.Printf("Sorted numbers: %v\nMedian: %d\n", numbers, med)
	// Calculate fuel consumption
	var consumption int
	for _, val := range numbers {
		if val > med {
			consumption += val - med
		} else {
			consumption += med - val
		}
	}
	fmt.Printf("Fuel consumption: %d\n", consumption)

}
