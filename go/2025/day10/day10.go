package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	part1 := 0

	// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
	for l, line := range lines {
		target := []int{}
		split := strings.Split(line, " ")

		for _, char := range split[0][1 : len(split[0])-1] {
			if char == '#' {
				target = append(target, 1)
			} else {
				target = append(target, 0)
			}
		}

		buttons := []map[int]bool{}
		for _, s := range split[1 : len(split)-1] {
			button := map[int]bool{}

			numbers := strings.Split(s[1:len(s)-1], ",")
			for _, number := range numbers {
				button[utils.JustAtoi(number)] = true
			}

			buttons = append(buttons, button)
		}

		min := math.MaxInt
		for _, combination := range utils.CombinationsWithRepetitions(10, buttons) {
			// if combination == target {
			// 	fmt.Println("Found")
			// }

			result := make([]int, len(target))
			for i, button := range combination {
				for k := range button {
					if result[k] == 1 {
						result[k] = 0
					} else {
						result[k] = 1
					}
				}

				matches := true
				for k := range result {
					if result[k] != target[k] {
						matches = false
						break
					}
				}

				if matches {
					min = utils.Min(min, i)
					// fmt.Printf("Found %d %d\n", l, i)
					// fmt.Println(result)
					// fmt.Println(target)
					break
				}
			}
		}

		part1 += min + 1
		fmt.Printf("%d: %d\n", l, min)

		// fmt.Println(buttons)
		// fmt.Println(target)
	}

	fmt.Println("Part 1:", part1)
}
