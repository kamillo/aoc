package main

import (
	"fmt"
	"math"
)

func main() {
	input := 34000000

	for i := 1; ; i++ {
		presents := 0
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				if i/j == j {
					presents += j * 10
				} else {
					presents += i / j * 10
					presents += j * 10
				}
			}
		}

		if presents >= input {
			fmt.Println("Part 1: ", i, presents)
			break
		}
	}

	elfs := make(map[int]int)
	for i := 1; ; i++ {
		presents := 0
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				if i/j == j {
					if elfs[j] < 50 {
						elfs[j]++
						presents += j * 11
					}
				} else {
					if elfs[i/j] < 50 {
						elfs[i/j]++
						presents += i / j * 11
					}
					if elfs[j] < 50 {
						elfs[j]++
						presents += j * 11
					}
				}
			}
		}

		if presents >= input {
			fmt.Println("Part 2: ", i, presents)
			break
		}
	}
}
