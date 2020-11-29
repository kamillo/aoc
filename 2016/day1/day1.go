package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/fileutil"
	"github.com/kamillo/aoc/geometry"
)

func main() {
	lines := fileutil.GetLines("input.txt")
	// lines := []string{"R5, L5, R5, R3"}
	// lines := []string{"R8, R4, R4, R8"}
	x, y, dx, dy := 0.0, 0.0, 0.0, 1.0
	visited := make(map[string]bool)

	for _, line := range lines {
		splited := strings.Split(line, ", ")
		paths := make([]geometry.Line, len(splited))
		checkIntersection := true

		for index, cord := range splited {
			val, _ := strconv.ParseFloat(cord[1:], 64)

			if cord[0] == 'L' {
				if dx == 0 {
					if dy == 1 {
						dx = -1
					} else if dy == -1 {
						dx = 1
					}
					dy = 0
				} else if dy == 0 {
					if dx == 1 {
						dy = 1
					} else if dx == -1 {
						dy = -1
					}
					dx = 0
				}
			} else {
				if dx == 0 {
					if dy == 1 {
						dx = 1
					} else if dy == -1 {
						dx = -1
					}
					dy = 0
				} else if dy == 0 {
					if dx == 1 {
						dy = -1
					} else if dx == -1 {
						dy = 1
					}
					dx = 0
				}
			}

			newPath := geometry.Line{x, y, x + dx*val, y + dy*val}
			if index > 1 && checkIntersection {
				for _, path := range paths[:index-1] {
					x, y, ok := geometry.LineIntersection(path, newPath)
					if ok {
						fmt.Println("Part 2: ", math.Abs(x)+math.Abs(y))
						checkIntersection = false
					}
				}
			}
			paths[index] = newPath
			x += dx * val
			y += dy * val
			visited[fmt.Sprint(x)+fmt.Sprint(y)] = true
		}
	}
	fmt.Println("Part 1: ", math.Abs(float64(x))+math.Abs(float64(y)))
}
