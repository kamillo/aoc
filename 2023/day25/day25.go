package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Edge struct {
	src, dest int
}

type Graph struct {
	V, E int
	edge []Edge
}

type Subset struct {
	parent, rank int
}

func find(subsets []Subset, i int) int {
	if subsets[i].parent != i {
		subsets[i].parent = find(subsets, subsets[i].parent)
	}
	return subsets[i].parent
}

func union(subsets []Subset, x, y int) {
	xroot := find(subsets, x)
	yroot := find(subsets, y)

	if subsets[xroot].rank < subsets[yroot].rank {
		subsets[xroot].parent = yroot
	} else if subsets[xroot].rank > subsets[yroot].rank {
		subsets[yroot].parent = xroot
	} else {
		subsets[yroot].parent = xroot
		subsets[xroot].rank++
	}
}

func kargerMinCut(graph *Graph) map[Edge]bool {
	V := graph.V
	E := graph.E
	edge := graph.edge

	subsets := make([]Subset, V)

	for v := 0; v < V; v++ {
		subsets[v].parent = v
		subsets[v].rank = 0
	}

	vertices := V

	for vertices > 2 {
		i := rand.Intn(E)

		subset1 := find(subsets, edge[i].src)
		subset2 := find(subsets, edge[i].dest)

		if subset1 == subset2 {
			continue
		} else {
			vertices--
			union(subsets, subset1, subset2)
		}
	}

	cutedges := 0
	cutset := map[Edge]bool{}
	for i := 0; i < E; i++ {
		subset1 := find(subsets, edge[i].src)
		subset2 := find(subsets, edge[i].dest)
		if subset1 != subset2 {
			cutedges++
			cutset[edge[i]] = true
		}
	}

	return cutset
}

func createGraph(V, E int) *Graph {
	graph := &Graph{
		V:    V,
		E:    E,
		edge: make([]Edge, E),
	}
	return graph
}

func main() {
	lines := utils.GetLines("input.txt")

	id := 0
	vertices := map[string]int{}
	edges := []Edge{}
	adjList := map[int]map[int]bool{}

	for _, line := range lines {
		key := strings.Split(line, ": ")[0]
		nodes := strings.Split(strings.Split(line, ": ")[1], " ")

		if _, ok := vertices[key]; !ok {
			vertices[key] = id
			id++
		}

		if _, ok := adjList[vertices[key]]; !ok {
			adjList[vertices[key]] = map[int]bool{}
		}

		for _, n := range nodes {
			if _, ok := vertices[n]; ok {
				edges = append(edges, Edge{vertices[key], vertices[n]})
			} else {
				edges = append(edges, Edge{vertices[key], id})
				vertices[n] = id
				id++
			}

			adjList[vertices[key]][vertices[n]] = true
			if _, ok := adjList[vertices[n]]; !ok {
				adjList[vertices[n]] = map[int]bool{}
			}
			adjList[vertices[n]][vertices[key]] = true
		}
	}

	graph := createGraph(len(vertices), len(edges))
	graph.edge = edges

	s := kargerMinCut(graph)
	for m := 0; m != 3; {
		s = kargerMinCut(graph)
		m = len(s)
	}

	fmt.Println(s)

	for r := range s {
		delete(adjList[r.src], r.dest)
		delete(adjList[r.dest], r.src)
	}

	visited := map[int]bool{}
	queue := []int{0}
	visited[0] = true
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for n := range adjList[v] {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}

	fmt.Println("Part 1: ", len(visited)*(len(vertices)-len(visited)))
}
