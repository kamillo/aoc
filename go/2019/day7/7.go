package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/2019/intcode"
)

func heapPermutation(a []int) [][]int {
	var permutations [][]int
	var generate func([]int, int)

	generate = func(a []int, size int) {
		if size == 1 {
			A := make([]int, len(a))
			copy(A, a)
			permutations = append(permutations, A)
		}
		for i := 0; i < size; i++ {
			generate(a, size-1)
			if size%2 == 1 {
				a[0], a[size-1] = a[size-1], a[0]
			} else {
				a[i], a[size-1] = a[size-1], a[i]
			}
		}
	}
	generate(a, len(a))
	return permutations
}

func main() {
	lines := utils.GetLines("input.txt")
	//lines = []string{"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"}
	//lines = []string{"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"}
	//lines = []string{"3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"}

	splitted := strings.Split(lines[0], ",")
	ints := make([]int, len(splitted))
	for i := 0; i < len(splitted); i++ {
		ints[i], _ = strconv.Atoi(splitted[i])
	}

	// Part 1
	{
		intCode := intcode.Make(ints)
		a := []int{0, 1, 2, 3, 4}
		permutations := heapPermutation(a)
		max := 0
		for i := range permutations {
			res := 0
			for j := range permutations[i] {
				intCode.Put([]int{permutations[i][j], res})
				res, _ = intCode.Get()
				intCode = intcode.Make(ints)
			}
			if res > max {
				max = res
			}
		}
		fmt.Println("Part 1: ", max)
	}

	// Part 2
	{
		a := []int{5, 6, 7, 8, 9}
		permutations := heapPermutation(a)
		max := 0
		for i := range permutations {
			modules := [...]intcode.IntCode{intcode.Make(ints), intcode.Make(ints), intcode.Make(ints), intcode.Make(ints), intcode.Make(ints)}
			res := 0
			maxE := 0
			for j := range permutations[i] {
				modules[j].Put([]int{permutations[i][j], res})
				res, _ = modules[j].Get()
			}

			for {
				status := intcode.Exit
				for j := range modules {
					modules[j].Put([]int{res})
					res, status = modules[j].Get()
				}
				if res > maxE {
					maxE = res
				}
				if status == intcode.Exit {
					break
				}
			}

			if maxE > max {
				max = maxE
			}
		}
		fmt.Println("Part 2: ", max)
	}
}
