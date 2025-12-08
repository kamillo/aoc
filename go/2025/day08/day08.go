package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	boxes := []utils.PointD3D{}
	for _, line := range lines {
		x, y, z := 0, 0, 0
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		boxes = append(boxes, utils.PointD3D{X: x, Y: y, Z: z})
	}

	type Distance struct {
		box1 utils.PointD3D
		box2 utils.PointD3D
		dist int
	}
	distances := []Distance{}

	pairs := map[Distance]bool{}
	for _, box := range boxes {
		for _, box2 := range boxes {
			if box == box2 {
				continue
			}

			if pairs[Distance{box, box2, distance(box, box2)}] || pairs[Distance{box2, box, distance(box, box2)}] {
				continue
			}

			pairs[Distance{box, box2, distance(box, box2)}] = true

			distances = append(distances, struct {
				box1 utils.PointD3D
				box2 utils.PointD3D
				dist int
			}{box, box2, distance(box, box2)})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	// fmt.Println(distances)

	parent := map[utils.PointD3D]utils.PointD3D{}
	for _, b := range boxes {
		parent[b] = b
	}

	var find func(p utils.PointD3D) utils.PointD3D
	find = func(p utils.PointD3D) utils.PointD3D {
		if parent[p] != p {
			parent[p] = find(parent[p])
		}
		return parent[p]
	}

	union := func(p1, p2 utils.PointD3D) {
		root1 := find(p1)
		root2 := find(p2)
		if root1 != root2 {
			parent[root1] = root2
		}
	}

	for i := range distances {
		union(distances[i].box1, distances[i].box2)

		circuits := map[utils.PointD3D][]utils.PointD3D{}
		for _, box := range boxes {
			root := find(box)
			circuits[root] = append(circuits[root], box)
		}

		if i == 999 {
			circuitsLengths := []int{}
			for _, c := range circuits {
				circuitsLengths = append(circuitsLengths, len(c))
			}

			sort.Sort(sort.Reverse(sort.IntSlice(circuitsLengths)))

			fmt.Println("Part 1:", circuitsLengths[0]*circuitsLengths[1]*circuitsLengths[2])
		}

		if len(circuits) == 1 {
			fmt.Println("Part 2:", distances[i].box1.X*distances[i].box2.X)
			break
		}
	}

}

func distance(a, b utils.PointD3D) int {
	return int(math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2)))
}
