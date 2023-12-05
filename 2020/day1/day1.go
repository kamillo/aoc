package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strconv"
)

func main() {
	lines := utils.GetLines("input.txt")
	first, second := 0.0, 0.0

	numbers := make([]float64, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.ParseFloat(line, 64)
	}

	for i, number1 := range numbers {
		for j, number2 := range numbers {
			if i != j && first == 0 {
				if number1+number2 == 2020 {
					first = number1 * number2
				}
			}

			for k, number3 := range numbers {
				if i != j && j != k && i != k {
					if number1+number2+number3 == 2020 {
						second = number1 * number2 * number3
						break
					}
				}
			}
			if second != 0 && first != 0 {
				break
			}
		}
		if second != 0 && first != 0 {
			break
		}
	}
	fmt.Printf("First part: %f\n", first)
	fmt.Printf("Second part: %f\n", second)
}
