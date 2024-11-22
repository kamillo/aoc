package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type TreeNode string

// type TreeNode struct {
// 	Name string
// }
type Tree map[string]map[string]bool

func main() {
	pipes := Tree{}
	for _, line := range utils.GetLines("input.txt") {
		split := strings.Split(line, " <-> ")
		if _, ok := pipes[split[0]]; !ok {
			pipes[split[0]] = map[string]bool{}
		}

		for _, p := range strings.Split(split[1], ", ") {
			// pipes[p] = pipes[split[0]]
			if _, ok := pipes[p]; !ok {
				pipes[p] = map[string]bool{}
			}
			pipes[split[0]][p] = true
			pipes[p][split[0]] = true
		}
	}

	nodes := map[string]bool{}
	var depth func(node string, nodes map[string]bool)
	depth = func(node string, nodes map[string]bool) {
		for k, _ := range pipes[node] {
			if _, ok := nodes[k]; !ok {
				nodes[k] = true
				depth(k, nodes)
			}
		}
	}

	depth("0", nodes)

	fmt.Println("Part 1:", len(nodes))

	groups := []map[string]bool{}

	for k, _ := range pipes {
		found := false
		for _, g := range groups {
			if found = g[k]; found {
				break
			}
		}

		if !found {
			groups = append(groups, map[string]bool{})
		}

		last := len(groups) - 1

		if _, ok := groups[last][k]; !ok {
			groups[last][k] = true
			depth(k, groups[last])
		}
	}

	fmt.Println("Part 2:", len(groups))
}
