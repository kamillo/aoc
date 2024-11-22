package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"os"
	"strconv"
	"strings"
)

// ABCDE
// 1002
//
// DE - two-digit opcode,      02 == opcode 2
// C - mode of 1st parameter,  0 == position mode
// B - mode of 2nd parameter,  1 == immediate mode
// A - mode of 3rd parameter,  0 == position mode,
func runProgram(inCodes []int, systemId int) []int {
	codes := append([]int(nil), inCodes...)
	ptr := 0

	for {
		param1, param2 := 0, 0

		if ptr+1 < len(codes) {
			param1 = codes[ptr+1]
			if ((codes[ptr] / 100) % 10) == 0 {
				if codes[ptr+1] < len(codes) {
					param1 = codes[codes[ptr+1]]
				}
			}
		}

		if ptr+2 < len(codes) {
			param2 = codes[ptr+2]
			if ((codes[ptr] / 1000) % 10) == 0 {
				if codes[ptr+2] < len(codes) {
					param2 = codes[codes[ptr+2]]
				}
			}
		}

		//if ptr + 3 < len(codes) {
		//	param3 = codes[ptr+3]
		//	if ((codes[ptr] / 10000) % 10) == 0 {
		//		if codes[ptr+3] < len(codes) {
		//			param3 = codes[codes[ptr+3]]
		//		}
		//	}
		//}

		switch codes[ptr] % 100 {
		case 1:
			codes[codes[ptr+3]] = param1 + param2
			ptr += 4
		case 2:
			codes[codes[ptr+3]] = param1 * param2
			ptr += 4
		case 3:
			codes[codes[ptr+1]] = systemId
			ptr += 2
		case 4:
			fmt.Println(param1)
			ptr += 2
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
			codes[codes[ptr+3]] = 0
			if param1 < param2 {
				codes[codes[ptr+3]] = 1
			}
			ptr += 4
		case 8: // equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			codes[codes[ptr+3]] = 0
			if param1 == param2 {
				codes[codes[ptr+3]] = 1
			}
			ptr += 4
		case 99: // halt
			return codes
		}
	}
}

func main() {
	lines := utils.GetLines(os.Args[1])

	for _, line := range lines {
		splitted := strings.Split(line, ",")
		ints := make([]int, len(splitted))
		for i := 0; i < len(splitted); i++ {
			value, _ := strconv.Atoi(splitted[i])
			ints[i] = value
		}

		fmt.Println("------ Part 1 -------")
		runProgram(ints, 1)
		fmt.Println("------ Part 2 -------")
		runProgram(ints, 5)
	}
}
