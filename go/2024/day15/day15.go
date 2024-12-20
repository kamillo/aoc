package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	// lines := utils.GetLines("test.txt")
	lines := utils.GetLines("input.txt")

	part1 := 0
	part2 := 0

	grid := [][]byte{}
	grid2 := [][]byte{}
	steps := false
	robot := image.Point{}
	robot2 := image.Point{}

	for y, line := range lines {
		if !steps {
			if len(line) == 0 {
				steps = true
				continue
			}

			grid = append(grid, []byte(line))
			if i := strings.Index(line, "@"); i != -1 {
				robot.X = i
				robot.Y = y

				robot2.X = i * 2
				robot2.Y = y
			}

			grid2 = append(grid2, []byte{})
			for _, c := range []byte(line) {
				newChar1 := c
				newChar2 := c

				if c == 'O' {
					newChar1 = '['
					newChar2 = ']'
				}

				if c == '@' {
					newChar1 = '@'
					newChar2 = '.'
				}

				grid2[len(grid2)-1] = append(grid2[len(grid2)-1], newChar1, newChar2)
			}

		} else {
			for _, c := range line {
				switch c {
				case '>':
					robot = image.Pt(move(grid, robot.X, robot.Y, 1, 0))
					robot2 = image.Pt(move2(grid2, robot2.X, robot2.Y, 1, 0))
				case '<':
					robot = image.Pt(move(grid, robot.X, robot.Y, -1, 0))
					robot2 = image.Pt(move2(grid2, robot2.X, robot2.Y, -1, 0))
				case '^':
					robot = image.Pt(move(grid, robot.X, robot.Y, 0, -1))
					robot2 = image.Pt(move2(grid2, robot2.X, robot2.Y, 0, -1))
				case 'v':
					robot = image.Pt(move(grid, robot.X, robot.Y, 0, 1))
					robot2 = image.Pt(move2(grid2, robot2.X, robot2.Y, 0, 1))
				}
			}
		}
	}

	print(grid)
	print(grid2)

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				part1 += 100*y + x
			}
		}
	}

	for y := range grid2 {
		for x := range grid2[y] {
			if grid2[y][x] == '[' {
				part2 += 100*y + x
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func print(grid [][]byte) {
	for y := range grid {
		for x := range grid[y] {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
}

func move(grid [][]byte, x, y, dx, dy int) (rx, ry int) {
	retx, rety := x, y

	free := -1
	if dy != 0 {
		for ny := y; grid[ny][x] != '#'; ny += dy {
			if grid[ny][x] == '.' {
				free = ny
				break
			}
		}

		if free != -1 {
			rety += dy
			dy *= -1
			for ny := free; ny != y; ny += dy {
				grid[ny][x] = grid[ny+dy][x]
			}
		}
	}

	if dx != 0 {
		for nx := x; grid[y][nx] != '#'; nx += dx {
			if grid[y][nx] == '.' {
				free = nx
				break
			}
		}

		if free != -1 {
			retx += dx
			dx *= -1
			for nx := free; nx != x; nx += dx {
				grid[y][nx] = grid[y][nx+dx]
			}
		}
	}

	if free != -1 {
		grid[y][x] = '.'
	}

	return retx, rety
}

type Tree map[image.Point]byte

func move2(grid [][]byte, x, y, dx, dy int) (rx, ry int) {
	retx, rety := x, y

	if dy != 0 {
		tree := Tree{}
		tree[image.Pt(x, y)] = grid[y][x]

		var find func(tree Tree, grid [][]byte, x, y, dx, dy int) bool

		find = func(tree Tree, grid [][]byte, x, y, dx, dy int) bool {
			ret := true
			// fmt.Println(x, y, dx, dy)

			if y+dy >= len(grid) || grid[y+dy][x] == '#' {
				return false
			}

			if grid[y+dy][x] == '[' {
				tree[image.Pt(x, y+dy)] = grid[y+dy][x]
				ret = ret && find(tree, grid, x, y+dy, dx, dy)
				tree[image.Pt(x+1, y+dy)] = grid[y+dy][x+1]
				ret = ret && find(tree, grid, x+1, y+dy, dx, dy)
			} else if grid[y+dy][x] == ']' {
				tree[image.Pt(x, y+dy)] = grid[y+dy][x]
				ret = ret && find(tree, grid, x, y+dy, dx, dy)
				tree[image.Pt(x-1, y+dy)] = grid[y+dy][x-1]
				ret = ret && find(tree, grid, x-1, y+dy, dx, dy)
			}

			if !ret {
				return false
			}

			found := false
			for ny := y; grid[ny][x] != '#'; ny += dy {
				if grid[ny][x] == '.' {
					found = true
					break
				}
			}

			return found
		}

		if ok := find(tree, grid, x, y, dx, dy); ok {
			for k := range tree {
				grid[k.Y][k.X] = '.'
			}

			for k, v := range tree {
				grid[k.Y+dy][k.X+dx] = v
			}

			retx += dx
			rety += dy
		}
	}

	if dx != 0 {
		free := -1
		for nx := x; grid[y][nx] != '#'; nx += dx {
			if grid[y][nx] == '.' {
				free = nx
				break
			}
		}

		if free != -1 {
			retx += dx
			dx *= -1
			for nx := free; nx != x; nx += dx {
				grid[y][nx] = grid[y][nx+dx]
			}
		}
		if free != -1 {
			grid[y][x] = '.'
		}
	}

	return retx, rety
}
