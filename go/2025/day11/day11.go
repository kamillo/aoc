package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Node string

func main() {
	lines := utils.GetLines("input.txt")

	graph := utils.Graph[Node]{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		device := Node(split[0][:len(split[0])-1])

		for _, s := range split[1:] {
			graph[device] = append(graph[device], Node(s))
		}
	}
	compare := func(n Node) bool { return n == Node("out") }
	fmt.Println(len(utils.AllPaths(graph, Node("you"), compare)))

	type State struct {
		u      Node
		hasDac bool
		hasFft bool
	}

	memo := map[State]int{}
	visiting := map[Node]bool{}

	var countPaths func(Node, bool, bool) int
	countPaths = func(u Node, hasDac, hasFft bool) int {
		// Update state based on current node
		if u == "dac" {
			hasDac = true
		}
		if u == "fft" {
			hasFft = true
		}

		if compare(u) {
			if hasDac && hasFft {
				return 1
			}
			return 0
		}

		state := State{u, hasDac, hasFft}
		if cnt, ok := memo[state]; ok {
			return cnt
		}

		if visiting[u] {
			return 0
		}
		visiting[u] = true

		total := 0
		for _, v := range graph[u] {
			total += countPaths(v, hasDac, hasFft)
		}

		visiting[u] = false
		memo[state] = total
		return total
	}

	total := countPaths(Node("svr"), false, false)
	fmt.Printf("Part 2: %d\n", total)
}
