package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/fileutil"
)

// ABCDE
// 1002
//
// DE - two-digit opcode,      02 == opcode 2
// C - mode of 1st parameter,  0 == position mode  , 2 == relative mode
// B - mode of 2nd parameter,  1 == immediate mode
// A - mode of 3rd parameter,  0 == position mode,
func runProgram(inCodes []int, input []int, ptrOptional ...int) (int, []int, int) {
	codes := append([]int(nil), inCodes...)
	ptr := 0
	relativeBase := 0

	if len(ptrOptional) > 0 {
		ptr = ptrOptional[0]
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
			return param1, codes, ptr
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
			return ret, codes, -1
		}
	}
}

func main() {
	lines := fileutil.GetLines("input.txt")
	//lines = []string{"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"}
	//lines = []string{"1102,34915192,34915192,7,4,7,99,0"}
	//lines = []string{"104,1125899906842624,99"}

	splitted := strings.Split(lines[0], ",")
	ints := make([]int, len(splitted)*100)
	for i := 0; i < len(splitted); i++ {
		value, _ := strconv.Atoi(splitted[i])
		ints[i] = value
	}

	part1, _, _ := runProgram(ints, []int{1})
	part2, _, _ := runProgram(ints, []int{2})
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 1: ", part2)
}
