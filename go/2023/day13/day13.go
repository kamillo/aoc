package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	fmt.Println("Part 1:", calc(lines, equal))
	fmt.Println("Part 2:", calc(lines, isSmudged))
}

func calc(lines []string, compare func(a, b []string) bool) int {
	pattern := []string{}

	sum := 0
	for _, line := range lines {
		if line == "" {
			h := findHorizontalSymmetry(pattern, compare)
			v := findHorizontalSymmetry(rotate(pattern), compare)

			if h > v {
				sum += h * 100
			} else {
				sum += v
			}

			pattern = []string{}
		} else {
			pattern = append(pattern, line)
		}
	}

	return sum
}

func rotate(pattern []string) []string {
	n := len(pattern)
	m := len(pattern[0])
	rotated := make([][]rune, m)
	for i := range rotated {
		rotated[i] = make([]rune, n)
	}
	for i, row := range pattern {
		for j, val := range row {
			rotated[j][n-1-i] = val
		}
	}

	var strings []string
	for _, runeSlice := range rotated {
		strings = append(strings, string(runeSlice))
	}
	return strings
}

func findHorizontalSymmetry(problem []string, compare func(a, b []string) bool) int {
	for i := 1; i < len(problem); i++ {
		rows := min(i, len(problem)-i)
		up := problem[i-rows : i]
		down := reverse(problem[i : i+rows])
		if compare(up, down) {
			return i
		}
	}
	return 0
}

func isSmudged(a, b []string) bool {
	found := false

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				if found {
					return false
				} else {
					found = true
				}
			}
		}
	}

	return found
}

func reverse(s []string) []string {
	a := make([]string, len(s))
	copy(a, s)
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
