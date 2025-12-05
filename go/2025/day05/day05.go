package main

import (
	"fmt"
	"sort"
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

	ranges := make([]Range, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		rangeParts := strings.Split(line, "-")

		if len(rangeParts) == 2 {
			ranges = append(ranges, Range{utils.JustAtoi(rangeParts[0]), utils.JustAtoi(rangeParts[1])})
		} else {
			id := utils.JustAtoi(line)

			for _, r := range ranges {
				if id >= r.start && id <= r.end {
					part1++
					break
				}
			}
		}
	}

	part2 = countUniqueValues(ranges)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func countUniqueValues(ranges []Range) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	count := 0
	currentStart := ranges[0].start
	currentEnd := ranges[0].end

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]

		if next.start > currentEnd+1 {
			count += currentEnd - currentStart + 1
			currentStart = next.start
			currentEnd = next.end
		} else {
			if next.end > currentEnd {
				currentEnd = next.end
			}
		}
	}

	count += currentEnd - currentStart + 1

	return count
}
