package main

import (
	"fmt"
	"image"

	"github.com/kamillo/aoc/utils"
)

var directions = map[string]image.Point{
	"N": image.Pt(0, -1),
	"E": image.Pt(1, 0),
	"S": image.Pt(0, 1),
	"W": image.Pt(-1, 0),
}

type grid struct {
	size      image.Point
	corrupted map[image.Point]bool
}

var g grid

func main() {
	g = grid{}
	g.corrupted = map[image.Point]bool{}

	// lines := utils.GetLines("test.txt")
	// limit := 12
	// g.size = image.Pt(7, 7)

	lines := utils.GetLines("input.txt")
	g.size = image.Pt(71, 71)
	limit := 1024

	origin := image.Point(image.Pt(0, 0))
	end := image.Point(image.Pt(g.size.X-1, g.size.Y-1))

	heuristic := func(p image.Point) int {
		return utils.Abs(p.X-end.X) + utils.Abs(p.Y-end.Y)
	}

	stop := func(s image.Point) bool {
		return s == end
	}

	part1, part2 := 0, image.Point{}

	for i, line := range lines {
		x, y := 0, 0
		fmt.Sscanf(line, "%d,%d", &x, &y)
		g.corrupted[image.Pt(x, y)] = true

		if i == limit-1 {
			_, part1 = utils.Astar[image.Point](origin, stop, neighbors, cost, heuristic)
		}

		if i >= limit {
			_, d := utils.Astar[image.Point](origin, stop, neighbors, cost, heuristic)
			if d == 0 {
				part2 = image.Pt(x, y)
				break
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func cost(from, to image.Point) int {
	return utils.Abs(to.X-from.X) + utils.Abs(to.Y-from.Y)
}

func neighbors(path map[image.Point]image.Point, s image.Point) []image.Point {
	res := []image.Point{}
	for _, off := range directions {
		q := s.Add(off)
		if g.isFreeAt(q) && g.isAccessible(path, s, q) {
			res = append(res, q)
		}
	}
	return res
}

func (g grid) isFreeAt(p image.Point) bool {
	return g.isInBounds(p)
}

func (g grid) isInBounds(p image.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < g.size.Y && p.X < g.size.X
}

func (g grid) isAccessible(path map[image.Point]image.Point, p image.Point, q image.Point) bool {
	if _, ok := g.corrupted[q]; ok {
		return false
	}

	return true
}
