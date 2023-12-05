package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strconv"
	"strings"
)

const (
	multiplicationPrecedence = "*"
	additionPrecedence       = "+"
)

func calcArr(toCalc []string) int {
	calc := func(l int, r int, op string) int {
		if op == "*" {
			return l * r
		} else {
			return l + r
		}
	}

	ll, _ := strconv.Atoi(toCalc[0])
	rr, _ := strconv.Atoi(toCalc[2])
	cc := calc(ll, rr, toCalc[1])
	for c := 4; c < len(toCalc); c++ {
		rr, _ = strconv.Atoi(toCalc[c])
		cc = calc(cc, rr, toCalc[c-1])
	}

	return cc
}

func calcArr2(input []string) int {
	toCalc := make([]string, len(input))
	copy(toCalc, input)
	for {
		found := false
		for i, t := range toCalc {
			sum := 0
			if t == "+" {
				ll, _ := strconv.Atoi(toCalc[i-1])
				rr, _ := strconv.Atoi(toCalc[i+1])
				sum += ll + rr
				toCalc[i-1] = strconv.Itoa(sum)
				toCalc = append(toCalc[:i], toCalc[i+2:]...)
				found = true
				break
			}
		}
		if !found {
			break
		}
	}

	if len(toCalc) > 2 {
		return calcArr(toCalc)
	} else {
		res, _ := strconv.Atoi(toCalc[0])
		return res
	}
}

func eval(equation string, precedence string) int {
	equation = strings.ReplaceAll(equation, "(", "( ")
	equation = strings.ReplaceAll(equation, ")", " )")

	for {
		open := false
		toCalc := make([]string, 0)
		split := strings.Split(equation, " ")
		for _, r := range split {
			if r == "(" {
				open = true
				toCalc = make([]string, 0)
			}
			if _, err := strconv.Atoi(r); err == nil && open {
				toCalc = append(toCalc, r)
			}
			if r == "*" || r == "+" {
				toCalc = append(toCalc, r)
			}
			if r == ")" && open && len(toCalc) > 2 {
				var cc int
				if precedence == multiplicationPrecedence {
					cc = calcArr2(toCalc)
				} else {
					cc = calcArr(toCalc)
				}
				equation = strings.ReplaceAll(equation, "( "+strings.Join(toCalc, " ")+" )", strconv.Itoa(cc))
				open = false
				toCalc = make([]string, 0)
			}
		}
		if !strings.Contains(equation, "(") {
			break
		}
	}
	split := strings.Split(equation, " ")
	if precedence == multiplicationPrecedence {
		return calcArr2(split)
	} else {
		return calcArr(split)
	}
}

func main() {
	lines := utils.GetLines("input.txt")
	//lines = fileutil.GetLines("test.txt")

	part1 := 0
	for _, line := range lines {
		result := eval(line, additionPrecedence)
		part1 += result
	}
	fmt.Println("Part 1:", part1)

	part2 := 0
	for _, line := range lines {
		result := eval(line, multiplicationPrecedence)
		part2 += result
	}
	fmt.Println("Part 2:", part2)
}
