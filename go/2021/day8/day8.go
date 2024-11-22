package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	notes := utils.GetLines("input.txt")

	part1 := 0
	part2 := 0
	for _, line := range notes {
		split := strings.Split(line, " | ")

		pattern := strings.Fields(split[0])
		output := strings.Fields(split[1])

		digits := map[string]int{}
		numbers := map[int]string{}

		sort.Slice(pattern, func(i int, j int) bool {
			return len(pattern[i]) < len(pattern[j])
		})

		c := ""
		f := ""
		i := 0
		bd := ""
		for len(digits) != 10 {

			pat := pattern[i]
			i = (i + 1) % (len(pattern))

			ps := []rune(pat)
			sort.Slice(ps, func(i int, j int) bool {
				return ps[i] < ps[j]
			})
			p := string(ps)

			if len(p) == 2 {
				digits[p] = 1
				numbers[1] = p
			}

			if len(p) == 3 {
				digits[p] = 7
				numbers[7] = p
			}

			if len(p) == 4 {
				digits[p] = 4
				numbers[4] = p

				if c != "" && f != "" {
					bd = strings.Replace(p, c, "", 1)
					bd = strings.Replace(bd, f, "", 1)
				}
			}

			// 5: 2, 3, 5
			if len(p) == 5 {
				if strings.Contains(p, string(numbers[1][0])) && strings.Contains(p, string(numbers[1][1])) {
					digits[p] = 3
					numbers[3] = p
				}

				// 2, 5
				if c != "" {
					if strings.Contains(p, c) && !strings.Contains(p, f) {
						digits[p] = 2
						numbers[2] = p
					} else if strings.Contains(p, f) && !strings.Contains(p, c) {
						digits[p] = 5
						numbers[5] = p
					}
				}
			}

			// 6: 0, 6, 9
			if len(p) == 6 {
				if !strings.Contains(p, string(numbers[1][0])) || !strings.Contains(p, string(numbers[1][1])) {
					digits[p] = 6
					numbers[6] = p

					if !strings.Contains(p, string(numbers[1][0])) {
						c = string(numbers[1][0])
						f = string(numbers[1][1])
					} else {
						c = string(numbers[1][1])
						f = string(numbers[1][0])
					}

					// 0, 9
				} else if bd != "" {
					if strings.Contains(p, string(bd[0])) && strings.Contains(p, string(bd[1])) {
						digits[p] = 9
						numbers[9] = p
					} else {
						digits[p] = 0
						numbers[0] = p
					}
				}
			}

			if len(p) == 7 {
				digits[p] = 8
				numbers[8] = p
			}
		}

		n := 0
		for x, i := range output {
			if len(i) == 2 || len(i) == 3 || len(i) == 4 || len(i) == 7 {
				part1++
			}

			out := []rune(i)
			sort.Slice(out, func(i int, j int) bool {
				return out[i] < out[j]
			})
			o := string(out)

			switch x {
			case 0:
				n += 1000 * digits[o]
			case 1:
				n += 100 * digits[o]
			case 2:
				n += 10 * digits[o]
			case 3:
				n += 1 * digits[o]
			}
		}

		part2 += n
		fmt.Println(n)
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
