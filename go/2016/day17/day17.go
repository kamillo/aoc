package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"math"
	"strings"

	"github.com/kamillo/aoc/2016/day17/astar"
)

func main() {
	pascode := "edjrjqaa"
	grid := graph{
		"######",
		"#    #",
		"#    #",
		"#    #",
		"#     ",
		"#### V",
	}

	_, s := astar.FindShortestPath[image.Point](grid, image.Pt(1, 1), image.Pt(4, 4), nodeDist, nodeDist, pascode, calcPass)
	fmt.Println("Part 1:", strings.ReplaceAll(s, pascode, ""))

	paths := astar.FindLongestPath[image.Point](grid, image.Pt(1, 1), image.Pt(4, 4), nodeDist, nodeDist, pascode, calcPass)
	max := 0
	for _, p := range paths {
		if len(p) > max {
			max = len(p)
		}
	}

	fmt.Println("Part 2:", max-len(pascode))
}

func calcPass(c string, a, b image.Point) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(c)))[:4]
}

func nodeDist(p, q image.Point) float64 {
	return math.Abs(float64(p.X)-float64(q.X)) + math.Abs(float64(p.Y)-float64(q.Y))
}

func (g graph) Neighbours(p image.Point, context string) map[string]image.Point {
	mdfive := fmt.Sprintf("%x", md5.Sum([]byte(context)))[:4]

	offsets := map[string]image.Point{
		"U": image.Pt(0, -1), // North
		"R": image.Pt(1, 0),  // East
		"D": image.Pt(0, 1),  // South
		"L": image.Pt(-1, 0), // West
	}
	res := make(map[string]image.Point)
	for k, off := range offsets {
		q := p.Add(off)
		if g.isFreeAt(q, mdfive, k) {
			res[k] = q
		}
	}
	return res
}

func (g graph) isFreeAt(p image.Point, context string, direction string) bool {
	return g.isInBounds(p) && g[p.Y][p.X] != '#' && !g.isLocked(p, context, direction)
}

func (g graph) isInBounds(p image.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(g) && p.X < len(g[p.Y])
}

func (g graph) isLocked(p image.Point, context string, direction string) bool {
	i := -1
	switch direction {
	case "U":
		i = 0
	case "D":
		i = 1
	case "L":
		i = 2
	case "R":
		i = 3
	}

	return context[i] == 'a' || (context[i] >= '0' && context[i] <= '9')
}

type graph []string
