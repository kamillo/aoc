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

func main() {
	lines := utils.GetLines("input.txt")

	nodes := [100][100]Node{}
	nodesMap := []Node{}
	for _, line := range lines {
		x, y, size, used, avail, perc := 0, 0, 0, 0, 0, 0

		if _, err := fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT %d", &x, &y, &size, &used, &avail, &perc); err == nil {
			n := Node{x, y, size, used, avail}
			nodes[x][y] = n
			nodesMap = append(nodesMap, n)
		}
	}

	sum := 0
	for _, n1 := range nodesMap {
		for _, n2 := range nodesMap {
			if n1.used != 0 && n1 != n2 && n1.used < n2.avail {
				sum++
			}
		}
	}

	fmt.Println(sum)
	// for y := 0; y < len(nodes); y++ {
	// 	for x := 0; x < len(nodes); x++ {
	// 		if ()
	// 	}
	// }
}
