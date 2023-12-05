package main

import (
	"container/ring"
	"fmt"
	"github.com/kamillo/aoc/utils"
)

func main() {
	//input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	input := []int{2, 1, 9, 3, 4, 7, 8, 6, 5}

	part1 := game(input, 100)
	part1.Do(func(p interface{}) {
		fmt.Print(p.(int))
	})
	fmt.Println()

	m := utils.MaxInArray(input) + 1
	for i := len(input); i < 1000000; i++ {
		input = append(input, m)
		m++
	}
	part2 := game(input, 10000000)
	fmt.Println(part2.Next().Value.(int) * part2.Next().Next().Value.(int))
}

func game(cupsLabels []int, moves int) *ring.Ring {
	length := len(cupsLabels)
	cups := ring.New(length)
	cupsMap := map[int]*ring.Ring{}
	min := utils.MinInArray(cupsLabels)

	for i := 0; i < length; i++ {
		cups.Value = cupsLabels[i]
		cupsMap[cups.Value.(int)] = cups
		cups = cups.Next()
	}

	for round := 0; round < moves; round++ {
		three := cups.Unlink(3)
		dest := cups.Value.(int) - 1

		for {
			if dest < min {
				dest = length
			}
			inRemoved := false
			three.Do(func(p interface{}) {
				if p.(int) == dest {
					inRemoved = true
					dest--
				}
			})
			if !inRemoved {
				break
			}
		}

		cupsMap[dest].Link(three)
		cups = cups.Next()
	}

	return cupsMap[1]
}
