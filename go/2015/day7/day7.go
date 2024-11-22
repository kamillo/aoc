package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strconv"
	"strings"
)

func main() {
	defined, wires := parseFile()

	connectWires(wires, defined)
	fmt.Println("Part 1: ", defined["a"])

	defined, wires = parseFile()
	defined["b"] = 956
	connectWires(wires, defined)
	fmt.Println("Part 2: ", defined["a"])
}

func applyOperator(a uint16, b uint16, op string) uint16 {
	switch op {
	case "NOT":
		return ^a
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "LSHIFT":
		return a << b
	case "RSHIFT":
		return a >> b
	}

	return 0
}

func connectWires(wires map[string][]string, defined map[string]uint16) {
	for len(wires) > 0 {
		for key, con := range wires {
			val := uint16(0)
			newVal := uint16(0)
			found := false

			if len(con) == 1 {
				if _, ok := defined[con[0]]; ok {
					newVal = defined[con[0]]
					found = true
				}
			}
			if len(con) == 2 {
				if _, ok := defined[con[1]]; ok {
					val = defined[con[1]]
					newVal = applyOperator(val, 0, con[0])
					found = true
				}
			} else if len(con) == 3 {
				if _, ok := defined[con[0]]; ok {
					val = defined[con[0]]
					if _, ok := defined[con[2]]; ok {
						newVal = applyOperator(val, defined[con[2]], con[1])
						found = true
					} else {
						if a, err := strconv.ParseUint(con[0], 10, 16); err == nil {
							newVal = applyOperator(uint16(a), val, con[1])
							found = true
						} else if b, err := strconv.ParseUint(con[2], 10, 16); err == nil {
							newVal = applyOperator(val, uint16(b), con[1])
							found = true
						}
					}
				} else if _, ok := defined[con[2]]; ok {
					val = defined[con[2]]
					if a, err := strconv.ParseUint(con[0], 10, 16); err == nil {
						newVal = applyOperator(uint16(a), val, con[1])
						found = true
					} else if b, err := strconv.ParseUint(con[2], 10, 16); err == nil {
						newVal = applyOperator(uint16(b), val, con[1])
						found = true
					}
				}
			}

			if found {
				defined[key] = newVal
				delete(wires, key)
			}
		}
	}
}

func parseFile() (defined map[string]uint16, wires map[string][]string) {
	defined = map[string]uint16{}
	wires = map[string][]string{}
	for _, line := range utils.GetLines("input.txt") {
		left := strings.Split(line, " -> ")[0]
		right := strings.Split(line, " -> ")[1]

		if n, err := strconv.ParseUint(left, 10, 16); err == nil {
			defined[right] = uint16(n)
		} else {
			wires[right] = strings.Split(left, " ")
		}
	}

	return defined, wires
}
