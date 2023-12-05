package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	notes := utils.GetLines("test.txt")

	part1 := 0
	for _, line := range notes {
		split := strings.Split(line, " | ")

		pattern := strings.Fields(split[0])
		output := strings.Fields(split[1])

		connection := map[string]string{}
		digits := map[string]int{}
		missing := []string{}
		for _, p := range pattern {
			switch len(p) {
			case 2:
				digits[p] = 1
				connection["c"] = string(p[0])
				connection["f"] = string(p[1])
			case 3:
				digits[p] = 7
				connection["a"] = string(p[0])
				connection["c"] = string(p[1])
				connection["f"] = string(p[2])
			case 4:
				digits[p] = 4
				connection["b"] = string(p[0])
				connection["c"] = string(p[1])
				connection["d"] = string(p[2])
				connection["f"] = string(p[3])
			case 7:
				digits[p] = 8
				connection["a"] = string(p[0])
				connection["b"] = string(p[1])
				connection["c"] = string(p[2])
				connection["d"] = string(p[3])
				connection["e"] = string(p[4])
				connection["f"] = string(p[5])
				connection["g"] = string(p[6])
			default:
				missing = append(missing, p)
			}
		}

		// 6: 0, 6, 9
		// 5: 2, 3, 5
		for _, m := range missing {
			if len(m) == 5 {
				c := connection["c"]
				f := connection["f"]
				// 2 - c:1, f:0
				if strings.Contains(m, c) && !strings.Contains(m, f) {
					digits[m] = 2
				}

				// 3 - c:1, f:1
				if strings.Contains(m, c) && strings.Contains(m, f) {
					digits[m] = 3
				}

				// 5 - c:0, f:1
				if !strings.Contains(m, c) && strings.Contains(m, f) {
					digits[m] = 5
				}
			}

			if len(m) == 6 {
				c := connection["c"]
				d := connection["d"]
				e := connection["e"]

				// 0 - e:1, d:0
				if strings.Contains(m, e) && !strings.Contains(m, d) {
					digits[m] = 0
				}

				// 6 - c:0
				if !strings.Contains(m, c) {
					digits[m] = 6
				}

				// 9 - e:0, d:1
				if !strings.Contains(m, e) && strings.Contains(m, d) {
					digits[m] = 9
				}
			}
		}

		// fmt.Println(missing)
		// fmt.Println(connection)
		fmt.Println(digits)
		for _, i := range output {
			if len(i) == 2 || len(i) == 3 || len(i) == 4 || len(i) == 7 {
				part1++
			}

			fmt.Print(i, digits[i])
		}
		fmt.Println()
	}

	fmt.Println("Part 1: ", part1)
}
