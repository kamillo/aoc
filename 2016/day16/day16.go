package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "01111001100111011"
	// input := "10000"
	//size := 272
	size := 35651584

	data := input
	for len(data) < size {
		data = step(data)
	}

	fmt.Println(data[:size])

	fmt.Println(checksum(data[:size]))
}

func step(a string) string {
	b := []byte(a)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	bb := strings.ReplaceAll(string(b), "1", "#")
	bb = strings.ReplaceAll(bb, "0", "1")
	bb = strings.ReplaceAll(bb, "#", "0")

	return a + "0" + bb
}

func checksum(data string) string {
	sum := ""
	for len(sum)%2 == 0 || len(sum) == 0 {
		sum = ""
		for i := 0; i < len(data); i += 2 {
			if data[i] == data[i+1] {
				sum += "1"
			} else {
				sum += "0"
			}
		}
		data = sum
	}

	return sum
}
