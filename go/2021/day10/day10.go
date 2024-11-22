package main

import (
	"fmt"
	"sort"

	"github.com/kamillo/aoc/utils"
)

func main() {
	close := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	points := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	autocomplete := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

	score := 0
	score2 := []int{}
	for _, line := range utils.GetLines("input.txt") {
		open := []rune{}

		corrupted := false

		for _, c := range line {
			corrupted = false

			switch c {
			case '(', '{', '[', '<':
				open = append(open, close[c])
			default:
				if open[len(open)-1] != c {
					score += points[c]
					corrupted = true
					break

				} else {
					open = open[:len(open)-1]
				}
			}

			if corrupted {
				break
			}
		}

		for i, j := 0, len(open)-1; i < j; i, j = i+1, j-1 {
			open[i], open[j] = open[j], open[i]
		}

		if !corrupted {
			s := 0

			for _, c := range open {
				s = s*5 + autocomplete[c]
			}
			score2 = append(score2, s)
		}

	}
	sort.Ints(score2)
	fmt.Println("Part 1: ", score)
	fmt.Println("Part 2: ", score2[len(score2)/2])
}
