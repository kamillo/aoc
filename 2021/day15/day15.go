package main

import (
	"fmt"

	"github.com/fzipp/astar"
	"github.com/kamillo/aoc/utils"
)

type graph map[astar.Node][]astar.Node

func main() {
	g := newGraph()
	risk := [][]int{}

	for y, line := range utils.GetLines("input.txt") {
		row := make([]int, len(line))
		risk = append(risk, row)

		for x, c := range utils.ToIntArr(line, "") {
			risk[y][x] = c
		}
	}

	findSafePath := func(risk [][]int) int {
		for y := range risk {
			for x := range risk[y] {
				a := utils.NewPointD3D(x, y, risk[y][x])
				if x+1 < len(risk[y]) {
					n1 := utils.NewPointD3D(x+1, y, risk[y][x+1])
					g.link(a, n1)
					g.link(n1, a)
				}

				if y+1 < len(risk) {
					n2 := utils.NewPointD3D(x, y+1, risk[y+1][x])
					g.link(a, n2)
					g.link(n2, a)
				}
			}
		}

		end := len(risk) - 1
		path := astar.FindPath(g, utils.NewPointD3D(0, 0, risk[0][0]), utils.NewPointD3D(end, end, risk[end][end]), nodeDist, nodeDist)

		cost := 0
		for i := 1; i < len(path); i++ {
			c := path[i].(utils.PointD3D)
			cost += c.Z
		}

		return cost
	}

	cave := make([][]int, len(risk)*5)
	for i := range cave {
		cave[i] = make([]int, len(risk[0])*5)
	}
	for y := range risk {
		for x := range risk[y] {
			cave[y][x] = risk[y][x]
		}
	}

	for l := 1; l < 5; l++ {
		for i := 1; i < 5; i++ {
			for y := range risk {
				for x := range risk[y] {
					yCurrent := y + len(risk)*i
					yOffset := len(risk) * (i - 1)

					cave[yCurrent][x] = (cave[y+yOffset][x] + 1)
					if cave[yCurrent][x] > 9 {
						cave[yCurrent][x] = 1
					}
				}
			}
		}
	}

	for i := 0; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for y := range risk {
				for x := range risk[y] {
					xCurrent := x + len(risk)*j
					yCurrent := y + len(risk)*i
					xOffset := len(risk) * (j - 1)

					cave[yCurrent][xCurrent] = (cave[yCurrent][x+xOffset] + 1)
					if cave[yCurrent][xCurrent] > 9 {
						cave[yCurrent][xCurrent] = 1
					}
				}
			}
		}
	}

	fmt.Println("Part 1: ", findSafePath(risk))
	fmt.Println("Part 2: ", findSafePath(cave))
}

func newGraph() graph {
	return make(map[astar.Node][]astar.Node)
}

func (g graph) link(a, b astar.Node) graph {
	g[a] = append(g[a], b)
	return g
}

func (g graph) Neighbours(n astar.Node) []astar.Node {
	return g[n]
}

func nodeDist(a, b astar.Node) float64 {
	q := b.(utils.PointD3D)
	return float64(q.Z)
}
