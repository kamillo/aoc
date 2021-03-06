package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"

	"github.com/kamillo/aoc/2020/bootcode"
)

func main() {
	lines := utils.GetLines("input.txt")
	bc := make(bootcode.BootCode, len(lines))

	for i, line := range lines {
		instruction := ""
		arg := 0

		fmt.Sscanf(line, "%s %d", &instruction, &arg)
		bc[i] = bootcode.Code{Instruction: instruction, Arg: arg}
	}

	// Part1
	fmt.Println(bootcode.RunBootCode(bc))

	for j, code := range bc {
		if code.Instruction == "nop" {
			bc[j].Instruction = "jmp"
		} else if code.Instruction == "jmp" {
			bc[j].Instruction = "nop"
		}

		if success, accumulator := bootcode.RunBootCode(bootcode); success {
			fmt.Println("Part 2:", accumulator)
			break
		}
		bc[j].Instruction = code.Instruction
	}
}
