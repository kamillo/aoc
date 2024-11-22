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

	// for i := 0; ; i++ {
	// 	clock(i)
	// }

	for i := 0; ; i++ {
		assembunny(program, i)
	}

}

func clock(a int) {
	fmt.Println("clock", a)

	d := a
	c := 15

	b := 170
	d += b * c
	prevB := 1

	for {
		a = d
		for a != 0 {
			b = a % 2
			a /= 2
			if b == prevB {
				return
			}
			prevB = b
		}
	}
}

func assembunny(lines [][]string, init int) {
	fmt.Println("assembuny", init)

	registers := map[string]int{"a": init, "b": 0, "c": 0, "d": 0}
	prevOut := 1

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

		case "out":
			v := 0
			if n, err := strconv.Atoi(param1); err == nil {
				v = n
			} else {
				v = registers[param1]
			}

			if v == prevOut {
				return
			}

			prevOut = v
			i++
		}
	}

	fmt.Println(registers)
}
