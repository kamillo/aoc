package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	packages := utils.GetLinesAsInterface("input.txt")
	sum := utils.Sum(packages)

	fmt.Println("Part 1: ", getQuantumEntanglement(sum/3, packages))
	fmt.Println("Part 2: ", getQuantumEntanglement(sum/4, packages))
}

func getQuantumEntanglement(groupLoad int, packages []interface{}) int {
	group1 := []int{}

	for i := 2; i <= len(packages); i++ {
		found := false

		for _, c := range utils.Combinations(packages, i) {
			if utils.Sum(c) == groupLoad {
				group1 = append(group1, utils.Product(c))
				found = true
			}
		}

		if found {
			return utils.MinInArray(group1)
		}
	}

	return 0
}
