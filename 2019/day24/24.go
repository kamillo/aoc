package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"math"
	"strings"
)

const gridSize = 5

type Grid [][]rune

func bugsLife(layers []Grid, tmpLayers []Grid, index int) {
	if index < 10 || index > 220 {
		return
	}
	grid := layers[index]
	var bugsInNeighborhood = func(x, y int) int {
		res := 0
		if x == 2 && y == 2 {
			return 0
		}
		if y < 4 && grid[y+1][x] == '#' {
			res++
		}
		if x < 4 && grid[y][x+1] == '#' {
			res++
		}
		if x > 0 && grid[y][x-1] == '#' {
			res++
		}
		if y > 0 && grid[y-1][x] == '#' {
			res++
		}
		if x == 4 && layers[index+1][2][3] == '#' {
			res++
		}
		if y == 4 && layers[index+1][3][2] == '#' {
			res++
		}
		if x == 0 && layers[index+1][2][1] == '#' {
			res++
		}
		if y == 0 && layers[index+1][1][2] == '#' {
			res++
		}

		if y == 2 && x == 1 {
			for i := range layers[index-1][0] {
				if layers[index-1][i][0] == '#' {
					res++
				}
			}
		}
		if y == 1 && x == 2 {
			for i := range layers[index-1][0] {
				if layers[index-1][0][i] == '#' {
					res++
				}
			}
		}
		if y == 3 && x == 2 {
			for i := range layers[index-1][0] {
				if layers[index-1][4][i] == '#' {
					res++
				}
			}
		}
		if y == 2 && x == 3 {
			for i := range layers[index-1][0] {
				if layers[index-1][i][4] == '#' {
					res++
				}
			}
		}
		return res
	}

	tmpGrid := make(Grid, gridSize)
	for y := range grid {
		tmpGrid[y] = make([]rune, gridSize)
		copy(tmpGrid[y], grid[y])
		for x := range grid[y] {
			if grid[y][x] == '#' && bugsInNeighborhood(x, y) != 1 {
				tmpGrid[y][x] = '.'
			} else if grid[y][x] == '.' && (bugsInNeighborhood(x, y) == 1 || bugsInNeighborhood(x, y) == 2) {
				tmpGrid[y][x] = '#'
			}
			//fmt.Printf("%c", tmpGrid[y][x])
		}
		//println()
	}

	tmpLayers[index] = tmpGrid
}

func main() {
	lines := fileutil.GetLines("input.txt")

	grid := make(Grid, gridSize)
	for i := range grid {
		grid[i] = []rune(lines[i])
	}
	fmt.Println(grid)

	var bugsInNeighborhood = func(x, y int) int {
		res := 0
		if y < 4 && grid[y+1][x] == '#' {
			res++
		}
		if x < 4 && grid[y][x+1] == '#' {
			res++
		}
		if x > 0 && grid[y][x-1] == '#' {
			res++
		}
		if y > 0 && grid[y-1][x] == '#' {
			res++
		}
		return res
	}

	// Part 1
	{
		layouts := make(map[string]bool)
		for {
			tmpGrid := make(Grid, 5)
			layout := ""
			for y := range grid {
				tmpGrid[y] = make([]rune, 5)
				copy(tmpGrid[y], grid[y])
				for x := range grid[y] {
					if grid[y][x] == '#' && bugsInNeighborhood(x, y) != 1 {
						tmpGrid[y][x] = '.'
					} else if grid[y][x] == '.' && (bugsInNeighborhood(x, y) == 1 || bugsInNeighborhood(x, y) == 2) {
						tmpGrid[y][x] = '#'
					}
					//fmt.Printf("%c", tmpGrid[y][x])
				}
				//println()
				layout += string(tmpGrid[y])
			}
			//fmt.Println(layout)
			if layouts[layout] {
				bioDiversity := 0.0
				for i := strings.LastIndex(layout, "#"); i != -1; i = strings.LastIndex(layout[:i], "#") {
					bioDiversity += math.Pow(float64(2), float64(i))
				}
				fmt.Println("Part 1:", int(bioDiversity))
				break
			}
			layouts[layout] = true
			grid = tmpGrid
		}
	}

	// Part 2
	{
		grid := make(Grid, gridSize)
		for i := range grid {
			grid[i] = []rune(lines[i])
		}

		layers := make([]Grid, 250)
		for i := range layers {
			layers[i] = make(Grid, gridSize)
			for j := range layers[i] {
				layers[i][j] = []rune(".....")
			}
		}

		layers[110] = grid
		for i := 0; i < 200; i++ {
			tmpLayers := make([]Grid, 250)
			for i := range tmpLayers {
				tmpLayers[i] = make(Grid, gridSize)
				for j := range tmpLayers[i] {
					tmpLayers[i][j] = make([]rune, gridSize)
					copy(tmpLayers[i][j], layers[i][j])
				}
			}

			bugsLife(layers, tmpLayers, 110)
			for j := 1; j < len(layers)/2; j++ {
				bugsLife(layers, tmpLayers, 110+j)
				bugsLife(layers, tmpLayers, 110-j)
			}

			layers = tmpLayers
		}
		bugs := 0
		for l := range layers {
			for r := range layers[l] {
				for c := range layers[l][r] {
					if layers[l][r][c] == '#' {
						bugs++
					}
				}
			}
		}

		fmt.Println("Part 2:", bugs)
	}
}
