// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package astar implements the A* shortest path finding algorithm.
package astar

import (
	"container/heap"
)

// The Graph interface is the minimal interface a graph data structure
// must satisfy to be suitable for the A* algorithm.
type Graph[Node any] interface {
	// Neighbours returns the neighbour nodes of node n in the graph.
	Neighbours(n Node, context string) map[string]Node
}

// A CostFunc is a function that returns a cost for the transition
// from node a to node b.
type CostFunc[Node any] func(a, b Node) float64

type ContextFunc[Node any] func(c string, a, b Node) string

// A Path is a sequence of nodes in a graph.
type Path[Node any] []Node

// newPath creates a new path with one start node. More nodes can be
// added with append().
func newPath[Node any](start Node) Path[Node] {
	return []Node{start}
}

// last returns the last node of path p. It is not removed from the path.
func (p Path[Node]) last() Node {
	return p[len(p)-1]
}

// cont creates a new path, which is a continuation of path p with the
// additional node n.
func (p Path[Node]) cont(n Node) Path[Node] {
	newPath := make([]Node, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, n)
	return newPath
}

// Cost calculates the total cost of path p by applying the cost function d
// to all path segments and returning the sum.
func (p Path[Node]) Cost(d CostFunc[Node]) (c float64) {
	for i := 1; i < len(p); i++ {
		c += d(p[i-1], p[i])
	}
	return c
}

type Value[Node any] struct {
	path    Path[Node]
	context string
}

// FindPath finds the shortest path between start and dest in graph g
// using the cost function d and the cost heuristic function h.
func FindShortestPath[Node comparable](g Graph[Node], start, dest Node, d, h CostFunc[Node], initialContext string, cf ContextFunc[Node]) (Path[Node], string) {
	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &item{value: Value[Node]{newPath(start), initialContext}})

	for pq.Len() > 0 {
		p := heap.Pop(pq).(*item).value.(Value[Node])
		n := p.path.last()
		c := p.context

		if n == dest {
			// Path found
			return p.path, p.context
		}

		for k, nb := range g.Neighbours(n, c) {
			newPath := p.path.cont(nb)

			heap.Push(pq, &item{
				value:    Value[Node]{newPath, c + k},
				priority: -(newPath.Cost(d) + h(nb, dest)),
			})
		}
	}

	// No path found
	return nil, ""
}

func FindLongestPath[Node comparable](g Graph[Node], start, dest Node, d, h CostFunc[Node], initialContext string, cf ContextFunc[Node]) []string {
	// closed := make(map[Node]bool)
	paths := []string{}

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &item{value: Value[Node]{newPath(start), initialContext}})

	for pq.Len() > 0 {
		p := heap.Pop(pq).(*item).value.(Value[Node])
		n := p.path.last()
		c := p.context

		// if closed[n] {
		// 	continue
		// }
		if n == dest {
			paths = append(paths, p.context)
			continue
			// Path found
			// return p.path, p.context
		}
		//closed[n] = true

		for k, nb := range g.Neighbours(n, c) {
			newPath := p.path.cont(nb)

			heap.Push(pq, &item{
				value:    Value[Node]{newPath, c + k},
				priority: (newPath.Cost(d) + h(nb, dest)),
			})
		}
	}

	// No path found
	return paths
}

func FindAllPaths[Node comparable](g Graph[Node], start, dest Node, d, h CostFunc[Node]) []Path[Node] {
	closed := make(map[Node]bool)
	paths := []Path[Node]{}

	pq := &priorityQueue{}

	heap.Init(pq)
	heap.Push(pq, &item{value: newPath(start)})

	for pq.Len() > 0 {
		// p := heap.Pop(pq).(*item[Path[Node]]).value
		p := heap.Pop(pq).(*item).value.(Path[Node])
		n := p.last()

		if closed[n] {
			continue
		}
		if n == dest {
			paths = append(paths, p)
			continue
			// Path found
			// return p.path, p.context
		}
		closed[n] = true

		for _, nb := range g.Neighbours(n, "") {
			newPath := p.cont(nb)

			heap.Push(pq, &item{
				value:    newPath,
				priority: (newPath.Cost(d) + h(nb, dest)),
			})
		}
	}

	// No path found
	return paths
}
