package main

import (
	"bufio"
	"fmt"
	"github.com/kamillo/aoc/2019/intcode"
	"github.com/kamillo/aoc/utils"
	"os"
)

func main() {
	lines := utils.GetLines("input.txt")
	ints := intcode.ParseInput(lines[0])

	intCode := intcode.Make(ints)

	for {
		out, state := intCode.GetLine()
		fmt.Println(out)
		if state != intcode.Output {
			break
		}
		if out == "Command?" {
			fmt.Print("Command: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // use `for scanner.Scan()` to keep reading
			command := scanner.Text()
			intCode.PutLine(command)
		}
	}
}
