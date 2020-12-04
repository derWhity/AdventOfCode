package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/derWhity/AdventOfCode/lib/input"
)

type validationFn func(string) bool

func validateIntLimits(min int64, max int64) validationFn {
	return func(value string) bool {
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return false
		}
		return intVal >= min && intVal <= max
	}
}

func validateHeight(value string) bool {
	reg := regexp.MustCompile(`^([0-9]+)(in|cm)$`)
	matches := reg.FindStringSubmatch(value)
	fmt.Printf("'%s': Match: %#v\n", value, matches)
	if len(matches) != 3 {
		return false
	}
	if matches[2] == "cm" {
		return validateIntLimits(150, 193)(matches[1])
	} else if matches[2] == "in" {
		return validateIntLimits(59, 76)(matches[1])
	}
	return false
}

func validateRegEx(regex string) validationFn {
	reg := regexp.MustCompile(regex)
	return func(value string) bool {
		return reg.MatchString(value)
	}
}

var requiredFields = map[string]validationFn{
	"byr": validateIntLimits(1920, 2002),
	"iyr": validateIntLimits(2010, 2020),
	"eyr": validateIntLimits(2020, 2030),
	"hgt": validateHeight,
	"hcl": validateRegEx(`^#[0-9a-f]{6}$`),
	"ecl": validateRegEx(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
	"pid": validateRegEx(`^[0-9]{9}$`),
	// "cid", - not needed since optional
}

func fieldsValid(passport map[string]string) bool {
	for key, fn := range requiredFields {
		value, ok := passport[key]
		if !ok {
			fmt.Printf("Field %s not existing", key)
			return false
		}
		if !fn(value) {
			fmt.Printf("Failed test for %s\n", key)
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
			if fieldsValid(passport) {
				fmt.Println("All fields are valid")
				validPassports++
			}
			// Reset
			passport = map[string]string{}
			fmt.Println("")
		}

	}
	fmt.Printf("Valid passports: %d\n", validPassports)
}
