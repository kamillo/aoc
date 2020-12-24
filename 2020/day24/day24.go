package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"github.com/kamillo/aoc/utils/hex"
	"strings"
)

type hexgrid map[hex.Hex]bool

func main() {
	grid := hexgrid{}
	dirs := map[string]hex.Direction{
		"se": hex.DirectionSE,
		"e":  hex.DirectionE,
		"ne": hex.DirectionNE,
		"nw": hex.DirectionNW,
		"sw": hex.DirectionSW,
		"w":  hex.DirectionW,
	}
	ref := hex.NewHex(0, 0)

	for _, line := range utils.GetLines("input.txt") {
		current := ref
		for len(line) > 0 {
			diag := ""
			if len(line) >= 2 {
				diag = line[:2]
			} else {
				diag = line[:1]
			}

			d, found := dirs[diag]
			n := hex.Neighbor(current, d)

			if !found {
				diag = string(line[:1])
				d, found = dirs[diag]
				n = hex.Neighbor(current, d)
			}
			if found {
				current = n
			}
			line = strings.TrimPrefix(line, diag)
		}

		if _, exist := grid[current]; !exist {
			grid[current] = true
		} else {
			delete(grid, current)
		}
	}

	fmt.Printf("Part 1: %d\n", len(grid))

	for i := 0; i < 100; i++ {
		neigh := map[hex.Hex]int{}
		for k := range grid {
			for _, d := range hex.Neighbors(k) {
				neigh[d]++
			}
		}
		newHexgrid := map[hex.Hex]bool{}
		for p, n := range neigh {
			if _, exist := grid[p]; exist && n == 1 || n == 2 {
				newHexgrid[p] = true
			}
		}

		grid = newHexgrid
	}

	fmt.Printf("Part 2: %d\n", len(grid))
}
