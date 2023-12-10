package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	start := utils.Point2D{X: 0, Y: 0}
	tiles := make([][]rune, len(lines))
	for y, line := range lines {
		tiles[y] = []rune(line)
		if strings.Contains(line, "S") {
			start = utils.Point2D{X: strings.Index(line, "S"), Y: y}
		}
	}

	// TODO: translate starting point
	tiles[start.Y][start.X] = '|'
	// tiles[start.Y][start.X] = '7'

	loop := map[utils.Point2D]bool{}
	current := start
	for {
		n := getNeighbours(current, tiles[current.Y][current.X])
		if !loop[n[0]] {
			current = n[0]
			loop[n[0]] = true
		} else if !loop[n[1]] {
			current = n[1]
			loop[n[1]] = true
		} else {
			break
		}
	}

	fmt.Println("Part 1:", len(loop)/2)

	knees := map[rune]bool{
		'F': true,
		'7': true,
		'L': true,
		'J': true,
	}

	enclosed := 0
	for y := range tiles {
		count := 0
		prev := '.'
		for x := range tiles[y] {
			tile := tiles[y][x]

			if loop[utils.Point2D{X: x, Y: y}] {
				if tile == '|' ||
					(tile == '7' && prev == 'L') ||
					(tile == 'J' && prev == 'F') {
					count++
				}

				if knees[tile] {
					prev = tile
				}

			} else {
				tiles[y][x] = '.'
				if count%2 == 1 {
					enclosed++
					tiles[y][x] = '*'
				}

				prev = '.'
			}
		}
	}

	fmt.Println("Part 2:", enclosed)
}

func getNeighbours(pos utils.Point2D, pipe rune) []utils.Point2D {
	tileTypes := map[rune][]utils.Point2D{
		'|': {{X: 0, Y: 1}, {X: 0, Y: -1}},
		'-': {{X: 1, Y: 0}, {X: -1, Y: 0}},
		'F': {{X: 1, Y: 0}, {X: 0, Y: 1}},
		'7': {{X: -1, Y: 0}, {X: 0, Y: 1}},
		'L': {{X: 1, Y: 0}, {X: 0, Y: -1}},
		'J': {{X: -1, Y: 0}, {X: 0, Y: -1}},
		'.': {{X: 0, Y: 0}, {X: 0, Y: 0}},
	}

	return []utils.Point2D{pos.Add(tileTypes[pipe][0]), pos.Add(tileTypes[pipe][1])}
}
