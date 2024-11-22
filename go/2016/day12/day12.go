package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	assembunny(lines, 0)
	assembunny(lines, 1)
}

func assembunny(lines []string, init int) {
	registers := map[string]int{"a": 0, "b": 0, "c": init, "d": 0}

	for i := 0; i < len(lines); {
		instruction, param1, param2 := "", "", ""
		fields := strings.Fields(lines[i])

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
				}
			} else {
				i++
			}
		}
	}

	fmt.Println(registers)
}
