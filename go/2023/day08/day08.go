package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Node struct {
	Left  string
	Right string
}

func main() {
	lines := utils.GetLines("input.txt")

	seq := lines[0]
	lines = lines[2:]
	network := map[string]Node{}

	for _, line := range lines {
		key := strings.Split(line, " = ")[0]
		value := strings.Split(line, " = ")[1]
		l := strings.Split(value, ", ")[0]
		l = l[1:]
		r := strings.Split(value, ", ")[1]
		r = r[:len(r)-1]

		network[key] = Node{l, r}
	}

	calc := func(node string) int {
		step := 0
		for !strings.HasSuffix(node, "Z") {
			for _, s := range seq {
				if s == 'L' {
					node = network[node].Left
				} else {
					node = network[node].Right
				}
				step++
			}
		}

		return step
	}

	fmt.Println("Part 1:", calc("AAA"))

	steps := []int{}
	for n := range network {
		if strings.HasSuffix(n, "A") {
			steps = append(steps, calc(n))
		}
	}

	fmt.Println("Part 2:", utils.Lcm(steps[0], steps[1], steps[2:]...))
}
