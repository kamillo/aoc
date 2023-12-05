package main

import (
	"encoding/json"
	"fmt"
	"github.com/kamillo/aoc/utils"
)

func main() {
	charsLength := 0
	stringLength := 0
	marshaled := 0
	for i, line := range utils.GetLines("input.txt") {
		charsLength += len(line)
		stringLength += len(lines[i])
		n, _ := json.Marshal(line)
		marshaled += len(n)
	}
	fmt.Println("Part 1: ", charsLength-stringLength)
	fmt.Println("Part 2: ", marshaled-charsLength)
}
