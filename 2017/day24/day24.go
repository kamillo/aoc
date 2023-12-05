package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

// func main() {
// 	lines := utils.GetLines("input.txt")

// 	components := []interface{}{}

// 	for _, line := range lines {
// 		a, b := 0, 0
// 		fmt.Sscanf(line, "%d/%d", &a, &b)

// 		components = append(components, image.Pt(a, b))
// 	}

// 	perm := utils.HeapPermutation(components)

// 	fmt.Println(len(perm))

// 	max := 0
// 	for _, p := range perm {
// 		prev := image.Point{}
// 		sum := 0

// 		for c := range p {
// 			if c == 0 {
// 				if p[c].(image.Point).X != 0 && p[c].(image.Point).Y == 0 {
// 					break
// 				}
// 			}

// 			if c != 0 {
// 				if p[c].(image.Point).X == prev.X || p[c].(image.Point).Y == prev.Y {
// 					break
// 				}
// 			}

// 			prev = p[c].(image.Point)
// 			sum += p[c].(image.Point).X + p[c].(image.Point).Y
// 		}

// 		if sum > max {
// 			max = sum
// 		}
// 	}
// }

// func connections(a []image.Point) [][]image.Point {
// 	var permutations [][]image.Point
// 	var generate func([]image.Point, int)

// 	generate = func(a []image.Point, size int) {
// 		if size == 1 {
// 			A := make([]image.Point, len(a))
// 			copy(A, a)
// 			permutations = append(permutations, A)
// 		}
// 		for i := 0; i < size; i++ {
// 			generate(a, size-1)
// 			if size%2 == 1 {
// 				a[0], a[size-1] = a[size-1], a[0]
// 			} else {
// 				a[i], a[size-1] = a[size-1], a[i]
// 			}
// 		}
// 	}
// 	generate(a, len(a))
// 	return permutations
// }

type Component struct {
	a, b                   int
	aAvailable, bAvailable bool
}

func main() {
	lines := utils.GetLines("input.txt")

	components := map[string]Component{}
	startComponents := map[string]Component{}

	for _, line := range lines {
		a, b := 0, 0
		fmt.Sscanf(line, "%d/%d", &a, &b)

		c := New(a, b)
		components[line] = c

		if a == 0 || b == 0 {
			startComponents[line] = c
		}
	}

	fmt.Println(components)

	max := 0
	for k, c := range startComponents {
		next := 0
		if c.a == 0 {
			c.aAvailable = false
			next = c.b
		} else {
			c.aAvailable = false
			next = c.a
		}

		used := map[string]bool{}
		used[k] = true

		tmp := map[string]Component{}
		for kk, vv := range components {
			tmp[kk] = vv
		}

		s := getConnection(c, next, tmp, used)
		fmt.Println(s)
		if s > max {
			max = s
		}
	}

	fmt.Println("Part 1:", max)

}

func New(a, b int) Component {
	c := Component{}

	c.a = a
	c.b = b
	c.aAvailable = true
	c.bAvailable = true
	// c.ports = map[int]bool{a: true, b: true}

	return c
}

func getConnection(current Component, port int, components map[string]Component, used map[string]bool) int {
	sum := 0
	for c := range components {
		tmpUsed := map[string]bool{}
		for k, v := range used {
			tmpUsed[k] = v
		}

		tmp := map[string]Component{}
		for kk, vv := range components {
			tmp[kk] = vv
		}

		component := tmp[c]
		// if canConnect(components[c], port) {
		if !used[c] && canConnect(component, port) {
			tmpUsed[c] = true

			if component.a == port {
				component.aAvailable = false
			} else {
				component.bAvailable = false
			}

			if !component.aAvailable && !component.bAvailable {
				fmt.Println(component)
			}

			tmp[c] = component

			sum = getConnection(tmp[c], getFreePort(component), tmp, tmpUsed)
		}
	}

	return sum + current.a + current.b
}

func canConnect(c Component, port int) bool {
	return (c.a == port && c.aAvailable) || (c.b == port && c.bAvailable)
}

func getFreePort(c Component) int {
	if c.aAvailable {
		return c.a
	}

	if c.bAvailable {
		return c.b
	}

	panic("no free port")

	return -1
}
