package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLinesAs2dIntArray("input.txt")

	part1 := 0
	part2 := 0
	for _, line := range lines {
		part1 += getMaxJoltage(line, 2)
		part2 += getMaxJoltage(line, 12)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func getMaxJoltage(bank []int, digits int) int {
	start := 0
	maxJoltage := 0
	for i := range digits {
		max := 0
		end := len(bank) - (digits - (i + 1))
		max, start = getRangedMax(bank, start, end)
		maxJoltage = utils.JustAtoi(fmt.Sprintf("%d%d", maxJoltage, max))
	}

	return maxJoltage
}

func getRangedMax(bank []int, start int, end int) (int, int) {
	max := 0
	pos := 0
	for i := start; i < end; i++ {
		if bank[i] > max {
			max = bank[i]
			pos = i + 1
		}
	}

	return max, pos
}
