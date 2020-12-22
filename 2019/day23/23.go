package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"

	"github.com/kamillo/aoc/2019/intcode"
)

func main() {
	lines := utils.GetLines("input.txt")
	ints := intcode.ParseInput(lines[0])
	computers := [50]intcode.IntCode{}
	for i := range computers {
		computers[i] = intcode.Make(ints)
		computers[i].Put([]int{i})
		computers[i].Put([]int{-1})
	}

	NAT := [2]int{}

	for {
		idles := 0
		for i, _ := range computers {
			if len(computers[i].GetInput()) == 0 {
				idles++
				computers[i].Put([]int{-1})
			}
			address, state := computers[i].Get()
			if state == 0 {
				X, _ := computers[i].Get()
				Y, _ := computers[i].Get()

				// fmt.Println(i, "->", address, X, Y)
				if address == 255 {
					fmt.Println("Setting NAT: ", Y)
					NAT[0] = X
					NAT[1] = Y
				} else {
					computers[address].Put([]int{X, Y})
				}
			}
		}
		if idles == len(computers) {
			fmt.Println("Sending to 0: ", NAT[0], NAT[1])
			computers[0].Put([]int{NAT[0], NAT[1]})
		}
	}
}
