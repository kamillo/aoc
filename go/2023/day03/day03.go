package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/kamillo/aoc/utils"
)

type set map[int]bool

func main() {
	lines := utils.GetLines("input.txt")
	engines := map[string]set{}
	sum := 0

	for y, line := range lines {
		isAdj := false
		level := 1
		digit := 0
		stars := []string{}

		for x := len(line) - 1; x >= 0; x-- {
			if d, err := strconv.Atoi(string(line[x])); err == nil {
				digit += level * d
				level *= 10

				adj, star := CountAdj(y, x, lines, '.', '*')
				isAdj = isAdj || adj > 0
				stars = append(stars, star...)

				if x > 0 {
					continue
				}
			}

			// not digit or first char
			if isAdj {
				sum += digit
			}

			for _, s := range stars {
				if _, ok := engines[s]; !ok {
					engines[s] = map[int]bool{}
				}
				engines[s][digit] = true
			}

			isAdj = false
			level = 1
			digit = 0
			stars = []string{}
		}
	}

	fmt.Println("Part 1:", sum)

	sum = 0
	for _, digits := range engines {
		if len(digits) == 2 {
			ratio := 1
			for d := range digits {
				ratio *= d
			}
			sum += ratio
		}
	}

	fmt.Println("Part 2:", sum)
}

func CountAdj(x int, y int, grid []string, char byte, star byte) (adj int, stars []string) {
	adjPoints := [][2]int{
		{x + 1, y},
		{x, y + 1},
		{x + 1, y + 1},
		{x - 1, y},
		{x, y - 1},
		{x - 1, y - 1},
		{x - 1, y + 1},
		{x + 1, y - 1},
	}

	var check = func(xp int, yp int) {
		if xp >= 0 && yp >= 0 && yp < len(grid) && xp < len(grid) && !unicode.IsDigit(rune(grid[xp][yp])) && grid[xp][yp] != char {
			adj++
			if grid[xp][yp] == star {
				stars = append(stars, fmt.Sprintf("%d-%d", xp, yp))
			}
		}
	}

	for _, p := range adjPoints {
		check(p[0], p[1])
	}

	return adj, stars
}
