package main

import (
	"fmt"

	"github.com/kamillo/aoc/fileutil"
)

func main() {
	lines := fileutil.GetLines("input.txt")

	group := ""
	personsCount, part1, part2 := 0, 0, 0

	for _, line := range lines {
		if len(line) == 0 {
			answers := make(map[string]int)

			for _, answ := range group {
				answers[string(answ)]++
				if answers[string(answ)] == personsCount {
					part2++
				}
			}
			part1 += len(answers)
			group = ""
			personsCount = 0

		} else {
			group += line
			personsCount++
		}
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
