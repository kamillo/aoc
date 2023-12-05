package main

import (
	"strconv"
	"strings"

	"github.com/kamillo/aoc/2019/intcode"
	"github.com/kamillo/aoc/fileutil"
)

func main() {
	lines := fileutil.GetLines("input.txt")
	splitted := strings.Split(lines[0], ",")
	ints := make([]int, len(splitted)*1000)
	for i := 0; i < len(splitted); i++ {
		ints[i], _ = strconv.Atoi(splitted[i])
	}

	{
		intCode := intcode.Make(ints)
		scene := make(map[[2]int]int)
		status := intcode.Output
		var x, y, id int
		for status != intcode.Exit {
			x, status = intCode.Get()
			y, status = intCode.Get()
			id, status = intCode.Get()
			scene[[2]int{x, y}] = id
		}

		sum := 0
		for _, v := range scene {
			if v == 2 {
				sum++
			}
		}
		println("Part 1: ", sum)
	}
	// Part 2
	{
		scene2 := [25][41]int{}
		scene := make(map[int][2]int)
		ints[0] = 2

		// lets cheat ;)
		for i := 1541; i < 1583; i++ {
			ints[i] = 1
		}

		intCode := intcode.Make(ints)
		status := intcode.Output
		var x, y, id, score int
		for status != intcode.Exit {
			intCode.Put([]int{0})
			x, status = intCode.Get()
			y, _ = intCode.Get()
			id, _ = intCode.Get()

			if x != -1 {
				scene2[y][x] = id
				scene[id] = [2]int{x, y}
			} else {
				score = id
				//for i := range scene2 {
				//	for j := range scene2[i] {
				//		fmt.Printf("%d", scene2[i][j])
				//	}
				//	println()
				//}
			}
		}
		println("Part 2: ", score)
	}
}
