package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := utils.GetLines("input.txt")

	toBinary := func(r rune) rune {
		switch r {
		case 'B', 'R':
			return '1'
		case 'F', 'L':
			return '0'
		}
		return r
	}

	ids := make([]int, len(lines))
	for i, line := range lines {
		binary := strings.Map(toBinary, line)

		row, _ := strconv.ParseInt(binary[:7], 2, 0)
		col, _ := strconv.ParseInt(binary[7:], 2, 0)

		ids[i] = int(row)*8 + int(col)
	}

	fmt.Println("Part 1: ", utils.MaxInArray(ids))

	sort.Ints(ids)
	for i, value := range ids {
		if i != 0 && i != len(ids)-1 {
			if ids[i-1] != value-1 {
				fmt.Println("Part 2: ", value-1)
				break
			}
		}
	}
}
