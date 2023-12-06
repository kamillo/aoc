package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	times := utils.ToIntArr(lines[0], " ")
	distances := utils.ToIntArr(lines[1], " ")

	fmt.Println("Part 1:", calc(times, distances))

	time := utils.JustAtoi(utils.IntArrToStr(times))
	distance := utils.JustAtoi(utils.IntArrToStr(distances))

	fmt.Println("Part 2:", calc([]int{time}, []int{distance}))
}

func calc(times []int, distances []int) int {
	sum := 1
	for i := 0; i < len(times); i++ {
		c := 0
		for j := 0; j < times[i]; j++ {
			if j*(times[i]-j) > distances[i] {
				c++
			}
		}
		sum *= c
	}

	return sum
}
