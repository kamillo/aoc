package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	sum := 0
	sum2 := 0

	cards := make([]int, len(lines))
	for i := 0; i < len(cards); i++ {
		cards[i] = 1
	}

	for i, line := range lines {
		val := 0
		n := strings.Split(line, "|")
		winning := utils.ToIntSet(n[0], " ")
		scratched := utils.ToIntArr(n[1], " ")

		for _, card := range scratched {
			if winning[card] {
				val++
			}
		}

		if val > 0 {
			sum += int(math.Pow(2, float64(val-1)))
			for j := 0; j < val; j++ {
				cards[i+1+j] += (cards[i])
			}
		}

		sum2 += cards[i]
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}
