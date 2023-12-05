package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Node struct {
	val     int
	monkeys []string
	op      func(m1, m2 string) int
}

type Tree map[string]*Node

func main() {
	lines := utils.GetLines("input.txt")

	tree := Tree{}
	for _, line := range lines {
		n := Node{}
		name := ""
		op, m1, m2 := "", "", ""
		val := 0
		if _, err := fmt.Sscanf(line, "%s %s %s %s", &name, &m1, &op, &m2); err == nil {

			switch op {
			case "+":
				n.op = tree.add
			case "-":
				n.op = tree.sub
			case "*":
				n.op = tree.mul
			case "/":
				n.op = tree.div
			}

			n.monkeys = append(n.monkeys, m1, m2)
			n.val = math.MaxInt

		} else {
			fmt.Sscanf(line, "%s %d", &name, &val)
			n.val = val
		}

		name = strings.ReplaceAll(name, ":", "")
		tree[name] = &n
	}

	fmt.Println(tree.getValue("root"))
}

func (t Tree) getValue(m string) int {
	if t[m].val != math.MaxInt {
		return t[m].val
	}

	v := t[m].op(t[m].monkeys[0], t[m].monkeys[1])
	t[m].val = v

	return v
}

func (t Tree) add(m1, m2 string) int {
	a := t.getValue(m1)
	b := t.getValue(m2)

	return a + b
}

func (t Tree) sub(m1, m2 string) int {
	a := t.getValue(m1)
	b := t.getValue(m2)

	return a - b
}

func (t Tree) mul(m1, m2 string) int {
	a := t.getValue(m1)
	b := t.getValue(m2)

	return a * b
}

func (t Tree) div(m1, m2 string) int {
	a := t.getValue(m1)
	b := t.getValue(m2)

	return a / b
}
