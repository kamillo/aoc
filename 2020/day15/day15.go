package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	previous, recent int
}

func main() {
	input := "12,1,16,3,11,0"
	numbersMap := make(map[int]Pair)
	for i, s := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(s)
		numbersMap[n] = Pair{i + 1, i + 1}
	}

	last := 0
	for i := len(numbersMap); ; i++ {
		if _, ok := numbersMap[last]; ok {
			numbersMap[last] = Pair{numbersMap[last].recent, i}
			last = numbersMap[last].recent - numbersMap[last].previous
		} else {
			numbersMap[last] = Pair{i, i}
			last = 0
		}

		if i == 2020-1 {
			fmt.Println("Part 1: ", last)
		}
		if i == 30000000-1 {
			fmt.Println("Part 2: ", last)
			break
		}
	}
}
