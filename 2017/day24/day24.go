package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

type Component struct {
	port0 int
	port1 int
	used  bool
}

var components []Component
var max int
var maxLength int

func main() {
	components = []Component{}

	lines := utils.GetLines("input.txt")
	for _, line := range lines {
		a, b := 0, 0
		fmt.Sscanf(line, "%d/%d", &a, &b)
		components = append(components, Component{a, b, false})
	}

	max = 0
	findPath1(0, 0, 0)
	fmt.Println("Part 1:", max)
}

func findPath1(plug, accum, length int) {
	if max < accum {
		max = accum
	}
	for i, component := range components {
		if component.used {
			continue
		}
		if component.port0 == plug {
			components[i].used = true
			findPath1(component.port1, accum+component.port0+component.port1, length+1)
			components[i].used = false
		}
		if component.port1 == plug {
			components[i].used = true
			findPath1(component.port0, accum+component.port0+component.port1, length+1)
			components[i].used = false
		}
	}
}

func findPath2(plug, accum, length int) {
	if max < accum {
		max = accum
	}
	for i, component := range components {
		if component.used {
			continue
		}
		if component.port0 == plug {
			components[i].used = true
			findPath1(component.port1, accum+component.port0+component.port1, length+1)
			components[i].used = false
		}
		if component.port1 == plug {
			components[i].used = true
			findPath1(component.port0, accum+component.port0+component.port1, length+1)
			components[i].used = false
		}
	}
}
