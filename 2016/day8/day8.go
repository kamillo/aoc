package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	screen := [6][50]int{}

	for _, line := range lines {
		x := 0
		y := 0
		if n, err := fmt.Sscanf(line, "rect %dx%d", &x, &y); err == nil && n == 2 {
			for i := 0; i < y; i++ {
				for j := 0; j < x; j++ {
					screen[i][j] = 1
				}
			}
		}

		if n, err := fmt.Sscanf(line, "rotate row y=%d by %d", &y, &x); err == nil && n == 2 {
			tmpRow := screen[y]
			for i := range screen[y] {
				screen[y][(i+x)%50] = tmpRow[i]
			}
		}

		if n, err := fmt.Sscanf(line, "rotate column x=%d by %d", &x, &y); err == nil && n == 2 {
			tmpCol := [6]int{}
			for i := 0; i < len(screen); i++ {
				tmpCol[i] = screen[i][x]
			}

			for i := 0; i < len(screen); i++ {
				screen[(i+y)%6][x] = tmpCol[i]
			}
		}
	}

	cnt := 0
	for i := range screen {
		fmt.Println(screen[i])
		for j := range screen[i] {
			if screen[i][j] == 1 {
				cnt++
			}
		}
	}
	fmt.Println("Part 1:", cnt)
}
