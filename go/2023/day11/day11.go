package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	universe := [][]rune{}
	galaxies := []utils.Point2D{}

	for y, line := range lines {
		universe = append(universe, []rune(line))

		for x, cell := range line {
			if cell == '#' {
				galaxies = append(galaxies, utils.Point2D{Y: y, X: x})
			}
		}
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += distance(galaxies[i], galaxies[j], universe, 2)
		}
	}

	fmt.Println("Part 1:", sum)

	sum = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += distance(galaxies[i], galaxies[j], universe, 1000000)
		}
	}

	fmt.Println("Part 2:", sum)
}

func columnContains(universe [][]rune, col int, char rune) bool {
	for _, row := range universe {
		if row[col] == char {
			return true
		}
	}
	return false
}

func rowContains(universe [][]rune, row int, char rune) bool {
	for _, cell := range universe[row] {
		if cell == char {
			return true
		}
	}
	return false
}

func distance(p1, p2 utils.Point2D, universe [][]rune, enlarge int) int {
	distance := utils.Abs(p2.X-p1.X) + utils.Abs(p2.Y-p1.Y)

	for i := min(p1.Y, p2.Y) + 1; i < max(p1.Y, p2.Y); i++ {
		if !rowContains(universe, i, '#') {
			distance += (enlarge - 1)
		}
	}

	for j := min(p1.X, p2.X) + 1; j < max(p1.X, p2.X); j++ {
		if !columnContains(universe, j, '#') {
			distance += (enlarge - 1)
		}
	}

	return distance
}
