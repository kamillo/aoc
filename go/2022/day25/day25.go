package main

import (
	"fmt"
	"math"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	sum := 0
	for _, line := range lines {
		i := convertSnafuToDecimal(line)
		sum += i
	}

	fmt.Println("Part 1:", convertDecimalToSnafu(sum))
}

func convertDecimalToSnafu(decimal int) string {
	switch decimal {
	case -2:
		return "="
	case -1:
		return "-"
	case 0:
		return "0"
	}
	snafu := ""
	for decimal > 0 {
		digit := decimal % 5
		if digit > 2 {
			decimal += 5
		}
		snafu = []string{"0", "1", "2", "=", "-"}[digit] + snafu
		decimal = decimal / 5
	}
	return snafu
}

func convertSnafuToDecimal(five string) int {
	fiveValues := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'=': -2,
		'-': -1,
	}

	result := 0

	for i, ch := range five {
		result += fiveValues[ch] * int(math.Pow(5, float64(len(five)-1-i)))
	}

	return result
}
