package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	len, _ := decompress1(lines[0])
	fmt.Println("Part 1:", len)

	len, _ = decompress2(lines[0])
	fmt.Println("Part 2:", len)
}

func decompress2(substr string) (int, string) {
	// var builder strings.Builder
	decompressLen := 0
	marker := false
	repeater := false
	repeat := ""
	sub := ""

	for i := 0; i < len(substr); i++ {
		c := substr[i]

		if c == '(' {
			marker = true
		}

		if c == ')' {
			n, _ := strconv.Atoi(sub)
			r, _ := strconv.Atoi(repeat)

			substring := substr[i+1 : i+1+n]
			substring = strings.Repeat(substring, r)

			next := strings.Index(substring, "(")

			if next > -1 {
				nn, _ := decompress2(substring)
				decompressLen += nn
				// nn, ss := decompress(substring)
				//builder.Grow(nn)
				//builder.WriteString(ss)

			} else {
				decompressLen += n * r
				// builder.Grow(n * r)
				// builder.WriteString(substring)
			}

			i += n

			marker = false
			repeater = false
			sub = ""
			repeat = ""
		}

		if c >= 'A' && c <= 'Z' {
			decompressLen++
			// builder.Grow(1)
			// builder.WriteByte(c)
		}

		if c >= '0' && c <= '9' {
			if !repeater {
				sub += string(c)
			} else {
				repeat += string(c)
			}
		}

		if c == 'x' {
			if marker {
				repeater = true
			}
		}
	}

	//return builder.Len(), builder.String()
	return decompressLen, ""
}

func decompress1(substr string) (int, string) {
	var builder strings.Builder
	marker := false
	repeater := false
	repeat := ""
	sub := ""

	for i := 0; i < len(substr); i++ {
		c := substr[i]

		if c == '(' {
			marker = true
		}

		if c == ')' {
			n, _ := strconv.Atoi(sub)
			r, _ := strconv.Atoi(repeat)

			substring := substr[i+1 : i+1+n]
			substring = strings.Repeat(substring, r)

			builder.Grow(n * r)
			builder.WriteString(substring)

			i += n

			marker = false
			repeater = false
			sub = ""
			repeat = ""
		}

		if c >= 'A' && c <= 'Z' {
			builder.Grow(1)
			builder.WriteByte(c)
		}

		if c >= '0' && c <= '9' {
			if !repeater {
				sub += string(c)
			} else {
				repeat += string(c)
			}
		}

		if c == 'x' {
			if marker {
				repeater = true
			}
		}
	}

	return builder.Len(), builder.String()
}
