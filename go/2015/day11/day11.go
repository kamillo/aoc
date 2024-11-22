package main

import (
	"fmt"
	"strings"
)

// "strings"

func main() {
	input := "cqjxjnds"
	i := len(input) - 1
	stop := 0

	for stop < 2 {
		if allowed(input) {
			fmt.Println(input)
			stop++
		}

		letter := input[i]
		if 'a'+(letter+1-'a')%26 == 'a' {
			for j := i; j >= 0; j-- {
				input = replaceAtIndex(input, rune('a'+(input[j]+1-'a')%26), j)
				if 'a'+(input[j]-'a')%26 != 'a' {
					break
				}
			}
		}
		input = replaceAtIndex(input, rune('a'+(letter+1-'a')%26), i)
	}
}

func allowed(pass string) bool {
	last := rune(pass[0])
	pairs := map[rune]bool{}
	straight := 0
	straights := false

	for i, d := range pass {
		if strings.Contains("iol", string(d)) {
			return false
		}

		if i != 0 && d == last {
			pairs[d] = true
		}

		if i != 0 && d-last == 1 {
			straight++
		} else {
			straight = 0
		}

		if straight == 2 {
			straights = true
		}

		last = d
	}
	return len(pairs) >= 2 && straights
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
