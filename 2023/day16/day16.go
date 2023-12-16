package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/kamillo/aoc/utils"
)

var directions = map[string]image.Point{
	"N": image.Pt(0, -1),
	"E": image.Pt(1, 0),
	"S": image.Pt(0, 1),
	"W": image.Pt(-1, 0),
}

var energized = [][]string{}
var tile = [][]rune{}

func main() {
	lines := utils.GetLines("input.txt")

	for _, line := range lines {
		tile = append(tile, []rune(line))
	}

	initEnergized()
	beamMeUpScotty(image.Pt(0, 0), directions["E"])
	fmt.Println("Part 1:", calcEnergy())

	energy := []int{}
	// start from every edge of the tile
	for x := 0; x < len(tile[0]); x++ {
		for direction, pt := range map[string]image.Point{"S": image.Pt(x, 0), "N": image.Pt(x, len(tile)-1)} {
			initEnergized()
			beamMeUpScotty(pt, directions[direction])
			energy = append(energy, calcEnergy())
		}
	}

	for y := 0; y < len(tile); y++ {
		for direction, pt := range map[string]image.Point{"E": image.Pt(0, y), "W": image.Pt(len(tile[0])-1, y)} {
			initEnergized()
			beamMeUpScotty(pt, directions[direction])
			energy = append(energy, calcEnergy())
		}
	}

	fmt.Println("Part 2:", utils.MaxInArray(energy))
}

func beamMeUpScotty(start image.Point, direction image.Point) {

	for {
		//check if start point is in range of tile
		if start.X < 0 || start.X >= len(tile[0]) || start.Y < 0 || start.Y >= len(tile) {
			return
		}

		// check if we are in the loop
		if strings.Contains(energized[start.Y][start.X], directionString(direction)) {
			return
		}

		energized[start.Y][start.X] += directionString(direction)

		x := start.X
		y := start.Y

		switch tile[start.Y][start.X] {
		case '|': // split vertically
			if direction == directions["E"] || direction == directions["W"] {
				beamMeUpScotty(image.Pt(x, y-1), directions["N"])
				direction = directions["S"]
			}

		case '-': // split horizontally
			if direction == directions["N"] || direction == directions["S"] {
				beamMeUpScotty(image.Pt(x+1, y), directions["E"])
				direction = directions["W"]
			}

		case '\\': // reflect 90 degrees counterclockwise
			if direction == directions["N"] {
				direction = directions["W"]
			} else if direction == directions["E"] {
				direction = directions["S"]
			} else if direction == directions["S"] {
				direction = directions["E"]
			} else if direction == directions["W"] {
				direction = directions["N"]
			}

		case '/': // reflect 90 degrees clockwise
			if direction == directions["N"] {
				direction = directions["E"]
			} else if direction == directions["E"] {
				direction = directions["N"]
			} else if direction == directions["S"] {
				direction = directions["W"]
			} else if direction == directions["W"] {
				direction = directions["S"]
			}
		}

		start = start.Add(direction)
	}
}

func directionString(direction image.Point) string {
	for k, v := range directions {
		if v == direction {
			return k
		}
	}
	return ""
}

func calcEnergy() int {
	sum := 0
	for y := range energized {
		for x := range energized[y] {
			if energized[y][x] != "" {
				sum++
			}

		}
	}

	return sum
}

func initEnergized() {
	energized = make([][]string, len(tile))
	for y := range tile {
		energized[y] = make([]string, len(tile[y]))
	}
}
