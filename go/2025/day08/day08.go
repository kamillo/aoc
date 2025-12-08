package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("test.txt")

	part1 := 1
	part2 := 0

	boxes := []utils.PointD3D{}
	for _, line := range lines {
		x, y, z := 0, 0, 0
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		boxes = append(boxes, utils.PointD3D{X: x, Y: y, Z: z})
	}

	// distances := map[utils.PointD3D]map[utils.PointD3D]int{}
	// for _, box := range boxes {
	// 	for _, box2 := range boxes {
	// 		distances[box][box2] = distance(box, box2)
	// 	}
	// }

	distances := []struct {
		box1 utils.PointD3D
		box2 utils.PointD3D
		dist int
	}{}

	for _, box := range boxes {
		for _, box2 := range boxes {
			if box == box2 {
				continue
			}

			if utils.Contains(distances, struct {
				box1 utils.PointD3D
				box2 utils.PointD3D
				dist int
			}{box2, box, distance(box, box2)}) {
				continue
			}

			distances = append(distances, struct {
				box1 utils.PointD3D
				box2 utils.PointD3D
				dist int
			}{box, box2, distance(box, box2)})
		}
	}

	// sort distances by value
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	fmt.Println(distances)

	circuits := []map[utils.PointD3D]bool{}
	for i := 0; i < 10; i++ {
		found := false
		for _, c := range circuits {
			if _, ok := c[distances[i].box1]; ok {
				c[distances[i].box2] = true
				found = true
			}

			if _, ok := c[distances[i].box2]; ok {
				c[distances[i].box1] = true
				found = true
			}

		}
		if !found {
			circuits = append(circuits, map[utils.PointD3D]bool{
				distances[i].box1: true,
				distances[i].box2: true,
			})
		}
	}

	for _, c := range circuits {
		fmt.Println(c)
		part1 *= len(c)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

// euclidean distance
func distance(a, b utils.PointD3D) int {
	return int(math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2)))
}
