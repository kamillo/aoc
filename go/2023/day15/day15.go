package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type instruction struct {
	label string
	value int
}

func main() {
	lines := utils.GetLines("input.txt")

	sum := 0
	for _, step := range strings.Split(lines[0], ",") {
		current := 0
		for _, c := range step {
			current = hash(current, c)
		}
		sum += current
	}

	fmt.Println("Part 1:", sum)

	hashmap := make([][]instruction, 256)
	for i := range hashmap {
		hashmap[i] = []instruction{}
	}
	for _, step := range strings.Split(lines[0], ",") {
		current := 0
		label := ""
		value := 0
		add := false

		if strings.Contains(step, "=") {
			label = strings.Split(step, "=")[0]
			value = utils.JustAtoi(strings.Split(step, "=")[1])
			add = true

		} else {
			label = strings.Split(step, "-")[0]
		}

		for _, c := range label {
			current = hash(current, c)
		}

		if add {
			if i := findInstruction(hashmap[current], label); i == -1 {
				hashmap[current] = append(hashmap[current], instruction{label, value})
			} else {
				hashmap[current][i].value = value
			}
		} else {
			hashmap[current] = removeFromSlice(hashmap[current], label)
		}
	}

	sum = 0
	for i, m := range hashmap {
		for j, v := range m {
			box := (i + 1) * (j + 1) * v.value
			sum += box
		}
	}

	fmt.Println("Part 2:", sum)
}

func hash(c int, r rune) int {
	c += int(r)
	c *= 17
	c %= 256
	return c
}

func removeFromSlice(s []instruction, val string) []instruction {
	for i, v := range s {
		if v.label == val {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func findInstruction(s []instruction, val string) int {
	for i, v := range s {
		if v.label == val {
			return i
		}
	}
	return -1
}
