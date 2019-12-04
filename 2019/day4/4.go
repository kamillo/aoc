package main

import (
	"fmt"
	"strconv"
)

func main() {
	start := 136818
	end := 685979

	possibilitiesPart1 := 0
	possibilitiesPart2 := 0

	for i := start; i < end; i++ {
		pass := strconv.Itoa(i)
		increasing := true
		matching := make(map[int]int)

		for j := range pass {
			current, _ := strconv.Atoi(string(pass[j]))

			matching[current] += 1

			if j == 0 {
				continue
			}

			previous, _ := strconv.Atoi(string(pass[j-1]))
			if previous > current {
				increasing = false
				break
			}
		}

		if !increasing {
			continue
		}

		for _, value := range matching {
			if value >= 2 {
				possibilitiesPart1++
				break
			}
		}

		for _, value := range matching {
			if value == 2 {
				possibilitiesPart2++
				break
			}
		}
	}

	fmt.Println("Part 1: ", possibilitiesPart1)
	fmt.Println("Part 2: ", possibilitiesPart2)
}
