package main

import (
	"fmt"
	"math"

	"github.com/kamillo/aoc/utils"
)

type Point4D struct {
	X int
	Y int
	Z int
	W int
}

func main() {
	lines := utils.GetLines("input.txt")
	points := map[Point4D]bool{}
	constallations := []map[Point4D]bool{}

	for _, line := range lines {
		p := Point4D{}

		fmt.Sscanf(line, "%d,%d,%d,%d", &p.X, &p.Y, &p.Z, &p.W)
		points[p] = true
	}

	for len(points) > 0 {
		constallations = append(constallations, make(map[Point4D]bool))

		for a := range points {
			constallations[len(constallations)-1][a] = true
			delete(points, a)
			break
		}

		for a := range points {
			found := false

			for _, c := range constallations {
				for b := range c {
					if a != b {
						d := distance(a, b)
						if d < 4 {
							found = true
							// fmt.Println(d, a, b)
							c[a] = true
							delete(points, a)
							break
						}
					}
					if found {
						break
					}
				}
				if found {
					break
				}
			}
		}
	}

	fmt.Println(len(constallations))
	fmt.Println(constallations)
}

func distance(a, b Point4D) int {
	return int(
		math.Abs(float64(a.X)-float64(b.X)) +
			math.Abs(float64(a.Y)-float64(b.Y)) +
			math.Abs(float64(a.Z)-float64(b.Z)) +
			math.Abs(float64(a.W)-float64(b.W)))
}
