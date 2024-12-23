package main

import (
	"bytes"
	"fmt"
	"image"

	"github.com/kamillo/aoc/utils"
)

type Robot struct {
	pos image.Point
	vel image.Point
}

func main() {
	robots := []Robot{}
	// lines := utils.GetLines("test.txt")
	// maxX := 11
	// maxY := 7

	lines := utils.GetLines("input.txt")
	maxX := 101
	maxY := 103

	for _, line := range lines {
		p, v := image.Point{}, image.Point{}

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &p.X, &p.Y, &v.X, &v.Y)

		robots = append(robots, Robot{p, v})
	}

	part1, part2 := 0, 0
	middleX := maxX / 2
	middleY := maxY / 2

	for i := 0; ; i++ {
		for j, r := range robots {
			robots[j].pos = image.Pt(wrap(r.pos.X+r.vel.X, maxX), wrap(r.pos.Y+r.vel.Y, maxY))
		}

		if i == 99 {
			q1, q2, q3, q4 := 0, 0, 0, 0

			for _, r := range robots {
				if r.pos.X < middleX && r.pos.Y < middleY {
					q1++
				} else if r.pos.X > middleX && r.pos.Y < middleY {
					q2++
				} else if r.pos.X < middleX && r.pos.Y > middleY {
					q3++
				} else if r.pos.X > middleX && r.pos.Y > middleY {
					q4++
				}
			}
			part1 = q1 * q2 * q3 * q4
			fmt.Println("Part 1:", part1)
		}

		fmt.Println(i)
		print(robots, maxX, maxY)

	}
}

func print(robots []Robot, maxX, maxY int) {
	grid := make([][]byte, maxY)
	for i := range grid {
		grid[i] = make([]byte, maxX)
		grid[i] = bytes.Repeat([]byte{'.'}, maxX)
	}

	for _, r := range robots {
		grid[r.pos.Y][r.pos.X] = '#'
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func wrap(d, max int) int {
	if d < 0 {
		return max + d
	} else if d >= max {
		return d - max
	}

	return d
}
