package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/fileutil"
)

type vector3d struct {
	x int
	y int
	z int
}

type square struct {
	tl vector3d
	br vector3d
}

type state struct {
	position vector3d
	distance int
}

type portal struct {
	label string
	from  vector3d
	to    vector3d
	outer bool
}

func getGrid() (grid map[vector3d]bool, portals map[vector3d]portal, start, end vector3d) {
	lines := fileutil.GetLines("input.txt")

	outer, inner := square{tl: vector3d{2, 2, 0}, br: vector3d{len(lines[0]) - 3, len(lines) - 3, 0}}, square{}
	for y := outer.tl.y; ; y++ {
		if strings.Contains(lines[y][outer.tl.x:outer.br.x], " ") {
			inner.tl = vector3d{strings.Index(lines[y][outer.tl.x:outer.br.x], " ") + outer.tl.x - 1, y - 1, 0}
			inner.br = vector3d{outer.br.x - inner.tl.x + outer.tl.x, outer.br.y - inner.tl.y + outer.tl.y, 0}
			break
		}
	}

	grid, portals = make(map[vector3d]bool), make(map[vector3d]portal)
	for y := outer.tl.y; y <= outer.br.y; y++ {
		for x := outer.tl.x; x <= outer.br.x; x++ {
			if lines[y][x] == '.' {
				grid[vector3d{x, y, 0}] = true

				var label string
				var pos vector3d
				var outerPortal bool
				if y == outer.tl.y {
					label = lines[y-2][x:x+1] + lines[y-1][x:x+1]
					pos = vector3d{x, y - 1, 0}
					outerPortal = true
				} else if y == outer.br.y {
					label = lines[y+1][x:x+1] + lines[y+2][x:x+1]
					pos = vector3d{x, y + 1, 0}
					outerPortal = true
				} else if x == outer.tl.x {
					label = lines[y][x-2 : x]
					pos = vector3d{x - 1, y, 0}
					outerPortal = true
				} else if x == outer.br.x {
					label = lines[y][x+1 : x+3]
					pos = vector3d{x + 1, y, 0}
					outerPortal = true
				} else if y == inner.br.y && x > inner.tl.x && x < inner.br.x {
					label = lines[y-2][x:x+1] + lines[y-1][x:x+1]
					pos = vector3d{x, y - 1, 0}
				} else if y == inner.tl.y && x > inner.tl.x && x < inner.br.x {
					label = lines[y+1][x:x+1] + lines[y+2][x:x+1]
					pos = vector3d{x, y + 1, 0}
				} else if x == inner.br.x && y > inner.tl.y && y < inner.br.y {
					label = lines[y][x-2 : x]
					pos = vector3d{x - 1, y, 0}
				} else if x == inner.tl.x && y > inner.tl.y && y < inner.br.y {
					label = lines[y][x+1 : x+3]
					pos = vector3d{x + 1, y, 0}
				}

				if label == "AA" {
					start = vector3d{x, y, 0}
				} else if label == "ZZ" {
					end = vector3d{x, y, 0}
				} else if label != "" {
					portals[pos] = portal{label: label, from: vector3d{x, y, 0}, outer: outerPortal}
					grid[pos] = true
				}
			}
		}
	}

	for i, p1 := range portals {
		for _, p2 := range portals {
			if p1.label == p2.label && p1.from != p2.from {
				p1.to = p2.from
				portals[i] = p1
			}
		}
	}

	return
}

func part1() {
	grid, portals, start, end := getGrid()

	directions := []vector3d{{0, -1, 0}, {1, 0, 0}, {0, 1, 0}, {-1, 0, 0}}
	queue, visited := []state{state{position: start}}, map[vector3d]bool{start: true}
	var st state
	for {
		st, queue = queue[0], queue[1:]
		for _, d := range directions {
			next := vector3d{st.position.x + d.x, st.position.y + d.y, 0}

			if next == end {
				fmt.Println("Part 1:", st.distance+1)
				return
			}

			if grid[next] && !visited[next] {
				visited[next] = true

				p, ok := portals[next]
				if ok {
					next = p.to
				}

				queue = append(queue, state{next, st.distance + 1})
			}
		}
	}
}

func part2() {
	grid, portals, start, end := getGrid()

	directions := []vector3d{{0, -1, 0}, {1, 0, 0}, {0, 1, 0}, {-1, 0, 0}}
	queue, visited := []state{state{position: start}}, map[vector3d]bool{start: true}
	var st state
	for {
		st, queue = queue[0], queue[1:]
		for _, d := range directions {
			next := vector3d{st.position.x + d.x, st.position.y + d.y, st.position.z}

			if next == end {
				fmt.Println("Part 2:", st.distance+1)
				return
			}

			if grid[vector3d{next.x, next.y, 0}] && !visited[next] {
				visited[next] = true

				p, ok := portals[vector3d{next.x, next.y, 0}]
				if ok && (st.position.z > 0 || !p.outer) {
					next = vector3d{p.to.x, p.to.y, st.position.z}
					if p.outer {
						next.z--
					} else {
						next.z++
					}

					visited[next] = true
				}

				queue = append(queue, state{next, st.distance + 1})
			}
		}
	}
}

func main() {
	part1()
	part2()
}
