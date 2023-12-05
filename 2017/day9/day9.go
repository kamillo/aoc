package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	// input := "{{<!!>},{<!!>},{<!!>},{<!!>}}"
	input := utils.GetLines("input.txt")[0]
	fmt.Println(getScore(input))
}

const (
	nothing int = iota
	group
	garbage
	escape
)

func getScore(str string) (int, int) {
	state := Stack{nothing}
	level := 0
	score := 0
	garbageCount := 0

	for _, c := range str {
		if state.Peek() == garbage && c != '!' && c != '>' {
			garbageCount++
		}

		switch c {
		case '{':
			if state.Peek() != garbage && state.Peek() != escape {
				state.Push(group)
				level++
				score += level
			} else if state.Peek() == escape {
				state.Pop()
			}
		case '}':
			if state.Peek() != garbage && state.Peek() != escape {
				state.Pop()
				level--
			} else if state.Peek() == escape {
				state.Pop()
			}
		case '<':
			if state.Peek() != garbage && state.Peek() != escape {
				state.Push(garbage)
			} else if state.Peek() == escape {
				state.Pop()
			}
		case '>':
			if state.Peek() != escape {
				state.Pop()
			} else {
				state.Pop()
			}
		case '!':
			if state.Peek() != escape {
				state.Push(escape)
			} else {
				state.Pop()
			}
		default:
			if state.Peek() == escape {
				state.Pop()
			}
		}

	}

	return score, garbageCount
}

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val int) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.

		return element, true
	}
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	} else {
		index := len(*s) - 1
		element := (*s)[index]

		return element
	}
}
