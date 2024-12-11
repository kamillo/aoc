package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"

	"github.com/kamillo/aoc/2019/intcode"
)

func main() {
	lines := utils.GetLines("input.txt")
	ints := intcode.ParseInput(lines[0])
	intCode := intcode.Make(ints)

	scene := [43][49]int{}
	x, y := 0, 0
	for {
		c, state := intCode.Get()
		if state == intcode.Exit {
			break
		}
		if c != '\n' {
			scene[y][x] = c
			x++
		} else {
			y++
			x = 0
		}
	}

	//fmt.Println(scene)
	aligment := 0
	for y := range scene {
		for x := range scene[y] {
			if x != 0 && x != len(scene[y])-1 && y != 0 && y != len(scene)-1 {
				if scene[y][x] == '#' && scene[y-1][x] == '#' && scene[y+1][x] == '#' && scene[y][x-1] == '#' && scene[y][x+1] == '#' {
					aligment += x * y
				}
			}
		}
	}

	fmt.Println("Part 1: ", aligment)

	//..............#########..........................
	//..............#.......#..........................
	//..............#.......#..........................
	//..............#.......#..........................
	//..............#.......#..........................
	//..............#.......#..........................
	//..............#.......#..........................
	//..............#.......#..........................
	//..............#...#####.#######.....#############
	//..............#...#.....#.....#.....#...........#
	//..............#...#.....#.....#.....#...........#
	//..............#...#.....#.....#.....#...........#
	//..............#######...#...#############.......#
	//..................#.#...#...#.#.....#...#.......#
	//..................#.#...#############...#.......#
	//..................#.#.......#.#.........#.......#
	//..................#.#.......#.#.........#########
	//..................#.#.......#.#..................
	//..............#############.#.#...#######........
	//..............#...#.#.....#.#.#...#.....#........
	//############^.#...#############...#.....#........
	//#.............#.....#.....#.#.....#.....#........
	//#.............#.....#.....#.#.....#.....#........
	//#.............#.....#.....#.#.....#.....#........
	//#.............#######.....#.#############........
	//#.........................#.......#..............
	//#.....#########...........#.......#..............
	//#.....#.......#...........#.......#..............
	//#.....#.......#...........#.......#..............
	//#.....#.......#...........#.......#..............
	//#.....#.......#############.......#######........
	//#.....#.................................#........
	//#######.................................#........
	//........................................#........
	//........................................#........
	//........................................#........
	//........................................#........
	//........................................#........
	//........................................#........
	//........................................#........
	//........................................#........
	//........................................#........
	//................................#########........
	//L,6,6,L,6,6,L,6,L,6,R,8,R,4,L,6,6,L,6,6,L,6,6,L,6,L,6,L,6,6,L,6,R,6,6,R,8,R,8,R,4,L,6,6,L,6,6,L,6,6,L,6,L,6,L,6,6,L,6,R,6,6,R,8,R,8,R,4,L,6,6,L,6,6,L,6,6,L,6,L,6,L,6,6,L,6,R,6,6,R,8
	//
	//A L,6,6,L,6,6,L,6,L,6
	//B R,8,R,4,L,6,6
	//C L,6,6,L,6,R,6,6,R,8
	//
	//A,B,A,C,B,A,C,B,A,C

	input := []int{'A', ',', 'B', ',', 'A', ',', 'C', ',', 'B', ',', 'A', ',', 'C', ',', 'B', ',', 'A', ',', 'C', '\n'}
	A := []int{'L', ',', '6', ',', '6', ',', 'L', ',', '6', ',', '6', ',', 'L', ',', '6', ',', 'L', ',', '6', '\n'}
	B := []int{'R', ',', '8', ',', 'R', ',', '4', ',', 'L', ',', '6', ',', '6', '\n'}
	C := []int{'L', ',', '6', ',', '6', ',', 'L', ',', '6', ',', 'R', ',', '6', ',', '6', ',', 'R', ',', '8', '\n'}

	ints[0] = 2
	intCode = intcode.Make(ints)
	intCode.Put(input)
	intCode.Put(A)
	intCode.Put(B)
	intCode.Put(C)
	intCode.Put([]int{'n', '\n'})
	for {
		x, _ := intCode.Get()
		if x < 125 {
			fmt.Printf("%c", x)
		} else {
			fmt.Printf("Part 2: %d\n", x)
		}
	}
}