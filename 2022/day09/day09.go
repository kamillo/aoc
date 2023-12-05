package main

import (
	"fmt"
	"image"
	"math"

	"github.com/kamillo/aoc/utils"
)

var offsets = map[string]image.Point{
	"U": image.Pt(0, -1), // U
	"R": image.Pt(1, 0),  // R
	"D": image.Pt(0, 1),  // D
	"L": image.Pt(-1, 0), // L
}

func main() {
	lines := utils.GetLines("input.txt")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	tail := image.Pt(0, 0)
	head := image.Pt(0, 0)

	visited := map[image.Point]bool{}

	for _, line := range lines {
		direction := ""
		steps := 0
		fmt.Sscanf(line, "%s %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			head = head.Add(offsets[direction])
			tail = move(head, tail)
			visited[tail] = true
		}
	}

	fmt.Println("Part 1:", len(visited))
}

func part2(lines []string) {
	tail := []image.Point{}
	for i := 0; i < 10; i++ {
		tail = append(tail, image.Pt(0, 0))
	}
	visited := map[image.Point]bool{}

	for _, line := range lines {
		direction := ""
		steps := 0
		fmt.Sscanf(line, "%s %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			tail[0] = tail[0].Add(offsets[direction])

			for tI := 1; tI < len(tail); tI++ {
				tail[tI] = move(tail[tI-1], tail[tI])
			}
			visited[tail[9]] = true
		}
	}

	fmt.Println("Part 2:", len(visited))
}

func move(h, t image.Point) image.Point {
	s := h.Sub(t)
	absX := math.Abs(float64(s.X))
	absY := math.Abs(float64(s.Y))

	if absX > 1 || absY > 1 {
		if absX > absY {
			if s.X < 0 {
				s.X += 1
			} else {
				s.X -= 1
			}
		} else if absY > absX {
			if s.Y < 0 {
				s.Y += 1
			} else {
				s.Y -= 1
			}
		} else {
			if s.Y < 0 {
				s.Y += 1
			} else {
				s.Y -= 1
			}
			if s.X < 0 {
				s.X += 1
			} else {
				s.X -= 1
			}
		}

		t = t.Add(s)
	}

	return t
}

func print(rope []image.Point) {
	for y := -10; y < 10; y++ {
		for x := -10; x < 10; x++ {
			f := true
			for i, r := range rope {
				if r.X == x && r.Y == y {
					fmt.Print(i)
					f = false
					break
				}
			}
			if f {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
