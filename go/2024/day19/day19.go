package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	// lines := utils.GetLines("test.txt")

	patterns := map[string]bool{}
	patternsArr := strings.Split(lines[0], ", ")
	for _, p := range patternsArr {
		patterns[p] = true
	}

	lines = lines[2:]

	part1, part2 := 0, 0

	for _, line := range lines {
		ok := countPatterns(line, patternsArr)
		fmt.Println(line, ok)
		if ok > 0 {
			part1++
			part2 += ok
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func countPatterns(target string, substrings []string) int {
	dp := make([]int, len(target)+1)
	dp[0] = 1

	for i := 1; i <= len(target); i++ {
		for _, substr := range substrings {
			if i >= len(substr) {
				prevIndex := i - len(substr)
				matchingPart := target[prevIndex:i]

				if matchingPart == substr && dp[prevIndex] > 0 {
					dp[i] += dp[prevIndex]
				}
			}
		}
	}

	return dp[len(target)]
}
