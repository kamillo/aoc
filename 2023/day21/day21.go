package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Grid [][]byte

func main() {
	lines := utils.GetLines("input.txt")

	steps := 64
	grid := Grid{}
	start := image.Pt(0, 0)

	for y, line := range lines {
		grid = append(grid, []byte(line))
		if index := strings.Index(line, "S"); index != -1 {
			start = image.Pt(index, y)
		}
	}

	pointsToCheck := map[image.Point]bool{}
	pointsToCheck[start] = true

	grid[start.Y][start.X] = '.'
	for i := 0; i < steps; i++ {
		tmpPoints := map[image.Point]bool{}

		for p := range pointsToCheck {
			points := GetAdj(p, grid, '.')
			for _, point := range points {
				tmpPoints[point] = true
			}
		}

		pointsToCheck = tmpPoints
	}

	fmt.Println("Part 1:", len(pointsToCheck))

	// Part 2
	// grid = 131x131
	// start = 65,65
	// steps = 26501365 = 65 + x * 131 = 65 + 131 * 202300

}

func GetAdj(point image.Point, grid Grid, char byte) (adj []image.Point) {
	x := point.X
	y := point.Y

	if x+1 < len(grid) && grid[x+1][y] == char {
		adj = append(adj, image.Point{x + 1, y})
	}
	if y+1 < len(grid[x]) && grid[x][y+1] == char {
		adj = append(adj, image.Point{x, y + 1})
	}
	if x > 0 && grid[x-1][y] == char {
		adj = append(adj, image.Point{x - 1, y})
	}
	if y > 0 && grid[x][y-1] == char {
		adj = append(adj, image.Point{x, y - 1})
	}

	return adj
}
