package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	common := []rune{}
	for _, line := range lines {
		half1 := line[:len(line)/2]
		half2 := line[len(line)/2:]

		common = append(common, findCommonItem(half1, half2))
	}

	fmt.Println("Part 1:", calcSum(common))

	badges := []rune{}
	for i := 0; i < len(lines); i += 3 {
		l1 := lines[i]
		l2 := lines[i+1]
		l3 := lines[i+2]

		badges = append(badges, findBadge(l1, l2, l3))
	}

	fmt.Println("Part 2:", calcSum(badges))
}

func findCommonItem(a, b string) rune {
	for _, c1 := range a {
		for _, c2 := range b {
			if c1 == c2 {
				return c1
			}
		}
	}

	return rune(-1)
}

func calcSum(items []rune) int {
	sum := 0
	for _, c1 := range items {
		if c1 >= 'a' && c1 <= 'z' {
			sum += int(c1 - rune('a') + 1)
		}
		if c1 >= 'A' && c1 <= 'Z' {
			sum += int(c1 - rune('A') + 27)
		}
	}

	return sum
}

func findBadge(l1, l2, l3 string) rune {
	for _, c1 := range l1 {
		for _, c2 := range l2 {
			for _, c3 := range l3 {
				if c1 == c2 && c2 == c3 {
					return c1
				}
			}
		}
	}

	return rune(-1)
}
