package main

import (
	"fmt"

	"github.com/kamillo/aoc/2019/intcode"
	"github.com/kamillo/aoc/fileutil"
)

func main() {
	lines := fileutil.GetLines("input.txt")
	ints := intcode.ParseInput(lines[0])
	intCode := intcode.Make(ints)

	scene := [43][49]int{}
	x, y := 0, 0
	for  {
		c, state := intCode.Get()
		fmt.Printf("%c", c)
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
			fmt.Printf("%c", scene[y][x])
			if x != 0 && x != len(scene[y])- 1 && y != 0 && y != len(scene) -1 {
				if scene[y][x] == '#' && scene[y-1][x] == '#' && scene[y+1][x] == '#' && scene[y][x-1] == '#' && scene[y][x+1] == '#'{
					//println("cross")
					aligment += x * y
				}
			}
		}
		println()
	}

	fmt.Println(aligment)
}
