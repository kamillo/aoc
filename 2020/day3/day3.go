package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
)

type slope struct {
	right int
	down  int
}

func main() {
	lines := utils.GetLines("input.txt")

	fmt.Println("Part 1: ", countTrees(lines, slope{3, 1}))

	// Part 2
	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treesProduct := 1
	for _, slope := range slopes {
		treesProduct *= countTrees(lines, slope)
	}

	fmt.Println("Part 2: ", treesProduct)
}

func countTrees(grid []string, slope slope) (treesCount int) {
	x, y := 0, 0
	width := len(grid[0])

	for y < len(grid) {
		if grid[y][x%width] == '#' {
			treesCount++
		}
		y += slope.down
		x += slope.right
	}

	return
}
