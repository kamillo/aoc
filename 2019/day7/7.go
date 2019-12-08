package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/fileutil"
)

// ABCDE
// 1002
//
// DE - two-digit opcode,      02 == opcode 2
// C - mode of 1st parameter,  0 == position mode
// B - mode of 2nd parameter,  1 == immediate mode
// A - mode of 3rd parameter,  0 == position mode,
func runProgram(inCodes []int, input []int, ptrOptional ...int) (int, []int, int) {
	codes := append([]int(nil), inCodes...)
	ptr := 0
	if len(ptrOptional) > 0 {
		ptr = ptrOptional[0]
	}
	inputPtr := 0
	ret := 0

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
			codes[codes[ptr+1]] = input[inputPtr]
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
			return ret, codes, -1
		}
	}
}

func heapPermutation(a []int) [][]int {
	var permutations [][]int
	var generate func([]int, int)

	generate = func(a []int, size int) {
		if size == 1 {
			A := make([]int, len(a))
			copy(A, a)
			permutations = append(permutations, A)
		}
		for i := 0; i < size; i++ {
			generate(a, size-1)
			if size%2 == 1 {
				a[0], a[size-1] = a[size-1], a[0]
			} else {
				a[i], a[size-1] = a[size-1], a[i]
			}
		}
	}
	generate(a, len(a))
	return permutations
}

func main() {
	lines := fileutil.GetLines(os.Args[1])
	//lines = []string{"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"}
	//lines = []string{"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"}
	//lines = []string{"3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"}

	splitted := strings.Split(lines[0], ",")
	ints := make([]int, len(splitted))
	for i := 0; i < len(splitted); i++ {
		value, _ := strconv.Atoi(splitted[i])
		ints[i] = value
	}

	// Part 1
	{
		a := []int{0, 1, 2, 3, 4}
		permutations := heapPermutation(a)
		max := 0
		for i := range permutations {
			A, _, _ := runProgram(ints, []int{permutations[i][0], 0})
			B, _, _ := runProgram(ints, []int{permutations[i][1], A})
			C, _, _ := runProgram(ints, []int{permutations[i][2], B})
			D, _, _ := runProgram(ints, []int{permutations[i][3], C})
			E, _, _ := runProgram(ints, []int{permutations[i][4], D})

			if E > max {
				max = E
			}
		}
		fmt.Println("Part 1: ", max)
	}

	// Part 2
	{
		a := []int{5, 6, 7, 8, 9}
		permutations := heapPermutation(a)
		max := 0
		for i := range permutations {
			maxE := 0
			A, intsA, ptrA := runProgram(ints, []int{permutations[i][0], 0})
			B, intsB, ptrB := runProgram(ints, []int{permutations[i][1], A})
			C, intsC, ptrC := runProgram(ints, []int{permutations[i][2], B})
			D, intsD, ptrD := runProgram(ints, []int{permutations[i][3], C})
			E, intsE, ptrE := runProgram(ints, []int{permutations[i][4], D})

			for {
				A, intsA, ptrA = runProgram(intsA, []int{E}, ptrA)
				B, intsB, ptrB = runProgram(intsB, []int{A}, ptrB)
				C, intsC, ptrC = runProgram(intsC, []int{B}, ptrC)
				D, intsD, ptrD = runProgram(intsD, []int{C}, ptrD)
				E, intsE, ptrE = runProgram(intsE, []int{D}, ptrE)
				if E > maxE {
					maxE = E
				}
				if ptrE == -1 {
					break
				}
			}

			if maxE > max {
				max = maxE
			}
		}
		fmt.Println("Part 2: ", max)
	}
}
