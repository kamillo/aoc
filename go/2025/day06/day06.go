package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	part1 := 0
	part2 := 0

	numbers := [][]int{}
	for y, line := range lines {
		split := strings.Fields(line)

		if split[0] == "+" || split[0] == "*" {
			for i, s := range split {
				result := 0
				for j := 0; j < len(numbers); j++ {
					if s == "+" {
						result += numbers[j][i]
					} else {
						if result == 0 {
							result = numbers[j][i]
						} else {
							result *= numbers[j][i]
						}
					}
				}
				part1 += result
			}
		} else {
			numbers = append(numbers, []int{})
			for _, s := range split {
				numbers[y] = append(numbers[y], utils.JustAtoi(s))
			}
		}
	}

	numbers2 := []int{}
	for i := len(lines[0]) - 1; i >= 0; i-- {
		number := 0
		for l := 0; l < len(lines)-1; l++ {
			if lines[l][i] != ' ' {
				number *= 10
				number += utils.JustAtoi(string(lines[l][i]))
			}
		}
		if number != 0 {
			numbers2 = append(numbers2, number)
		}

		fmt.Printf("%d\n", numbers2)
		if lines[len(lines)-1][i] == '*' {
			calc := 1
			for _, n := range numbers2 {
				calc *= n
			}
			part2 += calc
			numbers2 = []int{}
		} else if lines[len(lines)-1][i] == '+' {
			calc := 0
			for _, n := range numbers2 {
				calc += n
			}
			part2 += calc
			numbers2 = []int{}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
