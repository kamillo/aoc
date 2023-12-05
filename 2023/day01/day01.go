package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	sum := 0

	for _, line := range lines {
		digits := utils.ToIntArr(line, "")

		sum += digits[0]*10 + digits[len(digits)-1]
	}

	fmt.Println("Part 1:", sum)

	digitsAsLetters := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	sum = 0

	for _, line := range lines {
		digit1 := -1
		digit2 := -1

		for i := 0; i < len(line); i++ {
			if i, err := strconv.Atoi(string(line[i])); err == nil {
				digit1 = i
			}

			for d, digit := range digitsAsLetters {
				if strings.Contains(line[:i], digit) {
					digit1 = d
				}
			}

			if digit1 > -1 {
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if i, err := strconv.Atoi(string(line[i])); err == nil {
				digit2 = i
			}

			for d, digit := range digitsAsLetters {
				if strings.Contains(line[i:], digit) {
					digit2 = d
				}
			}

			if digit2 > -1 {
				break
			}
		}

		sum += digit1*10 + digit2
	}

	fmt.Println("Part 2:", sum)
}
