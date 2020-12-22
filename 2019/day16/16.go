package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := utils.GetLines("input.txt")
	input := lines[0]
	//input = "03036732577212944063491565474664"
	//offset = 303673
	offset := 0

	splited := []int{}
	for _, s := range strings.Split(input, "") {
		i, _ := strconv.Atoi(s)
		splited = append(splited, i)
	}

	splited2 := []int{}
	for i := 0; i < 10000; i++ {
		splited2 = append(splited2, splited...)
	}

	pattern := [...]int{0, 1, 0, -1}
	result := make([]int, len(splited))
	for i := 0; i < 100; i++ {
		for j := 0; j < len(splited); j++ {
			sum := 0
			p := j
			patternIdx := 0

			for d := 0; d < len(splited); d++ {
				p--
				if p < 0 {
					patternIdx = (patternIdx + 1) % 4
					p = j
				}
				sum += splited[d] * pattern[patternIdx]
			}
			result[j] = int(math.Abs(float64(sum % 10)))
		}
		splited = result
	}
	fmt.Println("Part 1: ", result[:8])

	offset = 5970951
	splited = splited2
	result = make([]int, len(splited))
	for i := 0; i < 100; i++ {
		sum := 0
		for d := offset; d < len(splited); d++ {
			sum += splited[d]
		}

		result[offset] = int(math.Abs(float64(sum % 10)))

		for j := offset + 1; j < len(splited); j++ {
			sum -= splited[j-1]
			result[j] = int(math.Abs(float64(sum % 10)))
		}
		copy(splited, result)
	}
	fmt.Println("Part 2: ", result[offset:offset+8])
}
