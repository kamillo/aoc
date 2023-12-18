package main

import (
	"fmt"
	"image"
	"math"
	"sort"

	"github.com/kamillo/aoc/utils"
)

var directions = map[string]image.Point{
	"U": image.Pt(0, -1),
	"R": image.Pt(1, 0),
	"D": image.Pt(0, 1),
	"L": image.Pt(-1, 0),
}

type DigSequence struct {
	Direction string
	Num       int
}

func main() {
	lines := utils.GetLines("input.txt")

	directions2 := []string{"R", "D", "L", "U"}
	digPlan1 := []DigSequence{}
	digPlan2 := []DigSequence{}
	for _, line := range lines {
		var dir string
		var off int
		var intDir int
		var hexOff int
		fmt.Sscanf(line, "%s %d (#%5x%1d)", &dir, &off, &hexOff, &intDir)

		digPlan1 = append(digPlan1, DigSequence{dir, off})
		digPlan2 = append(digPlan2, DigSequence{directions2[intDir], hexOff})
	}

	fmt.Println("Part 1:", calculateArea(digPlan1))
	fmt.Println("Part 2:", calculateArea(digPlan2))
}

func calculateArea(digSequence []DigSequence) int {
	x, y := 0, 0
	corners := []image.Point{{x, y}}
	xPoints := map[int]bool{}
	yPoints := map[int]bool{}

	for _, dig := range digSequence {
		switch dig.Direction {
		case "U":
			y -= dig.Num
		case "D":
			y += dig.Num
		case "R":
			x += dig.Num
		default:
			x -= dig.Num
		}

		corners = append(corners, image.Point{x, y})
		xPoints[x] = true
		yPoints[y] = true
	}

	xPointsSorted := sortMapKeys(xPoints)
	yPointsSorted := sortMapKeys(yPoints)

	compression := map[image.Point]image.Point{}
	gridSizes := map[image.Point]int{}

	newX := 0
	for ix, x := range xPointsSorted {
		newY := 0
		for iy, y := range yPointsSorted {
			compression[image.Point{x, y}] = image.Point{newX, newY}
			gridSizes[image.Point{newX, newY}] = 1

			if ix > 0 {
				lastX := xPointsSorted[ix-1]
				gridSizes[image.Point{newX - 1, newY}] = x - lastX - 1
				if iy > 0 {
					lastY := yPointsSorted[iy-1]
					gridSizes[image.Point{newX - 1, newY - 1}] = (x - lastX - 1) * (y - lastY - 1)
				}
			}
			if iy > 0 {
				lastY := yPointsSorted[iy-1]
				gridSizes[image.Point{newX, newY - 1}] = y - lastY - 1
			}

			newY += 2
		}
		newX += 2
	}

	fmt.Println(len(corners))

	seen := map[image.Point]bool{}
	cornersPrim := corners[1:]
	cornersPrim = append(cornersPrim, corners[0])
	for i := 0; i < len(corners); i++ {
		e0, e1 := corners[i], cornersPrim[i]
		x0, y0 := compression[e0].X, compression[e0].Y
		x1, y1 := compression[e1].X, compression[e1].Y

		if x0 > x1 {
			x0, x1 = x1, x0
		}
		if y0 > y1 {
			y0, y1 = y1, y0
		}

		for x := x0; x <= x1; x++ {
			for y := y0; y <= y1; y++ {
				seen[image.Point{x, y}] = true
			}
		}
	}

	fmt.Println(len(seen))

	minX, maxX := minMaxMapKeys(seen, "x")
	minY, maxY := minMaxMapKeys(seen, "y")

	fmt.Println(minX, maxX, minY, maxY)

	fill := map[image.Point]bool{}
	for k, v := range seen {
		fill[k] = v
	}
	for y := minY; y <= maxY; y++ {
		isEnclosed := false
		for x := minX; x <= maxX; x++ {
			if seen[image.Point{x, y}] {
				if seen[image.Point{x, y - 1}] && seen[image.Point{x, y + 1}] {
					isEnclosed = !isEnclosed
				}
				if seen[image.Point{x, y + 1}] && seen[image.Point{x + 1, y}] {
					isEnclosed = !isEnclosed
				}
				if seen[image.Point{x, y + 1}] && seen[image.Point{x - 1, y}] {
					isEnclosed = !isEnclosed
				}
			} else {
				if isEnclosed {
					fill[image.Point{x, y}] = true
				}
			}
		}
	}

	total := 0
	for k := range fill {
		total += gridSizes[k]
	}

	return total
}

func sortMapKeys(m map[int]bool) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func minMaxMapKeys(m map[image.Point]bool, axis string) (int, int) {
	min := math.MaxInt32
	max := 0
	for k := range m {
		if axis == "x" {
			if k.X < min {
				min = k.X
			}
			if k.X > max {
				max = k.X
			}
		} else {
			if k.Y < min {
				min = k.Y
			}
			if k.Y > max {
				max = k.Y
			}
		}
	}
	return min, max
}
