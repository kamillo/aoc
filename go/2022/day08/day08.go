package main

import (
	"fmt"
	"image"

	"github.com/kamillo/aoc/utils"
)

var offsets = []image.Point{
	image.Pt(0, -1), // North
	image.Pt(1, 0),  // East
	image.Pt(0, 1),  // South
	image.Pt(-1, 0), // West
}

func main() {
	trees := utils.GetLinesAs2dIntArray("input.txt")

	cnt := 0
	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees)-1; x++ {
			adj := CountAdj(y, x, trees)
			if adj > 0 {
				cnt++
			}
		}
	}

	fmt.Println("Part 1:", 4*len(trees[0])-4+cnt)

	max := 0
	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees)-1; x++ {
			adj := CountAdj2(y, x, trees)
			if adj > max {
				max = adj
			}
		}
	}

	fmt.Println("Part 2:", max)
}

func CountAdj2(x int, y int, grid [][]int) (ret int) {
	ret = 1
	for _, off := range offsets {
		p := image.Pt(x, y)
		q := image.Pt(x, y)
		adj := 0

		for {
			q = q.Add(off)
			if q.X < 0 || q.Y < 0 || q.X >= len(grid) || q.Y >= len(grid) {
				break
			}

			adj++
			if grid[q.X][q.Y] >= grid[p.X][p.Y] {
				break
			}
		}

		ret *= adj
	}

	return ret
}

func CountAdj(x int, y int, grid [][]int) (adj int) {
	for _, off := range offsets {
		p := image.Pt(x, y)
		q := image.Pt(x, y)
		visible := true

		for {
			q = q.Add(off)
			if q.X < 0 || q.Y < 0 || q.X >= len(grid) || q.Y >= len(grid) {
				break
			}

			if grid[q.X][q.Y] >= grid[p.X][p.Y] {
				visible = false
				break
			}
		}

		if visible {
			adj++
		}
	}

	return adj
}
