package main

import (
	"fmt"
	"image"

	"github.com/kamillo/aoc/utils"
)

var (
	up    = image.Pt(0, -1)
	down  = image.Pt(0, 1)
	left  = image.Pt(-1, 0)
	right = image.Pt(1, 0)
)

func part1() {
	m := 1001
	lines := utils.GetLines("input.txt")
	grid := make([][]byte, len(lines)*m)
	for y := range grid {
		grid[y] = make([]byte, len(lines)*m)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	for y := range lines {
		for x := range lines[y] {
			grid[len(lines)*(m/2)+y][len(lines)*(m/2)+x] = lines[y][x]
		}
	}

	virus := image.Point{len(grid) / 2, len(grid) / 2}
	direction := up
	inf := 0

	for i := 0; i < 10000; i++ {
		if grid[virus.Y][virus.X] == '#' {
			grid[virus.Y][virus.X] = '.'
			direction = turnRight(direction, virus)
			virus = virus.Add(direction)

		} else {
			grid[virus.Y][virus.X] = '#'
			direction = turnLeft(direction, virus)
			virus = virus.Add(direction)
			inf++
		}

	}

	// for y := range grid {
	// 	for x := range grid[y] {
	// 		fmt.Printf("%c", grid[y][x])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()
	fmt.Println("Part 1:", inf)
}

func part2() {
	m := 1001
	lines := utils.GetLines("input.txt")
	grid := make([][]byte, len(lines)*m)
	for y := range grid {
		grid[y] = make([]byte, len(lines)*m)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	for y := range lines {
		for x := range lines[y] {
			grid[len(lines)*(m/2)+y][len(lines)*(m/2)+x] = lines[y][x]
		}
	}

	virus := image.Point{len(grid) / 2, len(grid) / 2}
	direction := up
	inf := 0

	for i := 0; i < 10000000; i++ {
		switch grid[virus.Y][virus.X] {
		case '.':
			grid[virus.Y][virus.X] = 'W'
			direction = turnLeft(direction, virus)
			virus = virus.Add(direction)
		case 'W':
			grid[virus.Y][virus.X] = '#'
			virus = virus.Add(direction)
			inf++
		case '#':
			grid[virus.Y][virus.X] = 'F'
			direction = turnRight(direction, virus)
			virus = virus.Add(direction)
		case 'F':
			grid[virus.Y][virus.X] = '.'
			direction = turnBack(direction, virus)
			virus = virus.Add(direction)
		}
	}
	fmt.Println("Part 2:", inf)
}

func main() {
	part1()
	part2()
}

func turnLeft(d, p image.Point) image.Point {
	switch d {
	case up:
		return left
	case down:
		return right
	case left:
		return down
	case right:
		return up
	}

	return d
}

func turnRight(d, p image.Point) image.Point {
	switch d {
	case up:
		return right
	case down:
		return left
	case left:
		return up
	case right:
		return down
	}

	return d
}

func turnBack(d, p image.Point) image.Point {
	switch d {
	case up:
		return down
	case down:
		return up
	case left:
		return right
	case right:
		return left
	}

	return d
}
