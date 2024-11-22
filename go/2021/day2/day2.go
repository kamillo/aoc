package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	horizontal, depth1, depth2, aim := 0, 0, 0, 0

	for _, line := range utils.GetLines("input.txt") {
		command := ""
		value := 0
		fmt.Sscanf(line, "%s %d", &command, &value)

		switch command {
		case "forward":
			horizontal += value
			depth2 += aim * value
		case "down":
			depth1 += value
			aim += value
		case "up":
			depth1 -= value
			aim -= value
		}
	}

	fmt.Println("Part 1: ", horizontal*depth1)
	fmt.Println("Part 2: ", horizontal*depth2)
}
