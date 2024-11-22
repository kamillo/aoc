package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	total := 0
	for _, line := range lines {
		fields := strings.Fields(line)

		score := 0

		switch fields[1] {
		case "X":
			score = 1
			switch fields[0] {
			case "A":
				score += 3
			case "B":
				score += 0
			case "C":
				score += 6
			}
		case "Y":
			score = 2
			switch fields[0] {
			case "A":
				score += 6
			case "B":
				score += 3
			case "C":
				score += 0
			}
		case "Z":
			score = 3
			switch fields[0] {
			case "A":
				score += 0
			case "B":
				score += 6
			case "C":
				score += 3
			}
		}

		total += score
	}

	fmt.Println("Part 1:", total)

	total = 0
	for _, line := range lines {
		fields := strings.Fields(line)

		score := 0

		switch fields[1] {
		case "X":
			score = 0
			switch fields[0] {
			case "A":
				score += 3
			case "B":
				score += 1
			case "C":
				score += 2
			}
		case "Y":
			score = 3
			switch fields[0] {
			case "A":
				score += 1
			case "B":
				score += 2
			case "C":
				score += 3
			}
		case "Z":
			score = 6
			switch fields[0] {
			case "A":
				score += 2
			case "B":
				score += 3
			case "C":
				score += 1
			}
		}

		total += score
	}

	fmt.Println("Part 2:", total)
}
