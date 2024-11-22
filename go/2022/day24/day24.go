package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Blizzard struct {
	pos, dir, wrap image.Point
}

type Grid [][]rune

func main() {
	lines := utils.GetLines("input.txt")
	fmt.Println("Part 1:", p1(lines))
	fmt.Println("Part 2:", p2(lines))
}

var directions = map[string]image.Point{
	"^": image.Pt(0, -1), // North
	"v": image.Pt(0, 1),  // South
	"<": image.Pt(-1, 0), // West
	">": image.Pt(1, 0),  // East
}

func (grid *Grid) parse(lines []string) (blizzards []Blizzard) {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		*grid = append(*grid, []rune(line))
	}
	for y, row := range *grid {
		for x, c := range row {
			dir := directions[string(c)]
			switch c {
			case '^':
				blizzards = append(blizzards, Blizzard{
					image.Pt(x, y), dir, image.Pt(x, len(*grid)-2),
				})
			case 'v':
				blizzards = append(blizzards, Blizzard{
					image.Pt(x, y), dir, image.Pt(x, 1),
				})
			case '<':
				blizzards = append(blizzards, Blizzard{
					image.Pt(x, y), dir, image.Pt(len((*grid)[0])-2, y),
				})
			case '>':
				blizzards = append(blizzards, Blizzard{
					image.Pt(x, y), dir, image.Pt(1, y),
				})
			}
		}
	}
	return
}

func (grid Grid) traverseField(blizzards []Blizzard, start, target image.Point) int {
	currentPath := map[image.Point]bool{}
	currentPath[start] = true

	time := 0
	for ; !currentPath[target]; time++ {
		moveBlizzards := map[image.Point]bool{}
		for i, b := range blizzards {
			newB := b.pos.Add(b.dir)
			if grid.isInBounds(newB) {
				if grid[newB.Y][newB.X] == '#' {
					blizzards[i].pos = b.wrap
				} else {
					blizzards[i].pos = newB
				}
			}
			moveBlizzards[blizzards[i].pos] = true
		}

		newStep := make(map[image.Point]bool)
		for pos := range currentPath {
			if !(moveBlizzards[pos]) {
				newStep[pos] = true
			}
			for _, d := range directions {
				new := pos.Add(d)
				if grid.isFreeAt(new) && !moveBlizzards[new] {
					newStep[new] = true
				}
			}
		}
		currentPath = newStep
	}

	return time
}

func (g Grid) isFreeAt(p image.Point) bool {
	return g.isInBounds(p) && g[p.Y][p.X] != '#'
}

func (g Grid) isInBounds(pos image.Point) bool {
	return pos.X >= 0 && pos.X < len(g[0]) && pos.Y >= 0 && pos.Y < len(g)
}

func p1(lines []string) int {
	start := image.Pt(1, 0)
	target := image.Pt(len(lines[0])-2, len(lines)-1)

	grid := Grid{}
	blizzards := grid.parse(lines)
	return grid.traverseField(blizzards, start, target)
}

func p2(lines []string) int {
	start := image.Pt(1, 0)
	target := image.Pt(len(lines[0])-2, len(lines)-1)

	grid := Grid{}
	blizzards := grid.parse(lines)
	return grid.traverseField(blizzards, start, target) + grid.traverseField(blizzards, target, start) + grid.traverseField(blizzards, start, target)
}
