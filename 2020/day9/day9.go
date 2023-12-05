package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strconv"
)

func main() {
	lines := utils.GetLines("input.txt")

	preambleLength := 25
	numbers := make([]int, len(lines))
	invalid := 0
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)

		if i >= preambleLength {
			found := false
			numbersToCheck := numbers[i-preambleLength : i]

			for j, num := range numbersToCheck {
				for k, num2 := range numbersToCheck {
					if j != k && (num+num2 == numbers[i]) {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				invalid = numbers[i]
				fmt.Println("Part 1: ", numbers[i])
			}
		}
	}

	// Part 2
	for begin, _ := range numbers {
		for end := begin; end < len(numbers); end++ {
			sum := 0
			subrange := numbers[begin:end]
			for _, num := range subrange {
				sum += num
			}
			if sum == invalid {
				fmt.Println("Part 2: ", utils.MaxInArray(subrange)+utils.MinInArray(subrange))
				return
			}
		}
	}
}
