package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/kamillo/aoc/utils"
)

var (
	up    = image.Pt(0, -1)
	down  = image.Pt(0, 1)
	left  = image.Pt(-1, 0)
	right = image.Pt(1, 0)
)

func main() {
	lines := utils.GetLines("input.txt")

	point := image.Point{strings.Index(lines[0], "|"), 0}
	nodes := []string{}
	steps := 1

	goToNextNode := func(direction image.Point, p image.Point) {
		getNextPoint := func(dir, p image.Point) (image.Point, image.Point) {
			next := p.Add(dir)
			if lines[next.Y][next.X] != ' ' {
				return dir, next
			}

			left := turnLeft(dir, p)
			next = p.Add(left)
			if lines[next.Y][next.X] != ' ' {
				return left, next
			}

			right := turnRight(dir, p)
			next = p.Add(right)
			if lines[next.Y][next.X] != ' ' {
				return right, next
			}

			return dir, image.Point{-1, -1}
		}

		for {
			direction, p = getNextPoint(direction, p)

			if p.X == -1 && p.Y == -1 {
				break
			}

			if (lines[p.Y][p.X] >= 'a' && lines[p.Y][p.X] <= 'z') || (lines[p.Y][p.X] >= 'A' && lines[p.Y][p.X] <= 'Z') {
				nodes = append(nodes, string(lines[p.Y][p.X]))
				//nodes[string(lines[p.Y][p.X])] = p
			}

			steps++
		}
	}

	goToNextNode(down, point)

	fmt.Print("Part 1: ")
	for _, k := range nodes {
		fmt.Print(k)
	}
	fmt.Println()
	fmt.Println("Part 2:", steps)
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
