package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	getStartMark := func(length int) int {
		for i := 0; i < len(lines[0])-length; i++ {
			chars := map[rune]bool{}

			for _, c := range lines[0][i : i+length] {
				chars[c] = true
			}

			if len(chars) == length {
				return i + length
			}
		}

		return -1
	}

	fmt.Println("Part 1:", getStartMark(4))
	fmt.Println("Part 2:", getStartMark(14))
}
