package main

import (
	"fmt"
	"math"

	"github.com/kamillo/aoc/utils"
)

func main() {
	crabs := utils.ToIntArr(utils.GetLines("input.txt")[0], ",")

	fmt.Println("Part 1: ", part1(crabs))
	fmt.Println("Part 2: ", part2(crabs))
}

func part1(crabs []int) int {
	costs := map[int]int{}
	max := utils.MaxInArray(crabs)

	for _, k := range crabs {
		for l := 0; l < max; l++ {
			costs[l] += int(math.Abs(float64(k - l)))
		}
	}

	min := math.MaxInt64
	for _, v := range costs {
		if v < min {
			min = v
		}
	}

	return min
}

func part2(crabs []int) int {
	costs := map[int]int{}
	max := utils.MaxInArray(crabs)
	tmpCosts := map[int]int{}

	for _, k := range crabs {
		for l := 0; l < max; l++ {
			if k == l {
				continue
			}

			dist := int(math.Abs(float64(k - l)))
			if tmpCosts[dist] > 0 {
				costs[l] += tmpCosts[dist]
			} else {
				for i := 1; i <= dist; i++ {
					costs[l] += i
					tmpCosts[dist] += i
				}
			}
		}
	}

	min := math.MaxInt64
	for _, v := range costs {
		if v < min {
			min = v
		}
	}

	return min
}
