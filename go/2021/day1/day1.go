package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	measurements := utils.GetLinesAsInts("input.txt")
	prev := 0
	count := 0

	for _, m := range measurements {
		if prev > 0 && m > prev {
			count++
		}
		prev = m
	}

	fmt.Println("Part 1: ", count)

	count = 0
	prev = 0
	for x := range measurements {
		window := 0

		if x+3 > len(measurements) {
			break
		}

		window = utils.SumInts(measurements[x : x+3])

		if x != 0 && window > prev {
			count++
		}

		prev = window
	}

	fmt.Println("Part 2: ", count)
}
