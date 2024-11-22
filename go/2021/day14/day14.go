package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	template := lines[0]
	lines = lines[2:]

	rules := map[string]string{}
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		rules[split[0]] = split[1]
	}

	polymer := template

	grow := func(iterations int) int {
		counts := map[string]int{}

		for i := 2; i <= len(polymer); i++ {
			pair := polymer[i-2 : i]
			counts[pair]++
		}

		for j := 0; j < iterations; j++ {
			newCounts := map[string]int{}
			for pair, c := range counts {
				newCounts[string(pair[0])+rules[pair]] += c
				newCounts[rules[pair]+string(pair[1])] += c
			}

			counts = newCounts
		}

		letters := map[byte]int{}
		for pair, c := range counts {
			letters[pair[0]] += c
		}
		letters[polymer[len(polymer)-1]]++

		max := 0
		min := math.MaxInt64
		for _, l := range letters {
			if l < min {
				min = l
			}

			if l > max {
				max = l
			}
		}

		return max - min
	}

	fmt.Println("Part 1:", grow(10))
	fmt.Println("Part 2:", grow(40))
}
