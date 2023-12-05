package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	cubes := map[utils.PointD3D]bool{}
	for _, line := range utils.GetLines("input.txt") {
		x, y, z := 0, 0, 0
		x1, y1, z1 := 0, 0, 0
		state := ""
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &state, &x, &x1, &y, &y1, &z, &z1)

		if x < -50 || x > 50 || x1 < -50 || x1 > 50 || y < -50 || y > 50 || y1 < -50 || y1 > 50 || z < -50 || z > 50 || z1 < -50 || z1 > 50 {
			continue
		}
		for i := x; i <= x1; i++ {
			for j := y; j <= y1; j++ {
				for k := z; k <= z1; k++ {
					if state == "on" {
						cubes[utils.NewPointD3D(i, j, k)] = true
					} else if cubes[utils.NewPointD3D(i, j, k)] {
						cubes[utils.NewPointD3D(i, j, k)] = false
					}
				}
			}
		}
	}

	count := 0
	for _, v := range cubes {
		if v {
			count++
		}
	}

	fmt.Println(count)
}
