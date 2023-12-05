package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "01111001100111011"
	// input := "10000"
	// size := 272
	size := 35651584

	data := input
	for len(data) < size {
		data = step(data)
	}

	fmt.Println(checksum(data[:size]))
}

func step(a string) string {
	b := []byte(a)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	bb := strings.Map(func(r rune) rune {
		if r == '1' {
			return '0'
		}
		return '1'
	}, string(b))

	return a + "0" + bb
}

func checksum(data string) string {
	builder := strings.Builder{}
	sum := ""
	for len(sum)%2 == 0 || len(sum) == 0 {
		builder.Reset()
		builder.Grow(len(data) / 2)
		for i := 0; i < len(data); i += 2 {
			if data[i] == data[i+1] {
				builder.WriteString("1")
			} else {
				builder.WriteString("0")
			}
		}
		sum = builder.String()
		data = sum
	}

	return sum
}
