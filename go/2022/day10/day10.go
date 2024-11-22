package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	fmt.Println("Part 1:", part1(lines))
	for _, c := range part2(lines) {
		fmt.Println(c)
	}
}

func part1(intructions []string) (signal int) {
	register := 1
	cycle := 0
	cycles := map[int]bool{20: true, 60: true, 100: true, 140: true, 180: true, 220: true}

	for _, instruction := range intructions {
		fields := strings.Fields(instruction)
		tick := func() {
			cycle++
			if cycles[cycle] {
				signal += cycle * register
			}
		}

		switch fields[0] {
		case "addx":
			tick()
			tick()
			if x, error := strconv.Atoi(fields[1]); error == nil {
				register += x
			} else {
				panic("atoi")
			}

		case "noop":
			tick()
		}
	}

	return signal
}

func part2(intructions []string) (crt [6][40]string) {
	register := 1
	cycle := 0
	row := 0

	for _, instruction := range intructions {
		tick := func() {
			col := cycle % 40

			if cycle != 0 && col == 0 {
				row++
			}

			if register-1 == col || register == col || register+1 == col {
				crt[row][col] = "#"
			} else {
				crt[row][col] = "."
			}

			cycle++
		}

		fields := strings.Fields(instruction)
		switch fields[0] {
		case "addx":
			tick()
			tick()
			if x, error := strconv.Atoi(fields[1]); error == nil {
				register += x
			} else {
				panic("atoi")
			}

		case "noop":
			tick()
		}
	}

	return crt
}
