package main

import (
	"fmt"
	"sort"

	"github.com/kamillo/aoc/utils"
)

func main() {
	grid := [][]int{}
	for _, line := range utils.GetLines("input.txt") {
		row := utils.ToIntArr(line, "")
		grid = append(grid, row)
	}

	risk := 0
	basins := []int{}
	visited := map[utils.Point2D]bool{}
	for y := range grid {
		for x := range grid[y] {
			if checkAdj(y, x, grid, grid[y][x]) {
				risk += grid[y][x] + 1

				basins = append(basins, calcBasin(y, x, grid, visited))
			}
		}
	}

	fmt.Println("Part 1: ", risk)

	//sort.Ints(basins)
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	fmt.Println("Part 2: ", basins[0]*basins[1]*basins[2])
}

func checkAdj(x int, y int, grid [][]int, char int) (adj bool) {
	adj = true
	if x+1 < len(grid) {
		adj = adj && grid[x+1][y] > char
	}

	if y+1 < len(grid[x]) {
		adj = adj && grid[x][y+1] > char
	}

	if x > 0 {
		adj = adj && grid[x-1][y] > char
	}

	if y > 0 {
		adj = adj && grid[x][y-1] > char
	}

	return adj
}

func calcBasin(x int, y int, grid [][]int, visited map[utils.Point2D]bool) (size int) {
	if x+1 < len(grid) && grid[x+1][y] < 9 && !visited[utils.Point2D{x + 1, y}] {
		size++
		visited[utils.Point2D{x + 1, y}] = true
		size += calcBasin(x+1, y, grid, visited)
	}

	if y+1 < len(grid[x]) && grid[x][y+1] < 9 && !visited[utils.Point2D{x, y + 1}] {
		size++
		visited[utils.Point2D{x, y + 1}] = true
		size += calcBasin(x, y+1, grid, visited)
	}

	if x > 0 && grid[x-1][y] < 9 && !visited[utils.Point2D{x - 1, y}] {
		size++
		visited[utils.Point2D{x - 1, y}] = true
		size += calcBasin(x-1, y, grid, visited)
	}

	if y > 0 && grid[x][y-1] < 9 && !visited[utils.Point2D{x, y - 1}] {
		size++
		visited[utils.Point2D{x, y - 1}] = true
		size += calcBasin(x, y-1, grid, visited)
	}

	return
}
