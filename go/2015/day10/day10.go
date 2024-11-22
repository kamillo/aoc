package main

import (
	"fmt"
)

func main() {
	input := "1321131112"

	d := []rune(input)
	for i := 0; i < 50; i++ {
		d = lookAndSay(d)
	}

	fmt.Println(len(string(d)))
}

func lookAndSay(input []rune) []rune {
	last := input[0]
	cnt := 0
	ret := []rune{}

	for _, digit := range input {
		if digit == last {
			cnt++
		} else {
			ret = append(ret, rune('0'+cnt))
			ret = append(ret, last)
			cnt = 1
		}
		last = digit
	}

	ret = append(ret, rune('0'+cnt))
	ret = append(ret, last)

	return ret
}
