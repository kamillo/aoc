package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Pair struct{ Id, Result int }

func main() {
	lines := utils.GetLines("input.txt")
	fmt.Println("Part 1:", debug(lines))
}

func debug(lines []string) int {
	registers := map[string]int{}
	mulCnt := 0
	registers["a"] = 1

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fields := strings.Fields(line)

		b := 0
		if len(fields) > 2 {
			b = registers[fields[2]]
			if x, error := strconv.Atoi(fields[2]); error == nil {
				b = x
			}
		}

		switch fields[0] {
		case "set":
			registers[fields[1]] = b
		case "sub":
			registers[fields[1]] -= b
		case "mul":
			mulCnt++
			registers[fields[1]] *= b
		case "jnz":
			a := 0
			if len(fields) > 2 {
				a = registers[fields[1]]
				if x, error := strconv.Atoi(fields[1]); error == nil {
					a = x
				}
			}
			if a != 0 {
				i += b
				i--
			}
		}

		fmt.Println(registers)
	}

	return mulCnt
}
