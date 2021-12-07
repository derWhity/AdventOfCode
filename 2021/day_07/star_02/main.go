package main

import (
	"fmt"
	"math"
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

func average(arr sort.IntSlice) float64 {
	var sum int
	for _, n := range arr {
		sum += n
	}
	return (float64(sum) / float64(len(arr)))
}

func gaussSum(num int) int {
	return (num * (num + 1)) / 2
}

func main() {
	fmt.Printf("GS(16->5): %d\n", gaussSum(16-5))
	items, err := input.ReadNonEmptyString(filepath.Join("..", "input.txt"), true)
	failOnError(err)
	parts := strings.Split(items[0], ",")
	numbers := sort.IntSlice{}
	for _, str := range parts {
		numbers = append(numbers, cast.ToInt(str))
	}
	fmt.Printf("Numbers length: %d\n", len(numbers))
	avg := average(numbers)
	fmt.Printf("Average: %f\n", avg)
	// Calculate fuel consumption
	var consumption int
	var consumptionRounded int
	for _, val := range numbers {
		var distance int
		var distanceRounded int
		if val > int(avg) {
			distance += val - int(avg)
		} else {
			distance += int(avg) - val

		}
		if val > int(math.Round(avg)) {
			distanceRounded += val - int(math.Round(avg))
		} else {
			distanceRounded += int(math.Round(avg)) - val
		}
		consumption += gaussSum(distance)
		consumptionRounded += gaussSum(distanceRounded)
	}
	fmt.Printf("Consumptions: %d | %d\n", consumption, consumptionRounded)
	if consumption > consumptionRounded {
		consumption = consumptionRounded
	}
	fmt.Printf("Fuel consumption: %d\n", consumption)

}
