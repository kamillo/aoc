package main

import (
	"fmt"
	"strings"
)

func main() {
	startRow := "^^^^......^...^..^....^^^.^^^.^.^^^^^^..^...^^...^^^.^^....^..^^^.^.^^...^.^...^^.^^^.^^^^.^^.^..^.^"

	tiles := []string{}
	tiles = append(tiles, startRow)

	for i := 0; i < 400000-1; i++ {
		row := ""

		for c := range startRow {
			left := false
			center := false
			right := false

			if c > 0 {
				left = startRow[c-1] == '^'
			}
			center = startRow[c] == '^'
			if c+1 < len(startRow) {
				right = startRow[c+1] == '^'
			}

			tile := "."
			if (left && center && !right) ||
				(!left && center && right) ||
				(left && !center && !right) ||
				(!left && !center && right) {
				tile = "^"
			}

			row += tile
		}

		tiles = append(tiles, row)
		startRow = row
	}

	safeCnt := 0
	for _, row := range tiles {
		fmt.Println(row)
		safeCnt += strings.Count(row, ".")
	}

	fmt.Println(safeCnt)
}
