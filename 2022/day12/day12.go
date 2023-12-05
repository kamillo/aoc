package main

import (
	"fmt"
	"image"
	"math"
	"strings"

	"github.com/fzipp/astar"
	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	{
		area := graph{}
		S := image.Pt(0, 0)
		E := image.Pt(0, 0)

		for y, line := range lines {
			if strings.Contains(line, "S") {
				S.X = strings.Index(line, "S")
				S.Y = y
				line = strings.Replace(line, "S", "a", 1)
			}

			if strings.Contains(line, "E") {
				E.X = strings.Index(line, "E")
				E.Y = y
				line = strings.Replace(line, "E", "z", 1)
			}

			area = append(area, line)
		}

		n := len(astar.FindPath[image.Point](area, S, E, nodeDist, nodeDist))

		fmt.Println("Part 1:", n-1)
	}

	{
		area := graph{}
		A := []image.Point{}
		E := image.Pt(0, 0)

		for y, line := range lines {
			for x, c := range line {
				if c == 'a' || c == 'S' {
					A = append(A, image.Pt(x, y))
				}

				line = strings.Replace(line, "S", "a", 1)
			}

			if strings.Contains(line, "E") {
				E.X = strings.Index(line, "E")
				E.Y = y
				line = strings.Replace(line, "E", "z", 1)
			}

			area = append(area, line)
		}

		min := math.MaxInt
		for _, s := range A {
			n := len(astar.FindPath[image.Point](area, s, E, nodeDist, nodeDist))
			if n < min && n > 0 {
				min = n
			}
		}

		fmt.Println("Part 2:", min-1)
	}
}

func nodeDist(p, q image.Point) float64 {
	return math.Abs(float64(p.X)-float64(q.X)) + math.Abs(float64(p.Y)-float64(q.Y))
}

func (g graph) Neighbours(p image.Point) []image.Point {
	offsets := []image.Point{
		image.Pt(0, -1), // North
		image.Pt(1, 0),  // East
		image.Pt(0, 1),  // South
		image.Pt(-1, 0), // West
	}
	res := make([]image.Point, 0, 4)
	for _, off := range offsets {
		q := p.Add(off)
		if g.isFreeAt(q) && g.IsAccessible(p, q) {
			res = append(res, q)
		}
	}
	return res
}

func (g graph) isFreeAt(p image.Point) bool {
	return g.isInBounds(p)
}

func (g graph) isInBounds(p image.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(g) && p.X < len(g[p.Y])
}

func (g graph) IsAccessible(p image.Point, q image.Point) bool {
	p1 := g[p.Y][p.X]
	q1 := g[q.Y][q.X]

	return p1+1 == q1 || p1 >= q1
}

type graph []string
