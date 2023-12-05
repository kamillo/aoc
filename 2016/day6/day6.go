package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	position := [8]map[string]int{}

	for _, line := range lines {
		for i, r := range line {
			if position[i] == nil {
				position[i] = make(map[string]int)
			}

			position[i][string(r)]++
		}
	}

	message := ""
	message2 := ""
	for _, p := range position {
		sorted := utils.SortMapStringInt(p, true)
		message += sorted[0].Key
		message2 += sorted[len(sorted)-1].Key
	}

	fmt.Println("Part 1:", message)
	fmt.Println("Part 2:", message2)
}
