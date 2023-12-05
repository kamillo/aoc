package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	paperMap := map[utils.Point2D]bool{}
	folds := []utils.Point2D{}
	maxX, maxY := 0, 0

	for _, l := range utils.GetLines("input.txt") {
		if len(l) == 0 {
			continue
		}

		if strings.HasPrefix(l, "fold") {
			axis := 0
			value := 0

			if _, err := fmt.Sscanf(l, "fold along %c=%d", &axis, &value); err == nil {
				if axis == 'x' {
					folds = append(folds, utils.NewPoint2D(value, 0))
				} else {
					folds = append(folds, utils.NewPoint2D(0, value))
				}
			}
			continue
		}

		line := utils.ToIntArr(l, ",")
		if line[0] > maxX {
			maxX = line[0]
		}

		if line[1] > maxY {
			maxY = line[1]
		}

		paperMap[utils.NewPoint2D(line[0], line[1])] = true
	}

	for _, fold := range folds {
		newPaper := map[utils.Point2D]bool{}

		for k := range paperMap {
			if fold.X > 0 {
				if k.X > fold.X {
					x := fold.X - (k.X - fold.X)
					newPaper[utils.NewPoint2D(x, k.Y)] = true

				} else {
					newPaper[k] = true
				}
			} else {

				if k.Y > fold.Y {
					y := fold.Y - (k.Y - fold.Y)
					newPaper[utils.NewPoint2D(k.X, y)] = true

				} else {
					newPaper[k] = true
				}
			}
		}
		paperMap = newPaper
		maxX = fold.X
		maxY = fold.Y
	}
	fmt.Println(paperMap)

	code := make([][]byte, maxY)
	for i := 0; i < maxY; i++ {
		code[i] = make([]byte, 50)
	}
	for k := range paperMap {
		code[k.Y][k.X] = '#'
	}
	for y := range code {
		for x := range code[y] {
			if code[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
