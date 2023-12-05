package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	stacks1 := parseStack(lines)
	stacks2 := parseStack(lines)

	for _, line := range lines {
		crates := 0
		a, b := 0, 0

		if _, err := fmt.Sscanf(line, "move %d from %d to %d", &crates, &a, &b); err == nil {
			for i := 0; i < crates; i++ {
				stackA := stacks1[a]
				stackB := stacks1[b]
				v, _ := stackA.Pop()
				stackB.Push(v)

				stacks1[a] = stackA
				stacks1[b] = stackB
			}
			// part2
			{
				stackA := stacks2[a]
				stackB := stacks2[b]
				x := len(stackA) - crates
				stackB = append(stackB, stacks2[a][x:]...)

				stacks2[a] = stackA[:x]
				stacks2[b] = stackB
			}
		}
	}

	fmt.Print("Part 1: ")
	for i := 1; i <= len(stacks1); i++ {
		s := stacks1[i]
		fmt.Print(s.Peek())
	}
	fmt.Println()

	fmt.Print("Part 2: ")
	for i := 1; i <= len(stacks2); i++ {
		s := stacks2[i]
		fmt.Print(s.Peek())
	}
}

func parseStack(lines []string) map[int]Stack {
	stacks := map[int]Stack{}

	for _, line := range lines {
		if line[1] == '1' {
			return stacks
		}

		s := 1
		for i := 1; i < len(line); i += 4 {
			if line[i] != ' ' {
				stack := stacks[s]
				stack.PushFront(string(line[i]))
				stacks[s] = stack
			}
			s++
		}
		// for i := 1; i <= len(stacks); i++ {
		// 	s := stacks[i]
		// 	fmt.Println(i, s)
		// }
	}

	return stacks
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) PushFront(val string) {
	*s = append(Stack{val}, *s...)
}

func (s *Stack) Push(val string) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]

		return element, true
	}
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]

		return element
	}
}
