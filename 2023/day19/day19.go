package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	workflows := map[string][]string{}
	parts := []map[string]int{}
	rulesRegexp := regexp.MustCompile(`^(.*?)\{(.*?)\}$`)
	partsRegexp := regexp.MustCompile(`^\{(.*?)\}$`)
	accepted := []map[string]int{}

	rules := true
	for _, line := range lines {
		if line == "" {
			rules = false
			continue
		}

		if rules {
			matches := rulesRegexp.FindStringSubmatch(line)
			name := matches[1]
			rules := strings.Split(matches[2], ",")
			workflows[name] = rules

		} else {
			matches := partsRegexp.FindStringSubmatch(line)

			ratings := map[string]int{}
			split := strings.Split(matches[1], ",")
			for _, rating := range split {
				r := strings.Split(rating, "=")
				ratings[r[0]] = utils.JustAtoi(r[1])
			}

			parts = append(parts, ratings)
		}
	}

	for _, part := range parts {
		next := "in"
		for next != "R" && next != "A" {
			next = checkRule(workflows[next], part)
			// fmt.Println(part, next)
			if next == "A" {
				accepted = append(accepted, part)
			}
		}
	}

	sum := 0
	for _, part := range accepted {
		fmt.Println(part)
		for _, value := range part {
			sum += value
		}

	}

	fmt.Println("Part 1:", sum)
}

func checkRule(rules []string, parts map[string]int) string {
	for _, rule := range rules {
		splitLess := strings.Split(rule, "<")
		splitMore := strings.Split(rule, ">")

		if len(splitLess) == 2 {
			key := splitLess[0]

			valueStr := strings.Split(splitLess[1], ":")

			value := utils.JustAtoi(valueStr[0])
			rule := valueStr[1]

			if parts[key] < value {
				return rule
			}

		} else if len(splitMore) == 2 {
			key := splitMore[0]

			valueStr := strings.Split(splitMore[1], ":")
			value := utils.JustAtoi(valueStr[0])
			rule := valueStr[1]

			if parts[key] > value {
				return rule
			}

		} else {
			// fallback
			return rule
		}
	}

	return ""
}
