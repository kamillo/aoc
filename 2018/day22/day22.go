package main

import (
	"fmt"
)

// const targetX = 10
// const targetY = 10
// const depth = 510

const depth = 3198
const targetX = 12
const targetY = 757
const buffer = 50

var weights [targetY + buffer][targetX + buffer][3]int
var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type Entry struct {
	x, y     int
	gear     int
	previous *Entry
}

const (
	NOTHING int = iota
	TORCH
	CLIMBING_GEAR
)

func (e *Entry) Weight() int {
	return weights[e.y][e.x][e.gear]
}

var tileChars = []string{".", "=", "|", "#"}

func main() {
	mod := 20183
	sumRisk := 0

	cave := [targetY + buffer][targetX + buffer]int{}

	for y := range cave {
		for x := range cave[y] {
			index := 0
			if x != 0 && y != 0 {
				index = cave[y][x-1] * cave[y-1][x]
			} else if x > 0 && y == 0 {
				index = x * 16807
			} else if y > 0 && x == 0 {
				index = y * 48271
			}

			erosionLevel := (index + depth) % mod
			if x == targetX && y == targetY {
				erosionLevel = depth % mod
			}

			cave[y][x] = erosionLevel
			if !(x > targetX || y > targetY) {
				sumRisk += erosionLevel % 3
			}
		}
	}

	fmt.Println("Part 1:", sumRisk)

	weights[0][0][TORCH] = 1
	tiles := cave
	for y := range cave {
		for x := range cave[y] {
			tiles[y][x] %= 3
		}
	}

	queue := []*Entry{
		&Entry{0, 0, TORCH, nil},
	}

	for len(queue) > 0 {
		// Find the minimum weight
		min := -1
		for _, entry := range queue {
			if min < 0 || entry.Weight() < min {
				min = entry.Weight()
			}
		}

		minCount := 0
		for i, entry := range queue {
			if entry.Weight() == min {
				if entry.x == targetX && entry.y == targetY && entry.gear == TORCH {
					fmt.Println("Part 2:", entry.Weight()-1)
					return
				}

				for _, dir := range directions {
					x := entry.x + dir[0]
					y := entry.y + dir[1]
					if x < 0 || y < 0 {
						continue
					} else if y >= len(tiles) || x >= len(tiles[y]) {
						continue
					}

					if entry.gear != tiles[y][x] {
						// Current gear can be used on this tile
						posWeight := &weights[y][x][entry.gear]
						if *posWeight == 0 || *posWeight > entry.Weight()+1 {
							*posWeight = entry.Weight() + 1
							queue = append(queue, &Entry{x, y, entry.gear, entry})
						}
					}
				}
				// Try switching gear
				gear := 3 ^ (entry.gear ^ tiles[entry.y][entry.x])
				posWeight := &weights[entry.y][entry.x][gear]
				if *posWeight == 0 || *posWeight > entry.Weight()+7 {
					*posWeight = entry.Weight() + 7
					queue = append(queue, &Entry{entry.x, entry.y, gear, entry})
				}

				// Move the processed entries to the beginning so they can be quickly removed
				queue[minCount], queue[i] = queue[i], queue[minCount]
				minCount++
			}
		}
		queue = queue[minCount:]
	}
}

func printCave(cave [][]int) {
	for _, row := range cave {
		for _, cell := range row {
			switch cell % 3 {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("=")
			case 2:
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}
