package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
)

func main() {
	theSue := map[string]int{
		"children:":    3,
		"cats:":        7,
		"samoyeds:":    2,
		"pomeranians:": 3,
		"akitas:":      0,
		"vizslas:":     0,
		"goldfish:":    5,
		"trees:":       3,
		"cars:":        2,
		"perfumes:":    1}

	compare := func(s string, v int) bool {
		switch s {
		case "cats:":
			fallthrough
		case "trees:":
			return v > theSue[s]
		case "pomeranians:":
			fallthrough
		case "goldfish:":
			return v < theSue[s]
		default:
			return v == theSue[s]
		}
	}

	for _, line := range utils.GetLines("input.txt") {
		number := 0
		a, b, c, s := "", "", "", ""
		x, y, z := 0, 0, 0

		fmt.Sscanf(line, "%s %d: %s %d, %s %d, %s %d", &s, &number, &a, &x, &b, &y, &c, &z)

		if theSue[a] == x && theSue[b] == y && theSue[c] == z {
			fmt.Println("Part 1: ", number)
		}

		if compare(a, x) && compare(b, y) && compare(c, z) {
			fmt.Println("Part 2: ", number)
		}
	}
}
