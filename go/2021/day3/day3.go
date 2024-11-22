package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	input := utils.GetLines("input.txt")
	l := len(input[0])
	counts := map[int]int{}
	cnt := 0
	for _, line := range input {
		cnt++
		for i, c := range line {
			if c == '1' {
				counts[i]++
			}
		}
	}

	gamma := make([]string, l)
	epsilon := make([]string, l)
	for k, v := range counts {
		if v > (cnt - v) {
			gamma[k] = "1"
			epsilon[k] = "0"
		} else {
			gamma[k] = "0"
			epsilon[k] = "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 0)
	epsilonInt, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 0)

	fmt.Println("Part 1: ", gammaInt*epsilonInt)

	rating := func(comp func(int, int) bool) int64 {
		left := input
		res := ""
		for i := 0; i < l; i++ {
			newLeft := []string{}
			bitCount := countBits(left, i)
			if comp(bitCount, len(left)-bitCount) {
				res += "1"
			} else {
				res += "0"
			}

			for _, line := range left {
				if strings.HasPrefix(line, res) {
					newLeft = append(newLeft, line)
				}
			}

			if len(newLeft) == 1 {
				ret, _ := strconv.ParseInt(newLeft[0], 2, 0)
				return ret
			}

			left = newLeft
		}

		return 0
	}

	gammaInt = rating(func(i1, i2 int) bool { return i1 >= i2 })
	epsilonInt = rating(func(i1, i2 int) bool { return i1 < i2 })
	fmt.Println("Part 2: ", gammaInt*epsilonInt)
}

func countBits(input []string, index int) (cnt int) {
	for _, line := range input {
		if line[index] == '1' {
			cnt++
		}
	}
	return
}
