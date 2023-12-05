package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

type Node struct {
	x     int
	y     int
	size  int
	used  int
	avail int
}

type Grid [][]Node

func main() {
	lines := utils.GetLines("input.txt")

	nodes := make(Grid, 40)
	for n := range nodes {
		nodes[n] = make([]Node, 40)
	}

	nodesMap := []Node{}
	maxX := 0
	maxY := 0
	for _, line := range lines {
		x, y, size, used, avail, perc := 0, 0, 0, 0, 0, 0

		if _, err := fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT %d", &x, &y, &size, &used, &avail, &perc); err == nil {
			n := Node{x, y, size, used, avail}
			nodes[y][x] = n
			nodesMap = append(nodesMap, n)

			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	fmt.Println(maxX, maxY)

	sum := 0
	for _, n1 := range nodesMap {
		for _, n2 := range nodesMap {
			if n1.used != 0 && n1 != n2 && n1.used < n2.avail {
				sum++
			}
		}
	}

	fmt.Println(sum)

	nodes.print()

	// 26 + 1 + 33 * 5
}

func (grid *Grid) print() {
	for y := range *grid {
		for x := range *grid {
			fmt.Printf("%3d/%3d,", (*grid)[y][x].used, (*grid)[y][x].size)
		}
		fmt.Println()
	}
}
