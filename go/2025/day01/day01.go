package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	part1 := 0
	part2 := 0

	currentValue := 50
	for _, line := range lines {
		rotation := ""
		value := 0
		fmt.Sscanf(string(line), "%1s%d", &rotation, &value)

		if rotation == "L" {
			value = -1 * value
		}

		temp := currentValue + value
		if value > 0 {
			part2 += (temp - 1) / 100
		} else {
			startFloor := floorDiv(currentValue-1, 100)
			endFloor := floorDiv(temp, 100)
			part2 += startFloor - endFloor
		}

		if temp%100 == 0 {
			part2++
		}

		currentValue = utils.ModWrap(temp, 100)

		if currentValue == 0 {
			part1 += 1
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func floorDiv(a, b int) int {
	if a >= 0 {
		return a / b
	}
	return (a - b + 1) / b
}
