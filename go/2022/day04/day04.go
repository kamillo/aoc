package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	cnt1 := 0
	cnt2 := 0
	for _, line := range lines {
		a, b := 0, 0
		aa, bb := 0, 0

		if _, err := fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &aa, &bb); err != nil {
			panic("scanf")
		}

		if (a <= aa && b >= bb) ||
			(aa <= a && bb >= b) {
			cnt1++
		}

		if (a >= aa && a <= bb) ||
			(b >= aa && b <= bb) ||
			(aa >= a && aa <= b) ||
			(bb >= a && bb <= b) {
			cnt2++
		}
	}

	fmt.Println("Part 1:", cnt1)
	fmt.Println("Part 2:", cnt2)
}
