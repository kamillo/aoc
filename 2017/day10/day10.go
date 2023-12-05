package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	{
		input := strings.Split(utils.GetLines("input.txt")[0], ",")
		lengths := []int{}
		for _, l := range input {
			length, _ := strconv.Atoi(l)
			lengths = append(lengths, length)
		}

		f := knotHash(lengths, 1)
		fmt.Println("Part 1: ", f.Value.(int)*f.Next().Value.(int))
	}

	{
		input := utils.GetLines("input.txt")[0]
		lengths := []int{}
		for _, l := range input {
			lengths = append(lengths, int(l))
		}
		lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

		f := knotHash(lengths, 64)
		dense := [16]int{}
		x := 0
		for i := 0; i < f.Len(); i += 16 {
			for j := 0; j < 16; j++ {
				if j == 0 {
					dense[x] = f.Value.(int)
				} else {
					dense[x] = dense[x] ^ f.Value.(int)
				}
				f = f.Next()
			}
			x++
		}

		fmt.Println("Part 2:")
		for i := 0; i < 16; i++ {
			fmt.Printf("%x", dense[i])
		}
	}

}

func knotHash(input []int, rounds int) *ring.Ring {
	r := ring.New(256)
	f := r
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	skip := 0

	for round := 0; round < rounds; round++ {
		for _, length := range input {
			if length > 0 {
				tmp := ring.New(length)
				for i := 0; i < length; i++ {
					tmp.Value = r.Value
					r = r.Next()
					tmp = tmp.Next()
				}

				r = r.Move(-length)

				tmp = tmp.Prev()
				for i := 0; i < length; i++ {
					r.Value = tmp.Value
					r = r.Next()
					tmp = tmp.Prev()
				}
			}
			r = r.Move(skip)
			skip++

			// f.Do(func(p any) {
			// 	fmt.Print(p.(int))
			// 	fmt.Print(" ")
			// })
			// fmt.Println("", r.Value.(int))
		}
	}

	return f
}
