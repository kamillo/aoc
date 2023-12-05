package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

type Scanner struct {
	Depth int
	Range int
}

func main() {
	scanners := []Scanner{}

	for _, line := range utils.GetLines("input.txt") {
		depth, rang := 0, 0
		fmt.Sscanf(line, "%d: %d", &depth, &rang)
		scanners = append(scanners, Scanner{depth, rang})
	}

	cost := 0
	for _, scanner := range scanners {
		if (scanner.Depth)%((scanner.Range-1)*2) == 0 {
			cost += scanner.Depth * scanner.Range
		}
	}
	fmt.Println("Part 1:", cost)

	caught := true
	delay := 0
	for ;;delay++ {
		caught = false
		for _, scanner := range scanners {
			if (scanner.Depth+delay)%((scanner.Range-1)*2) == 0 {
				caught = true
				break
			}
		}
		if !caught {
			break
		}
	}
	fmt.Println("Part 2:", delay)
}
