package main

import (
	"fmt"

	"github.com/kamillo/aoc/2019/intcode"
	"github.com/kamillo/aoc/fileutil"
)

func main() {
	lines := fileutil.GetLines("input.txt")
	ints := intcode.ParseInput(lines[0])
	computers := [50]intcode.IntCode{}
	for i := range computers {
		computers[i] = intcode.Make(ints)
		computers[i].Put([]int{i})
		computers[i].Put([]int{-1})
	}

	for {
		for i, _ := range computers {
			if (len(computers[i].GetInput()) == 0) {
				computers[i].Put([]int{-1})
			}
			address, state := computers[i].Get()
			if state == 0 {
				X, _ := computers[i].Get()
				Y, _ := computers[i].Get()

				fmt.Println(i, "->", address, X, Y)
				if address == 255 {
					return 
				}
				computers[address].Put([]int{X, Y})
			}
		}
	}
}
