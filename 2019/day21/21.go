package main

import (
	"fmt"

	"github.com/kamillo/aoc/2019/intcode"
	"github.com/kamillo/aoc/fileutil"
)

func main() {
	lines := fileutil.GetLines("input.txt")
	ints := intcode.ParseInput(lines[0])
	intCode := intcode.Make(ints)

	var out string = ""

	// Jump over 3 tile hole
	prog :=
		`NOT C T
		NOT A J
		OR T J
		AND D J`

	intCode.PutLine(prog)
	intCode.PutLine("WALK")
	for state := intcode.Output; state == intcode.Output; out, state = intCode.GetLine() {
		fmt.Println(out)
	}
	fmt.Println("Part 1:", out)

	prog =
		`NOT C T
		NOT A J
		OR T J
		NOT B T
		OR T J
		NOT J T
		OR E T
		OR H T
		AND T J
		AND D J`
	intCode = intcode.Make(ints)
	intCode.PutLine(prog)
	intCode.PutLine("RUN")
	for state := intcode.Output; state == intcode.Output; out, state = intCode.GetLine() {
		fmt.Println(out)
	}
	fmt.Println("Part 2:", out)
}
