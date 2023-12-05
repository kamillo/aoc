package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

/*
[
    {14 8 push}
    {13 8 push}
    {13 3 push}
    {12 10 push}
    {-12 8 pop}
    {12 8 push}
    {-2 8 pop}
    {-11 5 pop}
    {13 9 push}
    {14 3 push}
    {0 4 pop}
    {-12 9 pop}
    {-13 2 pop}
    {-6 7 pop}
]
*/

/*
    push w0 + 8
		push w1 + 8
							push w2 + 3
					push w3 + 10
					pop w4 == z - 12
						push w5 + 8
						pop w6 == z - 2
							pop w7 == z - 11
			push w8 + 9
				push w9 + 3
				pop w10 == z
			pop w11 == z -12
		pop w12 == z -13
	pop w13 == z -6

	w4  == w3 -2
	w6  == w5 + 6
	w7  == w2 - 8
	w10 == w9 + 3
	w11 == w8 - 3
	w12 == w1 - 5
	w13 == w0 + 2

	biggest:
	01234567890123
	79997391969649

	smallest:
	01234567890123
	16931171414113

*/

type instruction struct {
	x, w   int
	action string
}

func main() {

	ptr := -1
	tmp := 0
	instructions := []instruction{}
	tmpInstr := instruction{}

	for _, line := range utils.GetLines("input.txt") {
		fields := strings.Fields(line)
		tmp++
		if fields[0] == "inp" {
			tmp = 0
			ptr++
			tmpInstr = instruction{}
		}
		if tmp == 4 {
			if fields[2] == "1" {
				tmpInstr.action = "push"
			} else {
				tmpInstr.action = "pop"
			}
		}

		if tmp == 5 {
			tmpInstr.x, _ = strconv.Atoi(fields[2])
		}

		if tmp == 15 {
			tmpInstr.w, _ = strconv.Atoi(fields[2])
			instructions = append(instructions, tmpInstr)
		}
	}

	fmt.Println(instructions)
	number := "79997391969649"
	if _, ok := checkNumber(number); ok {
		fmt.Println("Part 1: ", number)
	}

	number = "16931171414113"
	if _, ok := checkNumber(number); ok {
		fmt.Println("Part 2: ", number)
	}
}

func head(s []int) int {
	if len(s) > 0 {
		return s[len(s)-1]
	}

	return 0
}

// func findNumber(number string, instructions []instruction) bool {
// 	zs := []int{0}
// 	for n, i := range instructions {
// 		z := head(zs)

// 		if i.action == "pop" && len(zs) > 0 {
// 			zs = zs[:len(zs)-1]
// 		}

// 		if digit, _ := strconv.Atoi(string(number[n])); digit != z+i.x {
// 			zs = append(zs, digit+i.w)
// 		}
// 	}

// 	if zs[len(zs)-1] == 0 {
// 	}

// 	return true
// }

func checkNumber(number string) (int, bool) {
	ptr := 0
	registers := map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}
	for _, line := range utils.GetLines("input.txt") {
		fields := strings.Fields(line)

		b := 0
		if len(fields) > 2 {
			b = registers[fields[2]]
			if x, error := strconv.Atoi(fields[2]); error == nil {
				b = x
			}
		}

		switch fields[0] {
		case "inp":
			fmt.Println(registers["z"])
			registers[fields[1]], _ = strconv.Atoi(string(number[ptr]))
			ptr++
		case "add":
			registers[fields[1]] += b
		case "mul":
			registers[fields[1]] *= b
		case "div":
			registers[fields[1]] /= b
		case "mod":
			if b == 0 {
				return registers["z"], false
			}
			registers[fields[1]] %= b
		case "eql":
			if registers[fields[1]] == b {
				registers[fields[1]] = 1
			} else {
				registers[fields[1]] = 0
			}
		}
	}

	return registers["z"], true
}
