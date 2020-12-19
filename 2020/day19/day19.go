package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"regexp"
	"strings"
)

func main() {
	rules := map[string]string{}
	cases := []string{}
	lines := fileutil.GetLines("input.txt")
	endRules := false

	for _, line := range lines {
		if len(line) == 0 {
			endRules = true
		}

		if endRules {
			cases = append(cases, line)
		} else {
			rules[strings.Split(line, ": ")[0]] = strings.Split(line, ": ")[1]
		}
	}

	resolved := map[string]string{}

	// part1
	p1 := 0
	countMatches(rules, resolved, func() bool {
		_, ok := resolved["0"]
		return ok
	})

	for _, c := range cases {
		if match, _ := regexp.MatchString("\\b"+resolved["0"]+"\\b", c); match { // \b - match whole word
			//fmt.Println(resolved["0"], c)
			p1 += 1
		}
	}

	fmt.Println("Part 1: ", p1)

	// part 2
	rules["8"] = "42 | 42 8"
	rules["11"] = "42 31 | 42 11 31"
	p2 := map[string]bool{}
	countMatches(rules, resolved, func() bool {
		added := false
		for _, c := range cases {
			if match, _ := regexp.MatchString("\\b"+resolved["0"]+"\\b", c); match { // \b - match whole word
				if ok := p2[c]; !ok {
					p2[c] = true
					added = true
				}
			}
		}

		return !added
	})
	fmt.Println("Part 2: ", len(p2))
}

func countMatches(rules map[string]string, resolved map[string]string, breakCondition func() bool) {
	for {
		for i, rule := range rules {
			if rule[0] == '"' {
				resolved[i] = string(rule[1])
			} else {
				parts := strings.Split(rule, " ")
				allFound := true
				for _, part := range parts {
					if part == "|" {
						continue
					}

					if _, ok := resolved[part]; !ok {
						allFound = false
					}
				}

				expr := ""
				if allFound {
					for _, part := range parts {
						if part == "|" {
							expr += "|"
						} else {
							expr = expr + resolved[part]
						}
					}
					resolved[i] = "(" + expr + ")"
				}
			}
		}
		if breakCondition() {
			break
		}
	}
}
