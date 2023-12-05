package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

type cube struct {
	x, y, z [2]int
	on      bool
}

func main() {
	lines := utils.GetLines("input.txt")

	fmt.Println("Part 1: ", run(lines, false))
	fmt.Println("Part 2: ", run(lines, true))
}

func run(lines []string, part2 bool) int {
	visited := []cube{}

	for _, line := range lines {
		var x, y, z [2]int
		state := ""

		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &state, &x[0], &x[1], &y[0], &y[1], &z[0], &z[1])

		cuboid := cube{x, y, z, state == "on"}

		if !part2 && isBig(cuboid) {
			continue
		}

		newCuboids := []cube{}
		if cuboid.on {
			newCuboids = append(newCuboids, cuboid)
		}

		for _, c := range visited {
			if intersection, ok := intersection(c, cuboid, !c.on); ok {
				newCuboids = append(newCuboids, intersection)
			}
		}

		visited = append(visited, newCuboids...)
	}

	sum := 0
	for _, v := range visited {
		if v.on {
			sum += size(v)
		} else {
			sum -= size(v)
		}
	}

	return sum
}

func intersect(aa, bb [2]int) ([2]int, bool) {
	l := utils.Max(aa[0], bb[0])
	r := utils.Min(aa[1], bb[1])

	if l <= r {
		return [2]int{l, r}, true
	}

	return [2]int{}, false
}

func intersection(c1, c2 cube, on bool) (cube, bool) {
	x, xx := intersect(c1.x, c2.x)
	y, yy := intersect(c1.y, c2.y)
	z, zz := intersect(c1.z, c2.z)

	if xx && yy && zz {
		return cube{x, y, z, on}, true
	}

	return cube{}, false
}

func isBig(c cube) bool {
	return c.x[0] < -50 || c.x[0] > 50 || c.x[1] < -50 || c.x[1] > 50 || c.y[0] < -50 || c.y[0] > 50 || c.y[1] < -50 || c.y[1] > 50 || c.z[0] < -50 || c.z[0] > 50 || c.z[1] < -50 || c.z[1] > 50
}

func size(c cube) int {
	return (c.x[1] - c.x[0] + 1) * (c.y[1] - c.y[0] + 1) * (c.z[1] - c.z[0] + 1)
}
