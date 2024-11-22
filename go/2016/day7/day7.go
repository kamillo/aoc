package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	cnt := 0
	cnt2 := 0
	for _, line := range lines {
		s := line
		supernetSeqs := []string{}
		hypernetSeqs := []string{}

		for {
			before, after, found := strings.Cut(s, "[")
			supernetSeqs = append(supernetSeqs, before)
			if !found {
				break
			}

			brackets, rest, _ := strings.Cut(after, "]")
			hypernetSeqs = append(hypernetSeqs, brackets)
			s = rest
		}

		abba := checkAbba(supernetSeqs)
		abbaInBrackets := checkAbba(hypernetSeqs)

		if abba && !abbaInBrackets {
			cnt++
		}

		if checkAbaBab(supernetSeqs, hypernetSeqs) {
			cnt2++
		}
		//fmt.Println(valid)
	}

	fmt.Println("Part 1:", cnt)
	fmt.Println("Part 1:", cnt2)
}

func checkAbba(text []string) bool {
	for _, t := range text {
		for i := 0; i <= len(t)-4; i++ {
			slice := t[i : i+4]
			if slice[0] == slice[3] && slice[1] == slice[2] && slice[0] != slice[1] {
				return true
			}
		}
	}
	return false
}

func checkAbaBab(supernetSeqs, hypernetSeqs []string) bool {
	for _, t := range supernetSeqs {
		for i := 0; i <= len(t)-3; i++ {
			slice := t[i : i+3]
			if slice[0] == slice[2] && slice[0] != slice[1] {
				for _, h := range hypernetSeqs {
					for j := 0; j <= len(h)-3; j++ {
						slice2 := h[j : j+3]
						if slice2[0] == slice[1] && slice2[2] == slice[1] && slice2[1] == slice[0] {
							return true
						}
					}
				}
			}
		}
	}

	return false
}
