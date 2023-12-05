package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	seafloor := utils.GetLinesAs2dArray("input.txt")

	moves := []byte{'>', 'v'}
	moved := 1
	step := 0

	for moved > 0 {
		step++

		seafloorTmp := copyGrid(seafloor)
		moved = 0
		for _, m := range moves {
			for y := range seafloor {
				for x := range seafloor[y] {
					if seafloor[y][x] != m {
						continue
					}

					if yy, xx, ok := move(y, x, seafloor); ok {
						seafloorTmp[y][x], seafloorTmp[yy][xx] = '.', seafloorTmp[y][x]
						moved++
					}
				}
			}
			seafloor = copyGrid(seafloorTmp)
		}
	}

	printGrid(seafloor)
	fmt.Println(step)
}

func move(y int, x int, grid [][]byte) (newY, newX int, valid bool) {
	if grid[y][x] == '>' {
		newX = (x + 1) % len(grid[y])
		if grid[y][newX] == '.' {
			return y, newX, true
		}
	} else if grid[y][x] == 'v' {
		newY = (y + 1) % len(grid)
		if grid[newY][x] == '.' {
			return newY, x, true
		}
	}

	return y, x, false
}

func printGrid(grid [][]byte) {
	for y := range grid {
		for x := range grid[y] {
			fmt.Print(string(grid[y][x]))
		}
		fmt.Println()
	}
}

func copyGrid(grid [][]byte) [][]byte {
	duplicate := make([][]byte, len(grid))
	for i := range grid {
		duplicate[i] = make([]byte, len(grid[i]))
		copy(duplicate[i], grid[i])
	}

	return duplicate
}
