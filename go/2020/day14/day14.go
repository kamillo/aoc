package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strconv"
	"strings"
)

func main() {
	lines := utils.GetLines("input.txt")

	var maskOr int64
	var maskAnd int64
	mem := make(map[int]int64)
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			maskString := strings.Split(line, " = ")[1]
			maskStringOr := strings.ReplaceAll(maskString, "X", "0")
			maskStringAnd := strings.ReplaceAll(maskString, "X", "1")
			maskOr, _ = strconv.ParseInt(maskStringOr, 2, 0)
			maskAnd, _ = strconv.ParseInt(maskStringAnd, 2, 0)
		} else {
			address := 0
			var value int64
			fmt.Sscanf(line, "mem[%d] = %d", &address, &value)
			mem[address] = (value | maskOr) & maskAnd
		}
	}

	sum := int64(0)
	for _, value := range mem {
		sum += value
	}
	fmt.Println("Part 1: ", sum)

	mem2 := make(map[int]int64)
	var mask string
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = strings.Split(line, " = ")[1]
		} else {
			address := 0
			var value int64
			fmt.Sscanf(line, "mem[%d] = %d", &address, &value)
			for i, f := 0, strings.Count(mask, "X"); i < 1<<f; i++ {
				mask := strings.NewReplacer("X", "F", "0", "X").Replace(mask)
				for _, r := range fmt.Sprintf("%0*b", f, i) {
					mask = strings.Replace(mask, "F", string(r), 1)
				}

				maskStringOr := strings.ReplaceAll(mask, "X", "0")
				maskStringAnd := strings.ReplaceAll(mask, "X", "1")
				maskOr, _ = strconv.ParseInt(maskStringOr, 2, 64)
				maskAnd, _ = strconv.ParseInt(maskStringAnd, 2, 64)
				address = (address | int(maskOr)) & int(maskAnd)
				mem2[address] = value
			}
		}
	}

	sum = 0
	for _, i := range mem2 {
		sum += i
	}
	fmt.Println("Part 2: ", sum)
}
