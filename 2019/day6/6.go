package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"os"
	"strings"

	"github.com/albertorestifo/dijkstra"
)

type Object struct {
	name            string
	orbitingObjects []*Object
}

func main() {
	lines := utils.GetLines(os.Args[1])
	//lines = []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}

	graph := dijkstra.Graph{}
	objects := make(map[string]bool)
	root := "COM"
	for _, line := range lines {
		splitted := strings.Split(line, ")")
		objects[splitted[0]] = true
		objects[splitted[1]] = true

		if _, ok := graph[splitted[0]]; !ok {
			graph[splitted[0]] = map[string]int{}
		}
		if _, ok := graph[splitted[1]]; !ok {
			graph[splitted[1]] = map[string]int{}
		}

		graph[splitted[0]][splitted[1]] = 1
		graph[splitted[1]][splitted[0]] = 1

	}

	sum := 0
	for k := range objects {
		if k != root {
			_, cost, _ := graph.Path(root, k)
			sum += cost
		}
	}
	fmt.Println("Part 1:", sum)

	_, cost, _ := graph.Path("YOU", "SAN")
	fmt.Println("Part 2: ", cost-2)
}
