package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

var cache map[string]int

func main() {
	lines := utils.GetLines("input.txt")

	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		groupsStr := parts[1]
		groups := utils.ToIntArr(groupsStr, ",")

		cache = make(map[string]int)
		score := findReplacements(strings.Split(parts[0], ""), groups, 0, 0, 0)
		sum += score
	}
	fmt.Println("Part 1:", sum)

	sum = 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		groupsStr := parts[1]
		pipes := strings.Join([]string{parts[0], parts[0], parts[0], parts[0], parts[0]}, "?")
		groupsStr = strings.Join([]string{groupsStr, groupsStr, groupsStr, groupsStr, groupsStr}, ",")
		groups := utils.ToIntArr(groupsStr, ",")

		cache = make(map[string]int)
		score := findReplacements(strings.Split(pipes, ""), groups, 0, 0, 0)
		sum += score
	}
	fmt.Println("Part 2:", sum)
}

func findReplacements(dots []string, blocks []int, i int, bi int, current int) int {
	key := fmt.Sprintf("%d-%d-%d", i, bi, current)
	if val, ok := cache[key]; ok {
		return val
	}

	if i == len(dots) {
		if bi == len(blocks) && current == 0 {
			return 1

		} else if bi == len(blocks)-1 && blocks[bi] == current {
			return 1

		} else {
			return 0
		}
	}

	res := 0
	for _, c := range []string{".", "#"} {
		if dots[i] == c || dots[i] == "?" {
			if c == "." && current == 0 {
				res += findReplacements(dots, blocks, i+1, bi, 0)

			} else if c == "." && current > 0 && bi < len(blocks) && blocks[bi] == current {
				res += findReplacements(dots, blocks, i+1, bi+1, 0)

			} else if c == "#" {
				res += findReplacements(dots, blocks, i+1, bi, current+1)
			}
		}
	}

	cache[key] = res

	return res
}
