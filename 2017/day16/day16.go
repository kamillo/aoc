package main

import (
	"container/ring"
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Instruction struct {
	t byte
	a any
	b any
}

func main() {
	input := strings.Split(utils.GetLines("test.txt")[0], ",")
	instructions := []Instruction{}
	for _, i := range input {
		a, b := 0, 0
		inst := Instruction{i[0], 0, 0}
		switch i[0] {
		case 's':
			fmt.Sscanf(i[1:], "%d", &a)

		case 'x':
			fmt.Sscanf(i[1:], "%d/%d", &a, &b)

		case 'p':
			fmt.Sscanf(i[1:], "%c/%c", &a, &b)
		}

		inst.a = a
		inst.b = b
		instructions = append(instructions, inst)
	}

	programs := ring.New(5)
	for i := 0; i < programs.Len(); i++ {
		programs.Value = 'a' + i
		programs = programs.Next()
	}

	print(programs)

	for _, i := range instructions {
		programs = applyInstuction(programs, i)
	}

	print(programs)

	for j := 0; j < 1000000000; j++ {
		for _, i := range instructions {
			programs = applyInstuction(programs, i)
		}
	}
	print(programs)
}

func applyInstuction(ring *ring.Ring, instruction Instruction) *ring.Ring {
	switch instruction.t {
	case 's':
		{
			ring = spin(ring, instruction.a.(int))
		}

	case 'x':
		{
			ring = exchange(ring, instruction.a.(int), instruction.b.(int))
		}

	case 'p':
		{
			ring = swap(ring, instruction.a.(int), instruction.b.(int))
		}
	}

	return ring
}

func spin(ring *ring.Ring, i int) *ring.Ring {
	ring = ring.Move(-i)
	return ring
}

func exchange(ring *ring.Ring, a, b int) *ring.Ring {
	max := utils.Max(a, b)
	aP := ring
	bP := ring
	ret := ring
	for i := 0; i <= max; i++ {
		if i == a {
			aP = ring
		}

		if i == b {
			bP = ring
		}

		ring = ring.Next()
	}

	tmp := aP.Value
	aP.Value = bP.Value
	bP.Value = tmp

	return ret
}

func swap(ring *ring.Ring, a, b int) *ring.Ring {
	aP := ring
	bP := ring
	ret := ring
	for i := 0; i < ring.Len(); i++ {
		if ring.Value == int(a) {
			aP = ring
		}

		if ring.Value == int(b) {
			bP = ring
		}

		ring = ring.Next()
	}

	tmp := aP.Value
	aP.Value = bP.Value
	bP.Value = tmp

	return ret
}

func print(ring *ring.Ring) {
	tmp := ring
	for i := 0; i < tmp.Len(); i++ {
		fmt.Printf("%c", tmp.Value)
		tmp = tmp.Next()
	}

	fmt.Println()
}

// func get(ring *ring.Ring) {
// 	tmp := ring
// 	for i := 0; i < tmp.Len(); i++ {
// 		fmt.Printf("%c", tmp.Value)
// 		tmp = tmp.Next()
// 	}

// 	return s
// }
