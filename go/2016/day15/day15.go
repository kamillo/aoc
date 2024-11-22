package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

type Disc struct {
	positions int
	start     int
}

func main() {
	lines := utils.GetLines("input.txt")

	discs := []Disc{}
	for _, line := range lines {
		var i, p, s int
		fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.", &i, &p, &s)

		discs = append(discs, Disc{p, s})
	}

	fmt.Println("Part 1:", wait(discs))

	discs = append(discs, Disc{11, 0})
	fmt.Println("Part 1:", wait(discs))
}

func wait(discs []Disc) int {
	for w := 0; ; w++ {
		success := true
		for i, d := range discs {
			time := w + i + 1
			if (d.start+time)%d.positions != 0 {
				success = false
				break
			}
		}

		if success {
			return w
		}
	}
}
