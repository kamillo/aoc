package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	cnt := 0
	for _, line := range utils.GetLines("input.txt") {
		a, b, c := 0, 0, 0
		fmt.Sscanf(line, "%d %d %d", &a, &b, &c)

		if b+c > a && a+c > b && a+b > c {
			cnt++
		}
	}
	fmt.Println("Part 1: ", cnt)

	cnt = 0
	aa, bb, cc := [3]int{}, [3]int{}, [3]int{}
	for i, line := range utils.GetLines("input.txt") {
		a, b, c := 0, 0, 0
		fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
		aa[i%3] = a
		bb[i%3] = b
		cc[i%3] = c

		if i%3 == 2 {
			if aa[0]+aa[1] > aa[2] && aa[1]+aa[2] > aa[0] && aa[2]+aa[0] > aa[1] {
				cnt++
			}

			if bb[0]+bb[1] > bb[2] && bb[1]+bb[2] > bb[0] && bb[2]+bb[0] > bb[1] {
				cnt++
			}

			if cc[0]+cc[1] > cc[2] && cc[1]+cc[2] > cc[0] && cc[2]+cc[0] > cc[1] {
				cnt++
			}
			fmt.Println(aa, bb, cc)
		}
	}
	fmt.Println("Part 2: ", cnt)
}
