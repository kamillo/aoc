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
	// part2 := 0

	ranges := make([]Range, 0)

	for _, line := range lines {
		inputRanges := strings.Split(line, ",")
		for _, r := range inputRanges {
			rangeParts := strings.Split(r, "-")
			start := utils.JustAtoi(rangeParts[0])
			end := utils.JustAtoi(rangeParts[1])

			ranges = append(ranges, Range{
				start: start,
				end:   end,
			})

			for i := start; i <= end; i++ {
				if checkRepeatedHalves(i) {
					part1 += i
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
}

func checkRepeatedHalves(num int) bool {
	s := strconv.Itoa(num)
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}
