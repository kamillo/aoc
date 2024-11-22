package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	containers := utils.GetLinesAsInterface("input.txt")

	part1 := 0
	min := len(containers)
	for _, c := range utils.Combinations(containers, 0) {
		if utils.Sum(c) == 150 {
			part1++
			if len(c) < min {
				min = len(c)
			}
		}
	}

	part2 := 0
	for _, c := range utils.Combinations(containers, min) {
		if utils.Sum(c) == 150 {
			part2++
		}
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
