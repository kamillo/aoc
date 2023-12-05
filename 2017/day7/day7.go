package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Program struct {
	name   string
	weight int
	parent string
	next   []string
	level  int
}
type Graph map[string]map[string]bool

type TreeNode struct {
	Name string
	Data int
}
type Tree map[TreeNode]Tree

func main() {
	graph := Graph{}
	programms := map[string]Program{}

	for _, line := range utils.GetLines("input.txt") {
		split := strings.Split(line, " -> ")
		name := ""
		weight := 0
		fmt.Sscanf(split[0], "%s (%d)", &name, &weight)
		if _, ok := programms[name]; !ok {
			programms[name] = Program{name, weight, "", []string{}, 0}
		} else {
			p := programms[name]
			p.weight = weight
			programms[name] = p
		}
		graph[name] = make(map[string]bool)

		if len(split) == 2 {
			for _, node := range strings.Split(split[1], ", ") {
				graph[name][node] = true
				if _, ok := programms[node]; !ok {
					programms[node] = Program{node, 0, name, []string{}, 0}
				} else {
					p := programms[node]
					p.parent = name
					programms[node] = p
				}
				p := programms[name]
				p.next = append(p.next, node)
				programms[name] = p
			}
		}
	}

	root := ""
	for k, v := range programms {
		if v.parent == "" {
			fmt.Println("Part 1:", k)
			root = v.name
			break
		}
	}

	setLevel(root, programms, 0)

	level := 0
	for level < 5 {
		for name, p := range programms {
			weights := map[int][]string{}
			if p.level == level {
				for _, nextRoot := range programms[name].next {
					w := checkWeights(nextRoot, programms)
					weights[w] = append(weights[w], nextRoot)
				}
			}
			if len(weights) >= 2 {
				diff := 0
				//fmt.Println(level, weights)
				wrongW := 0
				for k, w := range weights {
					diff = int(math.Abs(float64(diff - k)))
					if len(w) == 1 {
						wrongW = programms[w[0]].weight
					}
				}
				fmt.Println("Part 2:", wrongW-diff)
			}
		}

		level++
	}
}

func checkWeights(root string, programms map[string]Program) int {
	branchWeight := programms[root].weight
	for _, child := range programms[root].next {
		branchWeight += checkWeights(child, programms)
	}

	return branchWeight
}

func setLevel(root string, programms map[string]Program, level int) {
	node := programms[root]
	node.level = level
	programms[root] = node
	level++

	for _, child := range programms[root].next {
		setLevel(child, programms, level)
	}
}
