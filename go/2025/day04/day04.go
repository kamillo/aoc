package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	grid := utils.GetLinesAs2dArray("input.txt")

	part1, _ := countRolls(grid)
	part2 := 0

	for {
		count := 0
		count, grid = countRolls(grid)
		if count == 0 {
			break
		}

		part2 += count
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func countRolls(grid [][]byte) (int, [][]byte) {
	count := 0
	newGrid := make([][]byte, len(grid))
	copy(newGrid, grid)
	for x := range grid {
		newGrid[x] = make([]byte, len(grid[x]))
		copy(newGrid[x], grid[x])
		for y := range grid[x] {
			if grid[x][y] == '@' {
				if utils.CountAdj(x, y, grid, '@') < 4 {
					count++
					newGrid[x][y] = '.'
				}
			}
		}
	}
	return count, newGrid
}
