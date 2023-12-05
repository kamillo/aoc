package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	containers := utils.GetLinesAsInterface("input.txt")

	sum := 0
	for _, c := range utils.Combinations(containers) {
		if utils.Sum(c) == 150 {
			sum++
		}
	}

	fmt.Println("Part 1: ", sum)
}
