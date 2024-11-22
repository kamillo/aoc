package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Monkey struct {
	items     []int
	next      int
	prev      int
	operation func(a, b int) int
	b         int
	div       int
}

func main() {
	lines := utils.GetLines("input.txt")

	monkeys1 := parse(lines)
	fmt.Println("Part 1:", part1(monkeys1))

	monkeys2 := parse(lines)
	fmt.Println("Part 2:", part2(monkeys2))
}

func part1(monkeys map[int]*Monkey) int {
	insp := map[int]int{}

	for i := 0; i < 20; i++ {
		for i := 0; i < len(monkeys); i++ {
			for _, item := range monkeys[i].items {
				insp[i]++

				op := monkeys[i].b
				if monkeys[i].b == -1 {
					op = item
				}

				new := int(monkeys[i].operation(item, op) / 3)

				if new%monkeys[i].div == 0 {
					monkeys[monkeys[i].next].items = append(monkeys[monkeys[i].next].items, new)
				} else {
					monkeys[monkeys[i].prev].items = append(monkeys[monkeys[i].prev].items, new)
				}
			}

			monkeys[i].items = []int{}
		}
	}

	return utils.SortMapIntInt(insp, true)[0].Value * utils.SortMapIntInt(insp, true)[1].Value
}

func part2(monkeys map[int]*Monkey) int {
	insp := map[int]int{}

	mod := 1
	for _, m := range monkeys {
		mod *= m.div
	}

	for i := 0; i < 10000; i++ {
		for i := 0; i < len(monkeys); i++ {
			for _, item := range monkeys[i].items {
				insp[i]++

				op := monkeys[i].b
				if monkeys[i].b == -1 {
					op = item
				}

				new := int(monkeys[i].operation(item, op)) % mod

				if new%monkeys[i].div == 0 {
					monkeys[monkeys[i].next].items = append(monkeys[monkeys[i].next].items, new)
				} else {
					monkeys[monkeys[i].prev].items = append(monkeys[monkeys[i].prev].items, new)
				}
			}

			monkeys[i].items = []int{}
		}
	}

	return utils.SortMapIntInt(insp, true)[0].Value * utils.SortMapIntInt(insp, true)[1].Value
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func parse(lines []string) map[int]*Monkey {
	monkeys := map[int]*Monkey{}
	index := 0

	for _, line := range lines {
		tmp := 0
		if _, err := fmt.Sscanf(line, "Monkey %d:", &index); err == nil {
			monkeys[index] = &Monkey{}
		} else {
			if strings.HasPrefix(line, "  Starting") {
				items := strings.Split(strings.Split(line, ": ")[1], ", ")

				for _, i := range items {
					item, err := strconv.Atoi(i)
					if err != nil {
						panic("atoi")
					}
					monkeys[index].items = append(monkeys[index].items, item)
				}
			} else if strings.HasPrefix(line, "  Operation:") {
				b := 0
				if _, err := fmt.Sscanf(line, "  Operation: new = old + %d", &b); err == nil {
					monkeys[index].b = b
					monkeys[index].operation = add
				} else if _, err := fmt.Sscanf(line, "  Operation: new = old * %d", &b); err == nil {
					monkeys[index].b = b
					monkeys[index].operation = mul
				} else if _, err := fmt.Sscanf(line, "  Operation: new = old * old"); err == nil {
					monkeys[index].b = -1
					monkeys[index].operation = mul
				} else {
					panic("op")
				}

			} else if _, err := fmt.Sscanf(line, "  Test: divisible by %d", &tmp); err == nil {
				monkeys[index].div = tmp

			} else if _, err := fmt.Sscanf(line, "    If true: throw to monkey %d", &tmp); err == nil {
				monkeys[index].next = tmp
			} else if _, err := fmt.Sscanf(line, "    If false: throw to monkey %d", &tmp); err == nil {
				monkeys[index].prev = tmp
			}
		}
	}

	return monkeys
}
