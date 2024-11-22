package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	program := [][]string{}
	for _, line := range lines {
		program = append(program, strings.Fields(line))
	}
	assembunny(program, 7)

	// {
	// 	a := 12
	// 	b := 11
	// 	for {
	// 		d := a
	// 		a = 0
	// 		c := b
	// 		for d != 0 {
	// 			c = b

	// 			for c != 0 {
	// 				a++
	// 				c--
	// 			}

	// 			d--
	// 		}
	// 		b--
	// 		c = b
	// 		d = c
	// 		for d != 0 {
	// 			d--
	// 			c++
	// 		}

	// 		fmt.Println(a, b, c, d)

	// 		c = -16
	// 	}
	// c = 77
	// d = 87

	// (12)! + 77*87

	// }

	program = [][]string{}
	for _, line := range lines {
		program = append(program, strings.Fields(line))
	}
	assembunny(program, 12)
}

func assembunny(lines [][]string, init int) {
	registers := map[string]int{"a": init, "b": 0, "c": 0, "d": 0}

	for i := 0; i < len(lines); {
		instruction, param1, param2 := "", "", ""
		fields := lines[i]

		if len(fields) > 1 {
			instruction = fields[0]
			param1 = fields[1]
		}

		if len(fields) > 2 {
			param2 = fields[2]
		}

		switch instruction {
		case "cpy":
			if n, err := strconv.Atoi(param1); err == nil {
				registers[param2] = n
			} else {
				registers[param2] = registers[param1]
			}
			i++
		case "inc":
			registers[param1]++
			i++
		case "dec":
			registers[param1]--
			i++
		case "jnz":
			p1, err := strconv.Atoi(param1)

			if err != nil {
				p1 = registers[param1]
			}

			if p1 != 0 {
				if n, err := strconv.Atoi(param2); err == nil {
					i += n
				} else {
					i += registers[param2]
				}
			} else {
				i++
			}

		case "tgl":
			next := 0
			if n, err := strconv.Atoi(param1); err == nil {
				next = n
			} else {
				next = registers[param1]
			}

			if next+i < len(lines) {
				fields := lines[next+i]
				instruction := fields[0]

				if len(fields) == 2 {
					if instruction == "inc" {
						instruction = "dec"
					} else {
						instruction = "inc"
					}
				}

				if len(fields) == 3 {
					if instruction == "jnz" {
						instruction = "cpy"
					} else {
						instruction = "jnz"
					}
				}

				lines[next+i][0] = instruction
			}
			i++
		}
	}

	fmt.Println(registers)
}
