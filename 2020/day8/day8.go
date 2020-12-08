package main

import (
	"fmt"

	"github.com/kamillo/aoc/fileutil"
)

type code struct {
	instruction string
	arg         int
}

// BootCode - AoC 2020 BootCode
type BootCode []code

func main() {
	lines := fileutil.GetLines("input.txt")
	bootcode := make([]code, len(lines))

	for i, line := range lines {
		instruction := ""
		arg := 0

		fmt.Sscanf(line, "%s %d", &instruction, &arg)
		bootcode[i] = code{instruction: instruction, arg: arg}
	}

	// Part1
	fmt.Println(runBootCode(bootcode))

	for j, code := range bootcode {
		if code.instruction == "nop" {
			bootcode[j].instruction = "jmp"
		} else if code.instruction == "jmp" {
			bootcode[j].instruction = "nop"
		}

		if success, accumulator := runBootCode(bootcode); success {
			fmt.Println("Part 2:", accumulator)
			break
		}
		bootcode[j].instruction = code.instruction
	}
}

func runBootCode(bootcode BootCode) (success bool, accumulator int) {
	instructionsDone := make(map[int]bool)

	for index := 0; index < len(bootcode); {
		instr := bootcode[index].instruction
		arg := bootcode[index].arg

		if instructionsDone[index] {
			return false, accumulator
		}
		instructionsDone[index] = true

		switch instr {
		case "nop":
			index++
			break
		case "acc":
			accumulator += arg
			index++
			break
		case "jmp":
			index += arg
			break
		}
	}

	return true, accumulator
}
