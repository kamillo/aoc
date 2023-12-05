package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strconv"
	"strings"
)

// ABCDE
// 1002
//
// DE - two-digit opcode,      02 == opcode 2
// C - mode of 1st parameter,  0 == position mode  , 2 == relative mode
// B - mode of 2nd parameter,  1 == immediate mode
// A - mode of 3rd parameter,  0 == position mode,
func runProgram(inCodes []int, input []int, ptrOptional ...int) (int, []int, int, int) {
	codes := append([]int(nil), inCodes...)
	ptr := 0
	relativeBase := 0

	if len(ptrOptional) > 0 {
		ptr = ptrOptional[0]
		relativeBase = ptrOptional[1]
	}
	inputPtr := 0
	ret := 0

	for {
		pparam1, param2, param3 := 0, 0, 0

		if ptr+1 < len(codes) {
			pparam1 = ptr + 1
			if ((codes[ptr] / 100) % 10) == 0 {
				if codes[ptr+1] < len(codes) {
					pparam1 = codes[ptr+1]
				}
			} else if ((codes[ptr] / 100) % 10) == 2 {
				if relativeBase+codes[ptr+1] < len(codes) {
					pparam1 = relativeBase + codes[ptr+1]
				}
			}
		}

		if ptr+2 < len(codes) {
			param2 = codes[ptr+2]
			if ((codes[ptr] / 1000) % 10) == 0 {
				if codes[ptr+2] < len(codes) {
					param2 = codes[codes[ptr+2]]
				}
			} else if ((codes[ptr] / 1000) % 10) == 2 {
				if relativeBase+codes[ptr+2] < len(codes) {
					param2 = codes[relativeBase+codes[ptr+2]]
				}
			}
		}

		if ptr+3 < len(codes) {
			param3 = ptr + 3
			if ((codes[ptr] / 10000) % 10) == 0 {
				if codes[ptr+3] < len(codes) {
					param3 = codes[ptr+3]
				}
			} else if ((codes[ptr] / 10000) % 10) == 2 {
				param3 = relativeBase + codes[ptr+3]
			}
		}
		param1 := codes[pparam1]
		switch codes[ptr] % 100 {
		case 1:
			codes[param3] = param1 + param2
			ptr += 4
		case 2:
			codes[param3] = param1 * param2
			ptr += 4
		case 3:
			codes[pparam1] = input[inputPtr]
			inputPtr++
			ptr += 2
		case 4:
			//fmt.Println("out: ", param1)
			ret = param1
			ptr += 2
			return param1, codes, ptr, relativeBase
		case 5: // jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			ptr += 3
			if param1 != 0 {
				ptr = param2
			}
		case 6: // jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			ptr += 3
			if param1 == 0 {
				ptr = param2
			}
		case 7: // less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			codes[param3] = 0
			if param1 < param2 {
				codes[param3] = 1
			}
			ptr += 4
		case 8: // equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			codes[param3] = 0
			if param1 == param2 {
				codes[param3] = 1
			}
			ptr += 4
		case 9:
			relativeBase += param1
			ptr += 2
		case 99: // halt
			return ret, codes, -1, -1
		}
	}
}

func main() {
	lines := utils.GetLines("input.txt")
	splitted := strings.Split(lines[0], ",")
	ints := make([]int, len(splitted)*100)
	for i := 0; i < len(splitted); i++ {
		value, _ := strconv.Atoi(splitted[i])
		ints[i] = value
	}

	out := 0
	state := ints
	ptr := 0
	x, y := 0, 0
	plate := make(map[[2]int]int)
	xv, yv := 0, 1
	plate[[2]int{x, y}] = 1
	dir := 0
	base := 0
	maxX := 0
	maxY := 0
	minY := 0
	minX := 0
	for {
		//fmt.Println("start ", plate)
		out, state, ptr, base = runProgram(state, []int{plate[[2]int{x, y}]}, ptr, base)
		if ptr == -1 {
			break
		}
		dir, state, ptr, base = runProgram(state, []int{plate[[2]int{x, y}]}, ptr, base)
		if ptr == -1 {
			break
		}
		plate[[2]int{x, y}] = out
		switch dir {
		case 0:
			xv, yv = -yv, xv
		case 1:
			xv, yv = yv, -xv
		}

		x += xv
		y += yv

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
	}

	fmt.Println(len(plate))
	fmt.Println(minX, minY, maxX, maxY)

	for y := maxY; y >= minY; y-- {
		for x := minX; x < maxX; x++ {
			if c, ok := plate[[2]int{x, y}]; ok {
				color := "#"
				if c == 0 {
					color = " "
				}
				fmt.Printf(color)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
