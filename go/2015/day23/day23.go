package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	registers := map[string]uint{}
	lines := utils.GetLines("input.txt")
	fmt.Println("Part 1:", run(lines, registers))

	registers = map[string]uint{}
	registers["a"] = 1
	fmt.Println("Part 2:", run(lines, registers))
}

func run(program []string, registers map[string]uint) uint {
	i := 0
	for {
		splited := strings.Split(program[i], " ")
		r := splited[1]
		step := 1

		switch splited[0] {
		case "hlf":
			registers[r] = registers[r] / 2
		case "tpl":
			registers[r] *= 3
		case "inc":
			registers[r]++
		case "jmp":
			if offset, err := strconv.Atoi(r); err == nil {
				step = offset
			}
		case "jie":
			r = strings.Trim(splited[1], ",")
			if registers[r]%2 == 0 {
				if offset, err := strconv.Atoi(splited[2]); err == nil {
					step = offset
				}
			}
		case "jio":
			r = strings.Trim(splited[1], ",")
			if registers[r] == 1 {
				if offset, err := strconv.Atoi(splited[2]); err == nil {
					step = offset
				}
			}
		}

		i += step

		if i < 0 || i >= len(program) {
			return registers["b"]
		}
	}
}
