package main

import (
	"fmt"
	"image"
	"math"

	"github.com/kamillo/aoc/utils"
)

type graph [][]int
type node struct {
	p         image.Point
	direction string
	strike    int
}

var directions = map[string]image.Point{
	"N": image.Pt(0, -1),
	"E": image.Pt(1, 0),
	"S": image.Pt(0, 1),
	"W": image.Pt(-1, 0),
}

var g = graph{}

func main() {
	lines := utils.GetLines("input.txt")

	for _, line := range lines {
		g = append(g, utils.ToIntArr(line, ""))
	}

	origin := image.Pt(0, 0)
	starts := []node{{p: origin, direction: "N", strike: 0}}

	path, distance := utils.AstarMultipleStart(starts, stop, neighbors, cost, heuristic)

	printPath(g, path)

	fmt.Println("Part 1:", distance)
}

func heuristic(p node) int {
	q := node{p: image.Pt(len(g)-1, len(g[0])-1)}
	return int(math.Abs(float64(p.p.X-q.p.X)) + math.Abs(float64(p.p.Y-q.p.Y)))
}

func stop(s node) bool {
	return s.p == image.Pt(len(g)-1, len(g[0])-1)
}

func cost(from, to node) int {
	return g[to.p.Y][to.p.X]
}

func neighbors(path map[node]node, s node) []node {
	res := []node{}
	for dir, off := range directions {
		q := s.p.Add(off)
		if g.isFreeAt(q) && g.isAccessible(path, s, q) {
			strike := 0
			if dir == s.direction {
				strike = s.strike + 1
			}
			res = append(res, node{p: q, direction: dir, strike: strike})
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

func (g graph) isAccessible(path map[node]node, p node, q image.Point) bool {
	if len(path) >= 2 {
		reverseDirection := getDirection(p.p, path[p].p)
		if reverseDirection == getDirection(p.p, q) {
			return false
		}
	}

	if len(path) >= 4 {
		p3 := path[p]
		p2 := path[p3]
		p1 := path[p2]
		d1 := getDirection(p1.p, p2.p)
		d2 := getDirection(p2.p, p3.p)
		d3 := getDirection(p3.p, p.p)
		d4 := getDirection(p.p, q)

		if len(map[string]bool{d1: true, d2: true, d3: true, d4: true}) == 1 {
			return false
		}
	}

	return true
}

func getDirection(prev, next image.Point) string {
	for dir, off := range directions {
		if prev.Add(off) == next {
			return dir
		}
	}
	return ""
}

func printPath(g graph, path []node) {
	// var dirs = map[string]{'^', '>', 'v', '<'}
	for _, s := range path {
		g[s.p.Y][s.p.X] = 0 //dirs[s.direction]
	}
	for _, l := range g {
		fmt.Println(l)
	}
}
