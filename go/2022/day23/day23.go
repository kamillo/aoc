package main

import (
	"fmt"
	"image"
	"math"

	"github.com/kamillo/aoc/utils"
)

var offsets = [][]image.Point{
	{image.Pt(0, -1), image.Pt(1, -1), image.Pt(-1, -1)}, // North
	{image.Pt(0, 1), image.Pt(1, 1), image.Pt(-1, 1)},    // South
	{image.Pt(-1, 0), image.Pt(-1, 1), image.Pt(-1, -1)}, // West
	{image.Pt(1, 0), image.Pt(1, -1), image.Pt(1, 1)},    // East
}

type Elves map[image.Point]bool

func main() {
	lines := utils.GetLines("input.txt")

	elfs := Elves{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				elfs[image.Pt(x, y)] = true
			}
		}
	}

	for round := 0; ; round++ {
		possibleMoves := map[image.Point][]image.Point{}
		for e := range elfs {
			if checkAdjenced(e, elfs) {
				if n, ok := nextMove(e, elfs); ok {
					possibleMoves[n] = append(possibleMoves[n], e)
				}
			}
		}

		new := Elves{}
		for k, v := range elfs {
			new[k] = v
		}

		for move, movingElfs := range possibleMoves {
			if len(movingElfs) == 1 {
				delete(new, movingElfs[0])
				new[move] = true
			}
		}

		element := offsets[0]
		offsets = append(offsets[:0], offsets[0+1:]...)
		offsets = append(offsets, element)

		if mapsEquals(new, elfs) {
			fmt.Println("Part 2:", round+1)
			break
		}

		elfs = new
		if round == 9 {
			fmt.Println("Part 1:", getEmptyCount(elfs))
		}
	}
}

func getEmptyCount(e Elves) int {
	maxY, maxX, minY, minX := 0, 0, math.MaxInt, math.MaxInt
	for k := range e {
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
		if k.X < minX {
			minX = k.X
		}
		if k.Y < minY {
			minY = k.Y
		}
	}

	return ((maxX + 1 - minX) * (maxY + 1 - minY)) - len(e)
}

func nextMove(elf image.Point, elfs Elves) (image.Point, bool) {
	for i := 0; i < len(offsets); i++ {
		if elfs.isFreeAt(elf.Add(offsets[i][0])) && elfs.isFreeAt(elf.Add(offsets[i][1])) && elfs.isFreeAt(elf.Add(offsets[i][2])) {
			return elf.Add(offsets[i][0]), true
		}
	}

	return image.Point{}, false
}

func checkAdjenced(elf image.Point, elfs Elves) bool {
	var adjs = []image.Point{
		image.Pt(0, -1),
		image.Pt(0, 1),
		image.Pt(-1, 0),
		image.Pt(1, 0),
		image.Pt(1, 1),
		image.Pt(-1, 1),
		image.Pt(1, -1),
		image.Pt(-1, -1),
	}

	for _, a := range adjs {
		if _, ok := elfs[elf.Add(a)]; ok {
			return true
		}
	}

	return false
}

func (e Elves) isFreeAt(p image.Point) bool {
	return !e[p]
}

func mapsEquals(m, r Elves) bool {
	res := len(m) == len(r)
	if res {
		for k, v := range m {
			if r[k] != v {
				res = false
				break
			}
		}
	}
	return res
}
