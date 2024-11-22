package main

import (
	"fmt"
	"math"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	fmt.Println("Part 1:", p1(lines))
	fmt.Println("Part 2:", p2(lines))
}

func p1(lines []string) int {
	points := map[utils.PointD3D]bool{}
	for _, line := range lines {
		p := utils.PointD3D{}
		fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		points[p] = true
	}

	offsets := []utils.PointD3D{
		utils.NewPointD3D(0, 0, 1),
		utils.NewPointD3D(0, 1, 0),
		utils.NewPointD3D(1, 0, 0),
		utils.NewPointD3D(0, 0, -1),
		utils.NewPointD3D(0, -1, 0),
		utils.NewPointD3D(-1, 0, 0),
	}

	res := len(points) * 6
	for p := range points {
		for _, o := range offsets {
			if points[p.Add(o)] {
				res--
			}
		}
	}

	return res
}

func p2(lines []string) int {
	points := map[utils.PointD3D]int{}
	minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	maxX, maxY, maxZ := 0, 0, 0

	for _, line := range lines {
		p := utils.PointD3D{}
		fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		points[p] = 0

		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Z < minZ {
			minZ = p.Z
		}

		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Z > maxZ {
			maxZ = p.Z
		}
	}

	minCords := utils.PointD3D{X: minX - 1, Y: minY - 1, Z: minZ - 1}
	maxCords := utils.PointD3D{X: maxX + 1, Y: maxY + 1, Z: maxZ + 1}

	offsets := []utils.PointD3D{
		utils.NewPointD3D(0, 0, 1),
		utils.NewPointD3D(0, 1, 0),
		utils.NewPointD3D(1, 0, 0),
		utils.NewPointD3D(0, 0, -1),
		utils.NewPointD3D(0, -1, 0),
		utils.NewPointD3D(-1, 0, 0),
	}

	start := Stack{minCords}
	visited := map[utils.PointD3D]bool{}

	for len(start) > 0 {
		u, _ := start.Pop()
		p := u.(utils.PointD3D)

		if visited[p] {
			continue
		}
		visited[p] = true

		for _, o := range offsets {
			v := p.Add(o)
			if minCords.X <= v.X && v.X <= maxCords.X &&
				minCords.Y <= v.Y && v.Y <= maxCords.Y &&
				minCords.Z <= v.Z && v.Z <= maxCords.Z {
				if _, ok := points[v]; ok {
					points[v] += 1

				} else {
					start.Push(v)
				}
			}
		}
	}

	sum := 0
	for _, v := range points {
		sum += v
	}

	return sum
}

type Stack []any

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val any) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() (any, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.

		return element, true
	}
}

func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return -1
	} else {
		index := len(*s) - 1
		element := (*s)[index]

		return element
	}
}
