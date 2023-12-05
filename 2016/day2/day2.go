package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	{
		x, y := 1, 1
		keypad := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		for _, line := range lines {
			//found := false
			for _, char := range line {
				switch char {
				case 'U':
					if y-1 < 0 {
					} else {
						y--
					}
					break
				case 'D':
					if y+1 >= len(keypad[0]) {
					} else {
						y++
					}
					break
				case 'L':
					if x-1 < 0 {
					} else {
						x--
					}
					break
				case 'R':
					if x+1 >= len(keypad[0]) {
					} else {
						x++
					}
					break
				}
			}
			fmt.Println(keypad[y][x])
		}
	}

	x, y := 1, 1
	keypad := [][]int{
		{0, 0, 1, 0, 0},
		{0, 2, 3, 4, 0},
		{5, 6, 7, 8, 9},
		{0, 'A', 'B', 'C', 0},
		{0, 0, 'D', 0, 0},
	}

	for _, line := range lines {
		//found := false
		for _, char := range line {
			switch char {
			case 'U':
				if y-1 < 0 || keypad[y-1][x] == 0 {
				} else {
					y--
				}
				break
			case 'D':
				if y+1 >= len(keypad[0]) || keypad[y+1][x] == 0 {
				} else {
					y++
				}
				break
			case 'L':
				if x-1 < 0 || keypad[y][x-1] == 0 {
				} else {
					x--
				}
				break
			case 'R':
				if x+1 >= len(keypad[0]) || keypad[y][x+1] == 0 {
				} else {
					x++
				}
				break
			}
		}
		fmt.Printf("%v\n", keypad[y][x])
	}
}
