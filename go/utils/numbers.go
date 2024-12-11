package utils

import "strconv"

func NumDigits(n int) int {
	return len(strconv.Itoa(Abs(n)))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
