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

	// Reverse engineered asm, fast enought
	b := 106700 // 67
	c := 123700 //67

	g := 1
	h := 0
	for g != 0 {
		f := 1
		d := 2
		for g != 0 {
			// e := 2

			g = 0
			if b%d == 0 {
				f = 0
			}

			// for g != 0 {
			// 	g = d
			// 	g *= e
			// 	g -= b

			// 	if g == 0 {
			// 		f = 0
			// 	}
			// 	e++
			// 	g = e
			// 	g -= b
			// }

			d++
			g = d
			g -= b
		}

		if f == 0 {
			h++
		}

		g = b
		g -= c
		if g == 0 {
			break
			//return h
		}
		b += 17
	}

	fmt.Println("Part 2:", h)
}

func debug(lines []string) int {
	registers := map[string]int{}
	mulCnt := 0
	registers["a"] = 0

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

		//fmt.Println(registers)
	}

	return mulCnt
}
