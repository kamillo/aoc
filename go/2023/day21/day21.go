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
	// 26501365 % 131 = 65

	gridLen := len(grid)
	var options []int = make([]int, 65+gridLen*2)

	pointsToCheck = map[image.Point]bool{}
	pointsToCheck[start] = true

	grid[start.Y][start.X] = '.'
	for i := 0; i < 65+gridLen*2; i++ {
		tmpPoints := map[image.Point]bool{}

		for p := range pointsToCheck {
			points := GetAdj(p, grid, '.')
			for _, point := range points {
				tmpPoints[point] = true
			}
		}

		pointsToCheck = tmpPoints
		options[i] = len(pointsToCheck)
	}

	f := func(x int, a [3]int) int {
		b0 := a[0]
		b1 := a[1] - a[0]
		b2 := a[2] - a[1]

		return b0 + (b1 * x) + (x*(x-1)/2)*(b2-b1)
	}

	pt2 := f(26501365/gridLen, [3]int{options[64], options[64+gridLen], options[64+gridLen*2]})

	fmt.Println("Part 2:", pt2)
}

func GetAdj(point image.Point, grid Grid, char byte) (adj []image.Point) {
	x := point.X
	y := point.Y
	n := len(grid)

	if grid[utils.ModWrap(x+1, n)][utils.ModWrap(y, n)] == char {
		adj = append(adj, image.Point{x + 1, y})
	}
	if grid[utils.ModWrap(x, n)][utils.ModWrap(y+1, n)] == char {
		adj = append(adj, image.Point{x, y + 1})
	}
	if grid[utils.ModWrap(x-1, n)][utils.ModWrap(y, n)] == char {
		adj = append(adj, image.Point{x - 1, y})
	}
	if grid[utils.ModWrap(x, n)][utils.ModWrap(y-1, n)] == char {
		adj = append(adj, image.Point{x, y - 1})
	}

	return adj
}
