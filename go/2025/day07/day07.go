package main

import (
	"fmt"
	"image"

	"github.com/kamillo/aoc/utils"
)

type Node image.Point

func main() {
	grid := utils.GetLinesAs2dArray("input.txt")

	activeBeams := map[int][]Node{}
	part1 := 0

	var startNode Node
	graph := utils.Graph[Node]{}

	for y := range grid {
		nextActiveBeams := map[int][]Node{}

		for x, char := range grid[y] {
			if char == 'S' {
				startNode = Node{x, y}
				nextActiveBeams[x] = append(nextActiveBeams[x], startNode)
				if _, exists := graph[startNode]; !exists {
					graph[startNode] = []Node{}
				}
			}
		}

		for x, sources := range activeBeams {
			if len(sources) == 0 {
				continue
			}

			hitSplitter := false
			if x >= 0 && x < len(grid[y]) {
				if grid[y][x] == '^' {
					hitSplitter = true
				}
			}

			if hitSplitter {
				currentNode := Node{x, y}
				for _, sourceNode := range sources {
					graph[sourceNode] = append(graph[sourceNode], currentNode)
				}

				if _, exists := graph[currentNode]; !exists {
					graph[currentNode] = []Node{}
				}

				nextActiveBeams[x-1] = append(nextActiveBeams[x-1], currentNode)
				nextActiveBeams[x+1] = append(nextActiveBeams[x+1], currentNode)
				part1++
			} else {
				nextActiveBeams[x] = append(nextActiveBeams[x], sources...)
			}
		}
		activeBeams = nextActiveBeams
	}

	exitY := len(grid)
	for x, sources := range activeBeams {
		exitNode := Node{x, exitY}
		if _, exists := graph[exitNode]; !exists {
			graph[exitNode] = []Node{}
		}
		for _, sourceNode := range sources {
			graph[sourceNode] = append(graph[sourceNode], exitNode)
		}
	}

	fmt.Println("Part 1:", part1)
	// utils.Print2dArray(grid)

	memo := map[Node]int{}
	var countPaths func(n Node) int
	countPaths = func(n Node) int {
		if count, ok := memo[n]; ok {
			return count
		}

		if len(graph[n]) == 0 {
			return 1
		}

		total := 0
		for _, neighbor := range graph[n] {
			total += countPaths(neighbor)
		}

		memo[n] = total
		return total
	}

	fmt.Println("Part 2:", countPaths(startNode))
}
