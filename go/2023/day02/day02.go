package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	sum := 0
	sum2 := 0

	for _, line := range lines {
		id := 0
		max := map[string]int{"red": 0, "green": 0, "blue": 0}

		fmt.Sscanf(line, "Game %d:", &id)

		goodGame := id
		tmp := strings.Split(line, ":")
		sets := strings.Split(tmp[1], ";")

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {
				count := 0
				color := ""
				fmt.Sscanf(cube, "%d %s", &count, &color)

				if !checkLimit(count, color) {
					goodGame = 0
				}

				if count > max[color] {
					max[color] = count
				}
			}
		}

		sum += goodGame
		sum2 += max["red"] * max["green"] * max["blue"]
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}

func checkLimit(count int, color string) bool {
	switch color {
	case "red":
		return count <= 12
	case "green":
		return count <= 13
	case "blue":
		return count <= 14
	}

	return false
}
