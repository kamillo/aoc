package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strconv"
	"strings"
)

type keyvalue map[string]string

func main() {
	lines := utils.GetLines("input.txt")
	requiredPassportFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	newlinesCount := 0
	for _, line := range lines {
		if len(line) == 0 {
			newlinesCount++
		}
	}

	passports := make([]keyvalue, newlinesCount+1)
	index := 0
	for _, line := range lines {
		if len(line) == 0 {
			index++
			continue
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			if passports[index] == nil {
				passports[index] = make(keyvalue)
			}
			passports[index][strings.Split(field, ":")[0]] = strings.Split(field, ":")[1]
		}
	}

	part1 := 0
	part2 := 0
	for _, passport := range passports {
		valid := true
		for _, key := range requiredPassportFields {
			if _, ok := passport[key]; !ok {
				valid = false
				break
			}
		}
		if valid {
			part1++
			if validatePassport(passport) {
				part2++
			}
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func validatePassport(passport keyvalue) (valid bool) {
	colors := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	valid = true
	for key, value := range passport {
		switch key {
		case "byr":
			valid = validateNumberRange(value, 1920, 2002)
			break
		case "iyr":
			valid = validateNumberRange(value, 2010, 2020)
			break
		case "eyr":
			valid = validateNumberRange(value, 2020, 2030)
			break
		case "hgt":
			valid = false
			if strings.HasSuffix(value, "cm") {
				valid = validateNumberRange(value[:len(value)-2], 150, 193)
			} else if strings.HasSuffix(value, "in") {
				valid = validateNumberRange(value[:len(value)-2], 59, 76)
			}
			break
		case "hcl":
			_, err := strconv.ParseUint(value[1:], 16, 64)
			if !strings.HasPrefix(value, "#") || len(value[1:]) != 6 || err != nil {
				valid = false
			}
			break
		case "ecl":
			if !colors[value] {
				valid = false
			}
			break
		case "pid":
			_, err := strconv.Atoi(value)
			if err != nil || len(value) != 9 {
				valid = false
			}
			break
		case "cid":
			valid = true
			break
		}
	}
	return valid
}

func validateNumberRange(value string, min, max int) bool {
	number, err := strconv.Atoi(value)
	return err == nil && number >= min && number <= max
}
