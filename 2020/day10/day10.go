package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"sort"
	"strconv"
)

func main() {
	lines := fileutil.GetLines("test.txt")

	differences := make(map[int]int)
	joltages := make([]int, len(lines))
	for i, line := range lines {
		joltages[i], _ = strconv.Atoi(line)
	}

	sort.Ints(joltages)

	for j, _ := range joltages {
		prev := 0
		if j > 0 {
			prev = joltages[j-1]
		}
		differences[joltages[j]-prev]++
	}
	differences[3]++ // last one on device

	fmt.Println("Part 1: ", differences[1]*differences[3])

	ways := make(map[int]int)
	ways[0] = 1
	for _, i := range joltages {
		ways[i] = ways[i-1] + ways[i-2] + ways[i-3]
	}
	fmt.Println("Part 2: ", ways)
}
