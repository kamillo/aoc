package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Range struct {
	start int
	end   int
}

func main() {
	lines := utils.GetLines("input.txt")

	part1 := 0
	part2 := 0

	for _, line := range lines {
		inputRanges := strings.Split(line, ",")
		for _, r := range inputRanges {
			rangeParts := strings.Split(r, "-")
			start := utils.JustAtoi(rangeParts[0])
			end := utils.JustAtoi(rangeParts[1])

			for i := start; i <= end; i++ {
				if checkPeriodic(i, 2) {
					part1 += i
				}

				if checkPeriodic(i, len(strconv.Itoa(i))) {
					part2 += i
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func checkPeriodic(num int, maxModulo int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	for i := 2; i <= maxModulo; i++ {
		if n%i != 0 {
			continue
		}

		partLen := n / i
		if s[:n-partLen] == s[partLen:] {
			return true
		}
	}
	return false
}
