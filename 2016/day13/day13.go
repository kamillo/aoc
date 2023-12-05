package main

import (
	"fmt"
	"image"
	"math"
	"strings"

	"github.com/fzipp/astar"
)

const input = 1358
const targetX = 31
const targetY = 39

// const input = 10
// const targetX = 7
// const targetY = 4
const buffer = 10

func main() {
	office := graph{}

	for y := 0; y < targetY+buffer; y++ {
		builder := strings.Builder{}
		builder.Grow(targetX + buffer)

		for x := 0; x < targetX+buffer; x++ {
			cell := x*x + 3*x + 2*x*y + y + y*y
			cell += input
			cellType := strings.Count(fmt.Sprintf("%b", cell), "1") % 2

			if cellType == 0 {
				builder.WriteString(" ")
			} else {
				builder.WriteString("#")
			}
		}
		office = append(office, builder.String())
	}

	start := image.Pt(1, 1)
	dest := image.Pt(targetX, targetY)
	path := astar.FindPath[image.Point](office, start, dest, nodeDist, nodeDist)
	// for _, p := range path {
	// 	office.put(p, '.')
	// }
	// office.print()
	fmt.Println("Part 1:", len(path)-1)

	nodes := map[image.Point]bool{}
	office.dfs(50, start, image.Pt(-1, -1), nodes)

	fmt.Println("Part 2:", len(nodes))
}

func nodeDist(p, q image.Point) float64 {
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

type graph []string

// Neighbours implements the astar.Graph[Node] interface (with Node = image.Point).
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
	return g.isInBounds(p) && g[p.Y][p.X] == ' '
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

func (g graph) dfs(k int, node, parent image.Point, nodes map[image.Point]bool) map[image.Point]bool {
	if k < 0 {
		return nodes
	}

	nodes[node] = true

	for _, nb := range g.Neighbours(node) {
		if nb != parent {
			// node nb becomes the parent
			nodes = g.dfs(k-1, nb, node, nodes)
		}
	}

	return nodes
}
