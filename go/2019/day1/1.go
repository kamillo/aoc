package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"os"
	"strconv"
)

func calcFuel(mass int) int {
	if mass <= 0 {
		return 0
	}

	return mass + calcFuel(mass/3-2)
}

func main() {
	lines := utils.GetLines(os.Args[1])

	sum1, sum2 := 0, 0
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		moduleFuel := mass/3 - 2
		sum1 += moduleFuel
		sum2 += calcFuel(moduleFuel)
	}

	fmt.Println("Part 1: ", sum1)
	fmt.Println("Part 2: ", sum2)
}
