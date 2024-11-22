package main

import (
	"fmt"
	"time"

	"github.com/kamillo/aoc/utils"
)

type Points map[utils.Point2D]bool

func main() {
	grid := [][]int{}
	for _, line := range utils.GetLines("input.txt") {
		grid = append(grid, utils.ToIntArr(line, ""))
	}

	totalFlashes := 0
	for step := 0; ; step++ {
		toFlash := []utils.Point2D{}

		for x := range grid {
			for y := range grid[x] {
				grid[x][y]++

				if grid[x][y] > 9 {
					toFlash = append(toFlash, utils.Point2D{x, y})
				}
			}
		}

		flashed := Points{}
		for _, p := range toFlash {
			grid[p.X][p.Y] = 0
			flashed[p] = true
		}
		for _, p := range toFlash {
			modifyAdj(p.X, p.Y, grid, 9, flashed)
		}

		totalFlashes += len(flashed)

		printGrid(grid)
		if step == 99 {
			fmt.Println("Part 1:", totalFlashes)
		}

		if len(flashed) == len(grid)*len(grid[0]) {
			fmt.Println("Part 2:", step+1)
			break
		}
	}
}

func flash(grid [][]int, x int, y int, flashed Points) {
	grid[x][y] = 0
	flashed[utils.Point2D{x, y}] = true
	modifyAdj(x, y, grid, 9, flashed)
}

func modifyAdj(x int, y int, grid [][]int, char int, flashed Points) {
	if x+1 < len(grid) && !flashed[utils.Point2D{x + 1, y}] {
		grid[x+1][y]++
		if grid[x+1][y] > char {
			flash(grid, x+1, y, flashed)
		}
	}
	if y+1 < len(grid[x]) && !flashed[utils.Point2D{x, y + 1}] {
		grid[x][y+1]++
		if grid[x][y+1] > char {
			flash(grid, x, y+1, flashed)
		}
	}
	if x+1 < len(grid) && y+1 < len(grid[x]) && !flashed[utils.Point2D{x + 1, y + 1}] {
		grid[x+1][y+1]++
		if grid[x+1][y+1] > char {
			flash(grid, x+1, y+1, flashed)
		}
	}
	if x > 0 && !flashed[utils.Point2D{x - 1, y}] {
		grid[x-1][y]++
		if grid[x-1][y] > char {
			flash(grid, x-1, y, flashed)
		}
	}
	if x > 0 && y+1 < len(grid[x]) && !flashed[utils.Point2D{x - 1, y + 1}] {
		grid[x-1][y+1]++
		if grid[x-1][y+1] > char {
			flash(grid, x-1, y+1, flashed)
		}
	}
	if y > 0 && !flashed[utils.Point2D{x, y - 1}] {
		grid[x][y-1]++
		if grid[x][y-1] > char {
			flash(grid, x, y-1, flashed)
		}
	}
	if y > 0 && x+1 < len(grid) && !flashed[utils.Point2D{x + 1, y - 1}] {
		grid[x+1][y-1]++
		if grid[x+1][y-1] > char {
			flash(grid, x+1, y-1, flashed)
		}
	}
	if x > 0 && y > 0 && !flashed[utils.Point2D{x - 1, y - 1}] {
		grid[x-1][y-1]++
		if grid[x-1][y-1] > char {
			flash(grid, x-1, y-1, flashed)
		}
	}
}

func printGrid(grid [][]int) {
	clear()
	for x := range grid {
		for y := range grid[x] {
			fmt.Print("\u001b[48;5;"+fmt.Sprint(235+grid[x][y]*2)+"m ", " ")
		}
		fmt.Println("\u001b[0m")
	}
	time.Sleep(time.Millisecond * 200)
}

func clear() {
	fmt.Print("\033[H\033[2J")
}
