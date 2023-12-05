package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	max := 0
	elf := 0
	elfs := []int{}

	for _, line := range lines {
		if i, err := strconv.Atoi(line); err == nil {
			elf += i

		} else {
			if elf > max {
				max = elf
			}
			elfs = append(elfs, elf)
			elf = 0
		}
	}

	elfs = append(elfs, elf)
	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))

	fmt.Println("Part 1:", elfs[0])
	fmt.Println("Part 2:", elfs[0]+elfs[1]+elfs[2])
}
