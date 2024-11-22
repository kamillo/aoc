package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/kamillo/aoc/utils"
)

type Cost struct {
	pos  image.Point
	cost int
}

type Grid map[image.Point]byte
type Graph map[image.Point][]Cost

func main() {
	lines := utils.GetLines("input.txt")

	grid := make(Grid)
	for y, line := range lines {
		for x, c := range line {
			grid[image.Point{X: x, Y: y}] = byte(c)
		}
	}

	start := image.Point{X: 1, Y: 0}
	end := image.Point{X: len(lines[0]) - 2, Y: len(lines) - 1}

	neighbors := buildGraph(grid, start, false)
	visited := make(map[image.Point]bool)
	path := explore(neighbors, start, end, visited, 0, 0)
	fmt.Println("Part 1:", path)

	neighbors = buildGraph(grid, start, true)
	visited = make(map[image.Point]bool)
	path = explore(neighbors, start, end, visited, 0, 0)
	fmt.Println("Part 2:", path)
}

func neighbors(p image.Point) []image.Point {
	return []image.Point{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
	}
}

func exploreSinglePath(grid Grid, previous image.Point, current image.Point, cost int, part2 bool) (Cost, bool) {
	if isAllowed(grid, current) {
		cpt := 0

		for _, ne := range neighbors(current) {
			if isAllowed(grid, ne) {
				cpt++
			}
		}

		if cpt > 2 {
			return Cost{pos: current, cost: cost}, true
		}
	}

	if !part2 {
		if c, ok := grid[current]; ok && c != '.' {
			if current.X > previous.X && c != '>' ||
				current.X < previous.X && c != '<' ||
				current.Y > previous.Y && c != 'v' ||
				current.Y < previous.Y && c != '^' {
				return Cost{}, false
			}
		}
	}

	for _, n := range neighbors(current) {
		if c, ok := grid[n]; ok && c != '#' && n != previous {
			return exploreSinglePath(grid, current, n, cost+1, part2)
		}
	}

	return Cost{pos: current, cost: cost}, true
}

func explore(neighbors Graph, p, goal image.Point, visited map[image.Point]bool, cost int, maxCost int) int {
	if p == goal {
		if cost > maxCost {
			maxCost = cost
		}
		return maxCost
	}

	visited[p] = true
	for _, pc := range neighbors[p] {
		if !visited[pc.pos] {
			maxCost = explore(neighbors, pc.pos, goal, visited, cost+pc.cost, maxCost)
		}
	}
	visited[p] = false
	return maxCost
}

func buildGraph(grid Grid, start image.Point, part2 bool) Graph {
	res := make(map[image.Point][]Cost)
	queue := []image.Point{}
	queue = append(queue, start)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if !isAllowed(grid, p) {
			continue
		}

		for _, n := range neighbors(p) {
			if !isAllowed(grid, n) {
				continue
			}

			pc, ok := exploreSinglePath(grid, p, n, 1, part2)
			if ok && !slices.Contains(res[p], pc) {
				res[p] = append(res[p], pc)
				queue = append(queue, pc.pos)
			}
		}
	}

	return res
}

func isAllowed(grid Grid, p image.Point) bool {
	if c, ok := grid[p]; ok && c != '#' {
		return true
	}
	return false
}
