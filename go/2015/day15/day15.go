package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
)

func main() {
	ingredients := make([][]int, 0)
	for _, line := range utils.GetLines("input.txt") {
		capacity, durability, flavor, texture, calories := 0, 0, 0, 0, 0
		ingredient := ""
		fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &ingredient,
			&capacity,
			&durability,
			&flavor,
			&texture,
			&calories)

		ingredients = append(ingredients, []int{capacity, durability, flavor, texture, calories})
	}

	max := 0
	max2 := 0
	for a := 1; a <= 100; a++ {
		for b := 1; b <= 100-a; b++ {
			for c := 1; c <= 100-(a+b); c++ {
				d := 100 - (a + b + c)
				if a+b+c+d != 100 {
					break
				}
				result := 1
				caloriesTotal := 0
				amounts := [4]int{a, b, c, d}

				for i := 0; i < 4; i++ {
					prop := 0
					for j := 0; j < len(ingredients); j++ {
						prop += amounts[j] * ingredients[j][i]
					}
					if prop < 0 {
						result = 0
						break
					}
					result *= prop
				}
				for j := 0; j < len(ingredients); j++ {
					caloriesTotal += amounts[j] * ingredients[j][4]
				}
				if result > max {
					max = result
				}
				if result > max2 && caloriesTotal == 500 {
					max2 = result
				}
			}
		}
	}

	fmt.Println("Part1: ", max)
	fmt.Println("Part2: ", max2)
}
