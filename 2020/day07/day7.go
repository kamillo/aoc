package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strings"
)

type bag map[string]int

func main() {
	lines := utils.GetLines("input.txt")
	part1 := make(map[string]bool)

	bagsMap := make(map[string]bag)
	for _, line := range lines {
		splited := strings.Split(line, " contain ")
		outer := strings.TrimSuffix(splited[0], " bags") // don't need that
		inner := splited[1]
		var bagsInside bag

		if inner != "no other bags" {
			bags := strings.Split(inner, ", ")
			bagsInside = make(bag)
			for _, bag := range bags {
				amount := 0
				color1, color2 := "", ""
				fmt.Sscanf(bag, "%d %s %s", &amount, &color1, &color2)
				bagsInside[color1+" "+color2] = amount
			}
		}
		bagsMap[outer] = bagsInside
	}

	countBags("shiny gold", bagsMap, part1)
	fmt.Println("Part 1: ", len(part1))
	fmt.Println("Part 2: ", countBags2("shiny gold", bagsMap))
}

func countBags2(color string, bagsMap map[string]bag) int {
	sum := 0
	if color != "shiny gold" {
		sum = 1
	}

	for bag, amount := range bagsMap[color] {
		if count := countBags2(bag, bagsMap); count > 0 {
			sum += amount * count
		}
	}

	return sum
}

func countBags(color string, bagsMap map[string]bag, part1 map[string]bool) {
	for key, value := range bagsMap {
		if _, ok := value[color]; ok {
			part1[key] = true
			countBags(key, bagsMap, part1)
		}
	}
}
