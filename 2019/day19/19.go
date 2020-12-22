package main

import (
	"fmt"
	"github.com/kamillo/aoc/2019/intcode"
	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	ints := intcode.ParseInput(lines[0])
	intCode := intcode.Make(ints)

	sum := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			intCode = intcode.Make(ints)
			intCode.Put([]int{x, y})
			res, _ := intCode.Get()
			fmt.Printf("%d", res)
			sum += res
		}
		println()
	}
	fmt.Println("Part 1: ", sum)

	var check = func(x, y int) int {
		if x < 0 || y < 0 {
			return 0
		}

		intCode = intcode.Make(ints)
		intCode.Put([]int{x, y})
		res, _ := intCode.Get()

		return res
	}

	lastX := 0
	for y := 10; y < 10000; y++ {
		for x := lastX; x < 10000; x++ {
			if check(x, y) == 1 {
				if check(x, y-99) == 1 && check(x+99, y-99) == 1 {
					fmt.Println("Part 2: ", x*10000+(y-99), x, y-99)
					return
				}
				lastX = x - 1
				break
			}
		}
	}
}
