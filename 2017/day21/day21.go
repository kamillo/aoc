package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	makeArt := func(iterations int) int {
		art := []string{".#.", "..#", "###"}
		rules := map[string][]string{}

		for _, line := range utils.GetLines("input.txt") {
			pattern := strings.Split(strings.Split(line, " => ")[0], "/")
			out := strings.Split(strings.Split(line, " => ")[1], "/")

			for i := 0; i < 4; i++ {
				rules[strings.Join(utils.SliceFlipVString(pattern), "")] = out
				rules[strings.Join(utils.SliceFlipHString(pattern), "")] = out
				pattern = utils.SliceRotateString(pattern)
				rules[strings.Join(pattern, "")] = out
			}
		}

		for i := 0; i < iterations; i++ {
			size := len(art)
			div := 0
			if size%2 == 0 {
				div = 2
			} else if size%3 == 0 {
				div = 3
			}

			newArt := make([]string, size/div*(div+1))
			apply(art, newArt, div, rules)

			art = newArt
			// for a := range art {
			// 	fmt.Println(art[a])
			// }
			// fmt.Println()
		}

		sum := 0
		for _, s := range art {
			sum += strings.Count(s, "#")
		}

		return sum
	}

	fmt.Println("Part 1:", makeArt(5))
	fmt.Println("Part 2:", makeArt(18))
}

func apply(art, newArt []string, div int, rules map[string][]string) {
	size := len(art)
	for x := 0; x < size; x += div {
		lastY := 0
		for y := 0; y < size; y += div {
			pattern := art[y][x:x+div] + art[y+1][x:x+div]
			if div == 3 {
				pattern += art[y+2][x : x+div]
			}

			rule, ok := rules[pattern]
			if !ok {
				panic("no rule")
			}

			for r := range rule {
				newArt[lastY+r] += rule[r]
			}

			lastY += div + 1
		}
	}
}
