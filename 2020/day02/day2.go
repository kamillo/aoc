package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strings"
)

func main() {
	lines := utils.GetLines("input.txt")

	goodPasswords := 0
	goodPasswordsPart2 := 0
	for _, line := range lines {
		min, max := 0, 0
		var char string
		var password string

		fmt.Sscanf(line, "%d-%d %1s: %s", &min, &max, &char, &password)

		count := strings.Count(password, char)
		if count >= min && count <= max {
			goodPasswords++
		}

		// Part2
		valid := false
		if string(password[min-1]) == char {
			valid = true
		}
		if string(password[max-1]) == char {
			valid = !valid && true
		}

		if valid {
			goodPasswordsPart2++
		}
	}

	fmt.Println("Part 1: ", goodPasswords)
	fmt.Println("Part 2: ", goodPasswordsPart2)
}
