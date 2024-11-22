package main

import (
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Cave map[image.Point]bool

func main() {
	lines := utils.GetLines("input.txt")

	cave, maxY := parse(lines)
	part1(cave, maxY)
	part2(cave, maxY)
}

func part1(cave Cave, maxY int) {
	start := image.Pt(500, 0)

	offsets := []image.Point{
		image.Pt(0, 1),  // South
		image.Pt(-1, 1), // SW
		image.Pt(1, 1),  // SE
	}

	for i := 0; ; i++ {
		sand := start
		rest := false
		for !rest {
			rest = true
			for _, o := range offsets {
				q := sand.Add(o)
				if cave.isFreeAt(q) {
					sand = q
					rest = false
					break
				}
			}

			if sand.Y > maxY {
				fmt.Println("Part 1:", i)
				return
			}
		}

		cave[sand] = true
	}
}

func part2(cave Cave, maxY int) {
	start := image.Pt(500, 0)

	offsets := []image.Point{
		image.Pt(0, 1),  // South
		image.Pt(-1, 1), // SW
		image.Pt(1, 1),  // SE
	}

	for i := 0; ; i++ {
		sand := start
		rest := false
		for !rest {
			rest = true
			for _, o := range offsets {
				q := sand.Add(o)
				if cave.isFreeAt(q) && q.Y < maxY+2 {
					sand = q
					rest = false
					break
				}
			}

			if rest && sand == start {
				fmt.Println("Part 2:", i+1)
				return
			}
		}

		cave[sand] = true
	}
}

func (g Cave) isFreeAt(p image.Point) bool {
	return !g[p]
}

func print(cave [][]byte) {
	for i := 0; i < 10; i++ {
		for j := 480; j < 520; j++ {
			if cave[i][j] != 0 {
				fmt.Print(string(cave[i][j]))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func parse(lines []string) (Cave, int) {
	cave := Cave{}
	maxY := 0
	for _, line := range lines {
		fields := strings.Split(line, " -> ")

		for i := 1; i < len(fields); i++ {
			split1 := strings.Split(fields[i-1], ",")
			split2 := strings.Split(fields[i], ",")

			startY, _ := strconv.Atoi(split1[1])
			startX, _ := strconv.Atoi(split1[0])

			endY, _ := strconv.Atoi(split2[1])
			endX, _ := strconv.Atoi(split2[0])

			if maxY < int(math.Max(float64(startY), float64(endY))) {
				maxY = int(math.Max(float64(startY), float64(endY)))
			}

			if startX == endX {
				for y := int(math.Min(float64(startY), float64(endY))); y <= int(math.Max(float64(startY), float64(endY))); y++ {
					cave[image.Pt(startX, y)] = true
				}
			} else if startY == endY {
				for x := int(math.Min(float64(startX), float64(endX))); x <= int(math.Max(float64(startX), float64(endX))); x++ {
					cave[image.Pt(x, startY)] = true
				}
			}
		}
	}

	return cave, maxY
}
