package main

import (
	"fmt"
	"math"
)

// Register A: 64854237
// Register B: 0
// Register C: 0
// Program: 2,4,1,1,7,5,1,5,4,0,5,5,0,3,3,0

// Register A: 729
// Register B: 0
// Register C: 0
// Program: 0,1,5,4,3,0

func main() {

	registers := map[string]int{
		"A": 64854237,
		// "A": 729,
		"B": 0,
		"C": 0,
	}
	// program := []int{0, 1, 5, 4, 3, 0}
	program := []int{2, 4, 1, 1, 7, 5, 1, 5, 4, 0, 5, 5, 0, 3, 3, 0}

	var run = func(registers map[string]int, program []int) []int {
		ptr := 0
		out := []int{}

		for ptr < len(program) {
			switch program[ptr] {
			case 0:
				op := getOperand(registers, program[ptr+1])
				op = int(math.Pow(2, float64(op)))
				registers["A"] = int(registers["A"] / op)
				ptr += 2
			case 1:
				registers["B"] = registers["B"] ^ program[ptr+1]
				ptr += 2
			case 2:
				registers["B"] = getOperand(registers, program[ptr+1]) & 7
				ptr += 2

			case 3:
				if registers["A"] == 0 {
					ptr += 2
				} else {
					ptr = program[ptr+1]
				}
			case 4:
				registers["B"] = registers["B"] ^ registers["C"]
				ptr += 2
			case 5:
				out = append(out, getOperand(registers, program[ptr+1])&7)
				ptr += 2

			case 6:
				op := getOperand(registers, program[ptr+1])
				op = int(math.Pow(2, float64(op)))
				registers["B"] = int(registers["A"] / op)
				ptr += 2

			case 7:
				op := getOperand(registers, program[ptr+1])
				op = int(math.Pow(2, float64(op)))
				registers["C"] = int(registers["A"] / op)
				ptr += 2
			}
		}
		return out
	}

	fmt.Println("Part 1:", run(registers, program))

	type state []int

	queue := []state{}
	for i := 0; i < 8; i++ {
		queue = append(queue, state{i})
	}

	var part2 int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		var x int
		for i := len(cur) - 1; i >= 0; i-- {
			s := cur[i] << (3 * i)
			x = x | s
		}

		registers["A"] = x
		vals := run(registers, program)
		vp := 0
		matched := true
		for p := len(program) - len(vals); p < len(program); p++ {
			if vals[vp] != program[p] {
				matched = false
				break
			}
			vp++
		}

		done := matched && len(program) == len(vals)
		if done {
			part2 = x
			break
		}

		if matched {
			for i := 0; i < 8; i++ {
				nseg := make([]int, len(cur))
				copy(nseg, cur)
				nseg = append([]int{i}, nseg...)
				queue = append(queue, nseg)
			}
		}
	}

	fmt.Println("Part 2:", part2)
}

func getOperand(registers map[string]int, operand int) int {
	if operand == 4 {
		return registers["A"]
	}

	if operand == 5 {
		return registers["B"]
	}

	if operand == 6 {
		return registers["C"]
	}

	if operand >= 7 {
		panic("Invalid operand")
	}
	return operand
}
