package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	// "cid", - not needed since optional
}

func hasAllFields(passport map[string]string) bool {
	for _, requiredField := range requiredFields {
		if _, ok := passport[requiredField]; !ok {
			fmt.Printf("Missing field: %s\n", requiredField)
			return false
		}
	}
	return true
}

func main() {
	inputLines, err := input.ReadString(filepath.Join("..", "input.txt"), true)
	if err != nil {
		panic(err)
	}
	var validPassports uint
	passport := map[string]string{}
	for i, line := range inputLines {
		if line != "" {
			// Line containing fields
			fields := strings.Split(line, " ")
			for _, field := range fields {
				elements := strings.Split(field, ":")
				if len(elements) == 2 {
					passport[strings.TrimSpace(elements[0])] = strings.TrimSpace(elements[1])
				}
			}
		}
		if line == "" || i == len(inputLines)-1 {
			// Passport is finished - check the fields
			fmt.Printf("Passport:\n%#v\n", passport)
			if hasAllFields(passport) {
				fmt.Println("All fields present")
				validPassports++
			}
			// Reset
			passport = map[string]string{}
			fmt.Println("")
		}

	}
	fmt.Printf("Valid passports: %d\n", validPassports)
}
