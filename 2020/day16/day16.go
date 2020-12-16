package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"strconv"
	"strings"
)

type Range struct {
	name     string
	min, max int
}

func main() {
	lines := fileutil.GetLines("input.txt")
	//lines := fileutil.GetLines("test.txt")
	//lines := fileutil.GetLines("test2.txt")

	var myTicketValues []int
	var nearbyTicketValues [][]int
	near := 0
	var allRanges [][2]Range
	for _, line := range lines {
		if len(line) == 0 || line == "your ticket:" || line == "nearby tickets:" {
			continue
		}

		if strings.Contains(line, ":") {
			field := ""
			field2 := ""
			var a, b, c, d int

			if _, err := fmt.Sscanf(line, "%s %d-%d or %d-%d", &field, &a, &b, &c, &d); err != nil {
				fmt.Sscanf(line, "%s %s %d-%d or %d-%d", &field, &field2, &a, &b, &c, &d)
			}

			allRanges = append(allRanges, [2]Range{{field + field2, a, b}, {field + field2, c, d}})

		} else if len(myTicketValues) == 0 {
			values := strings.Split(line, ",")
			for i := range values {
				n, _ := strconv.Atoi(values[i])
				myTicketValues = append(myTicketValues, n)
			}
		} else {
			nearbyTicketValues = append(nearbyTicketValues, []int{})
			values := strings.Split(line, ",")
			for i := range values {
				n, _ := strconv.Atoi(values[i])
				nearbyTicketValues[near] = append(nearbyTicketValues[near], n)
			}
			near++
		}
	}

	sum := 0
	matching := make([][]bool, len(allRanges))
	for i := range matching {
		matching[i] = make([]bool, len(allRanges))
		for j := range matching[i] {
			matching[i][j] = true
		}
	}

	for _, nt := range nearbyTicketValues {
		validTicket := true
		for _, t := range nt {
			valid := false
			for _, r := range allRanges {
				if (t >= r[0].min && t <= r[0].max) || (t >= r[1].min && t <= r[1].max) {
					valid = true
				}
			}
			if !valid {
				sum += t
			}
			validTicket = validTicket && valid
		}
		if validTicket {
			for i, v := range nt {
				for j, r := range allRanges {
					if !((v >= r[0].min && v <= r[0].max) || (v >= r[1].min && v <= r[1].max)) {
						matching[i][j] = false
					}
				}
			}
		}
	}

	fmt.Println("Part 1: ", sum)

	used := make([]bool, len(allRanges))
	fields := make(map[int]int)
	for {
		for i := range used {
			validTmp := make([]int, 0)
			for j := range used {
				if !used[j] && matching[i][j] {
					validTmp = append(validTmp, j)
				}
			}
			if len(validTmp) == 1 {
				fields[i] = validTmp[0]
				used[validTmp[0]] = true
			}
		}
		if len(fields) == len(used) {
			break
		}
	}

	p2 := 1
	for i, j := range fields {
		if j < 6 {
			p2 *= myTicketValues[i]
		}
	}
	fmt.Println("Part 2: ", p2)
}
