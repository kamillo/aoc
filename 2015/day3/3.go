package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"os"
	"strings"
)

func main() {
	lines := utils.GetLines(os.Args[1])
	// lines := [1]string{"^v^v^v^v^v"}
	{
		houses := make(map[[2]int]int)
		// point := [2]int{0,0}
		x, y := 0, 0
		houses[[2]int{x, y}]++

		for _, line := range lines {
			for _, direction := range strings.Split(line, "") {
				switch direction {
				case "^":
					y++
				case "v":
					y--
				case ">":
					x++
				case "<":
					x--
				}

				houses[[2]int{x, y}]++
			}
		}

		// fmt.Println(houses)
		fmt.Println("Part 1: ", len(houses))
	}

	{
		x, y := 0, 0
		xx, yy := 0, 0

		houses := make(map[[2]int]int)
		houses[[2]int{x, y}]++
		houses[[2]int{xx, yy}]++

		xp := &x
		yp := &y

		toggle := func() {
			if xp == &x && yp == &y {
				xp, yp = &xx, &yy
			} else {
				xp, yp = &x, &y
			}
		}
		for _, line := range lines {
			for _, direction := range strings.Split(line, "") {
				switch direction {
				case "^":
					*yp++
				case "v":
					*yp--
				case ">":
					*xp++
				case "<":
					*xp--
				}

				houses[[2]int{*xp, *yp}]++
				toggle()
			}
		}
		// fmt.Println(houses)
		fmt.Println("Part 2: ", len(houses))
	}
}
