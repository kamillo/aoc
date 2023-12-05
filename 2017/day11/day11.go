package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
	"github.com/kamillo/aoc/utils/hex"
)

func main() {
	dirs := map[string]hex.Direction{
		"s":  hex.DirectionSE,
		"se": hex.DirectionE,
		"ne": hex.DirectionNE,
		"n":  hex.DirectionNW,
		"sw": hex.DirectionSW,
		"nw": hex.DirectionW,
	}

	start := hex.NewHex(0, 0)
	end := hex.NewHex(0, 0)
	max := 0

	for _, dir := range strings.Split(utils.GetLines("input.txt")[0], ",") {
		end = hex.Neighbor(end, dirs[dir])
		distance := hex.Distance(start, end)
		if max < distance {
			max = distance
			if max > 1580 {
				fmt.Println(max, start, end)
			}
		}
	}

	fmt.Println("Part 1", hex.Distance(start, end))
	fmt.Println("Part 2", max)
}
