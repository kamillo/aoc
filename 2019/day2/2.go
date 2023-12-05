package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"os"
	"strconv"
	"strings"
)

func runProgram(codes []int, noun int, verb int) []int {
	for i := 0;; i += 4 {
		switch  codes[i] {
		case 1: codes[codes[i + 3]] = codes[codes[i + 1]] + codes[codes[i + 2]]
		case 2: codes[codes[i + 3]] = codes[codes[i + 1]] * codes[codes[i + 2]]
		case 99:
			return codes
		}
	}
}

func main() {
	lines := fileutil.GetLines(os.Args[1])
	//lines := []string{"1,1,1,4,99,5,6,0,99"}

	for _, line := range lines {
		splitted := strings.Split(line, ",")
		ints := make([]int, len(splitted))
		for i := 0; i < len(splitted); i++ {
			value, _ := strconv.Atoi(splitted[i])
			ints[i] = value
		}

		part1 := runProgram(ints, 12, 2)
		part2 := runProgram(ints, 71, 95)

		fmt.Println(part1)
		fmt.Println(part2)
	}
}
