package main

import (
	"fmt"

	"github.com/kamillo/aoc/fileutil"
)

type Point4d struct {
	x int
	y int
	z int
	w int
}

type Grid struct {
	dimensions int
	cubes map[Point4d]bool
	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
	minW int
	maxW int
}

func getGrid(lines []string) *Grid {
	result := &Grid{cubes: make(map[Point4d]bool)}

	for y, line := range lines {
		for x, c := range line {
			result.cubes[Point4d{x, y, 0, 0}] = c == '#'
			result.maxX = x
		}
		result.maxY = y
	}

	return result
}

func (grid *Grid) grow() {
	grid.minX--
	grid.minY--
	grid.minZ--
	grid.minW--
	grid.maxX++
	grid.maxY++
	grid.maxZ++
	grid.maxW++
}

func (grid *Grid) countNeighbors(x, y, z, w int) (count int) {
	check := func(dx, dy, dz, dw int) {
		if grid.dimensions == 3 {
			dw = 0
		}

		if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
			return
		}

		if grid.cubes[Point4d{x+dx, y+dy, z+dz, w+dw}] {
			count++
		}
	}

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					check(dx, dy, dz, dw)
				}
			}
		}
	}

	return count
}

func (grid *Grid) cycle() {
	temp := &Grid{cubes: make(map[Point4d]bool)}

	modifyCube := func(x, y, z, w int) {
		if grid.dimensions == 3 {
			w = 0
		}

		active := grid.cubes[Point4d{x, y, z, w}]
		neighborCount := grid.countNeighbors(x, y, z, w)

		if (active && (neighborCount == 2 || neighborCount == 3)) || !active && neighborCount == 3 {
			temp.cubes[Point4d{x, y, z, w}] = true
		} else {
			temp.cubes[Point4d{x, y, z, w}] = false
		}
	}

	for x := grid.minX - 1; x <= grid.maxX+1; x++ {
		for y := grid.minY - 1; y <= grid.maxY+1; y++ {
			for z := grid.minZ - 1; z <= grid.maxZ+1; z++ {
				for w := grid.minW - 1; w <= grid.maxW+1; w++ {
					modifyCube(x, y, z, w)
				}
			}
		}
	}

	grid.cubes = temp.cubes
	grid.grow()
}

func (grid *Grid) countActive() (count int) {
	for _, active := range grid.cubes {
		if active {
			count++
		}
	}

	return count
}

func main() {
	lines := fileutil.GetLines("input.txt")

	grid := getGrid(lines)
	grid.dimensions = 3
	for i := 0; i < 6; i++ {
		grid.cycle()
	}
	fmt.Println("Part 1: ", grid.countActive())

	grid = getGrid(lines)
	grid.dimensions = 4
	for i := 0; i < 6; i++ {
		grid.cycle()
	}
	fmt.Println("Part 2: ", grid.countActive())
}
