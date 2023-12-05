package bootcode

// Code
type Code struct {
	Instruction string
	Arg         int
}

// BootCode - AoC 2020 BootCode
type BootCode []Code

// RunBootCode
func RunBootCode(bootcode BootCode) (success bool, accumulator int) {
	instructionsDone := make(map[int]bool)

	for index := 0; index < len(bootcode); {
		instr := bootcode[index].Instruction
		arg := bootcode[index].Arg

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
