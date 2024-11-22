package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	findOverlapping := func(part2 bool) int {
		vents := map[utils.Point2D]int{}

		for _, l := range lines {
			field := strings.Fields(l)

			p1 := utils.ToIntArr(field[0], ",")
			p2 := utils.ToIntArr(field[2], ",")

			var line utils.DiscreteLine
			if p1[0] > p2[0] {
				line = utils.DiscreteLine{p2[0], p2[1], p1[0], p1[1]}
			} else {
				line = utils.DiscreteLine{p1[0], p1[1], p2[0], p2[1]}
			}

			if line.StartX == line.EndX {
				for i := utils.Min(line.StartY, line.EndY); i <= utils.Max(line.StartY, line.EndY); i++ {
					vents[utils.Point2D{line.StartX, i}]++
				}
			} else if line.StartY == line.EndY {
				for i := utils.Min(line.StartX, line.EndX); i <= utils.Max(line.StartX, line.EndX); i++ {
					vents[utils.Point2D{i, line.StartY}]++
				}
			} else if part2 {
				if line.StartX < line.EndX && line.StartY < line.EndY {
					for i := line.StartX; i <= line.EndX; i++ {
						j := i - line.StartX
						vents[utils.Point2D{line.StartX + j, line.StartY + j}]++
					}
				} else {
					for i := line.StartX; i <= line.EndX; i++ {
						j := i - line.StartX
						vents[utils.Point2D{line.StartX + j, line.StartY - j}]++
					}
				}
			}
		}

		count := 0
		for _, v := range vents {
			if v > 1 {
				count++
			}
		}

		return count
	}

	fmt.Println("Part 1:", findOverlapping(false))
	fmt.Println("Part 2:", findOverlapping(true))
}
