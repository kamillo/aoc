package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	sum1 := 0
	sum2 := 0
	for _, line := range lines {
		values := utils.ToIntArr(line, " ")
		n, p := findNextValue(values)

		sum1 += n + values[len(values)-1]
		sum2 += values[0] - p
	}

	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}

func findNextValue(values []int) (int, int) {
	isZero := true
	next := []int{}
	diff := 0

	for i, v := range values {
		if i > 0 {
			diff = v - values[i-1]
			isZero = isZero && diff == 0
			next = append(next, diff)
		}
	}

	if isZero {
		return 0, 0

	} else {
		nextValue, prevValue := findNextValue(next)

		return nextValue + next[len(next)-1], next[0] - prevValue
	}
}
