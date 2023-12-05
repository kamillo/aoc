package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

type State struct {
	value     int
	move      int
	nextState string
}

func main() {
	lines := utils.GetLines("input.txt")

	states := map[string]map[int]State{}
	steps := 0

	state := ""
	value := 0

	for _, line := range lines {
		if _, err := fmt.Sscanf(line, "Perform a diagnostic checksum after %d steps.", &steps); err == nil {
			continue
		}

		if _, err := fmt.Sscanf(line, "In state %1s:", &state); err == nil {
			states[state] = map[int]State{}
			continue
		}

		if _, err := fmt.Sscanf(line, "  If the current value is %d", &value); err == nil {
			states[state][value] = State{}
			continue
		}

		v := 0
		if _, err := fmt.Sscanf(line, "    - Write the value %d.", &v); err == nil {
			s := states[state][value]
			s.value = v
			states[state][value] = s
			continue
		}

		m := ""
		if _, err := fmt.Sscanf(line, "    - Move one slot to the %s", &m); err == nil {
			s := states[state][value]
			if m == "right." {
				s.move = 1

			} else {
				s.move = -1
			}

			states[state][value] = s
			continue
		}

		n := ""
		if _, err := fmt.Sscanf(line, "    - Continue with state %1s.", &n); err == nil {
			s := states[state][value]
			s.nextState = n
			states[state][value] = s
			continue
		}
	}
	fmt.Println(states)

	ptr := 10000
	tape := [30000]int{}
	state = "A"
	for i := 0; i < steps; i++ {
		v := tape[ptr]

		tape[ptr] = states[state][v].value
		ptr += states[state][v].move
		state = states[state][v].nextState
	}

	fmt.Println(utils.SumInts(tape[:]))
}
