package main

import (
	"fmt"
	"math"

	"github.com/kamillo/aoc/utils"
)

type Range struct {
	low  int
	high int
}

func main() {
	lines := utils.GetLines("input.txt")

	ranges := []Range{}
	for _, line := range lines {
		low, high := 0, 0
		fmt.Sscanf(line, "%d-%d", &low, &high)

		ranges = append(ranges, Range{low, high})
	}

	min := 0
	for {
		blocked := false
		for _, r := range ranges {
			if min >= r.low && min <= r.high {
				min = r.high + 1
				blocked = true
			}
		}

		if !blocked {
			break
		}
	}

	fmt.Println("Part 1:", min)

	allowed := 0

	for i := 0; i < math.MaxUint32; i++ {
		blocked := false
		for _, r := range ranges {
			if i >= r.low && i <= r.high {
				i = r.high
				blocked = true
			}
		}

		if !blocked {
			allowed++
		}
	}

	fmt.Println("Part 2:", allowed)

}
