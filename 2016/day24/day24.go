package main

import (
	"fmt"
	"image"
	"math"
	"strconv"

	"github.com/fzipp/astar"
	"github.com/kamillo/aoc/utils"
)

func main() {
	area := graph{}
	points := []interface{}{}
	paths := map[string]int{}
	start := image.Point{}

	for y, line := range utils.GetLines("input.txt") {
		area = append(area, line)

		for x, c := range line {
			if c > '0' && c <= '9' {
				points = append(points, image.Point{x, y})
			}

			if c == '0' {
				start = image.Pt(x, y)
			}
		}
	}

	perm := utils.HeapPermutation(points)
	findRoute := func(endAtZero bool) int {
		min := math.MaxInt64
		for _, pe := range perm {
			route := 0

			pe = append([]interface{}{start}, pe...)
			if endAtZero {
				pe = append(pe, start)
			}

			for p := 0; p < len(pe)-1; p++ {
				n, exist := paths[fmt.Sprintf("%v %v", pe[p], pe[p+1])]
				if !exist {
					n = len(astar.FindPath[image.Point](area, pe[p].(image.Point), pe[p+1].(image.Point), nodeDist, nodeDist)) - 1
					paths[fmt.Sprintf("%v %v", pe[p], pe[p+1])] = n
				}

				//fmt.Printf("%v -> %v %d ", pe[p], pe[p+1], n)
				route += n
			}

			//fmt.Println(route, pe)

			if min > route {
				min = route
				// fmt.Println(min, pe)
			}
		}

		return min
	}

	fmt.Println("Part 1:", findRoute(false))
	fmt.Println("Part 2:", findRoute(true))
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
		if g.isFreeAt(q) {
			res = append(res, q)
		}
	}
	return res
}

func (g graph) isFreeAt(p image.Point) bool {
	_, err := strconv.Atoi(string(g[p.Y][p.X]))
	return g.isInBounds(p) && (g[p.Y][p.X] == '.' || err == nil)
}

func (g graph) isInBounds(p image.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(g) && p.X < len(g[p.Y])
}

func (g graph) put(p image.Point, c rune) {
	g[p.Y] = g[p.Y][:p.X] + string(c) + g[p.Y][p.X+1:]
}

func (g graph) print() {
	for _, row := range g {
		fmt.Println(row)
	}
}

type graph []string
