package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	lanterfish := utils.ToIntArr(lines[0], ",")
	population := map[int]int{}

	for _, f := range lanterfish {
		population[f]++
	}

	fmt.Println("Part 1: ", run(population, 80))
	fmt.Println("Part 2: ", run(population, 256))
}

func run(population map[int]int, days int) (count int) {
	for i := 0; i < days; i++ {
		newPopulation := map[int]int{}
		for k, v := range population {
			newPopulation[k-1] = v
		}

		if population[0] > 0 {
			if newPopulation[6] > 0 {
				newPopulation[6] += newPopulation[-1]
			} else {
				newPopulation[6] = newPopulation[-1]
			}
		}
		newPopulation[8] = population[0]

		for k := range newPopulation {
			if k <= -1 {
				delete(newPopulation, k)
			}
		}

		population = newPopulation
	}

	for _, v := range population {
		count += v
	}

	return
}
