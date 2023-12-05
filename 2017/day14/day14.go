package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	disk := make([][]byte, 128)
	for x := 0; x < 128; x++ {
		// input := "stpzcrnm" + "-" + strconv.Itoa(x)
		input := "flqrgnkx" + "-" + strconv.Itoa(x)
		lengths := []int{}
		for _, l := range input {
			lengths = append(lengths, int(l))
		}

		dense := knotHash(lengths, 64)

		for _, i := range dense {
			s := fmt.Sprintf("%04b%04b", i>>4, i&0x0f)
			sum += strings.Count(s, "1")
			// fmt.Printf("%04b%04b", i>>4, i&0x0f)
			disk[x] = append(disk[x], []byte(s)...)
		}

	}
	fmt.Println(sum)

	reg := 3
	regions := map[int]bool{}
	for y := range disk {
		for x := range disk[y] {
			if disk[y][x] == '1' {
				reg++
				MarkAdj(y, x, &disk, '1', byte(reg))
				regions[reg] = true
			}
		}
	}

	fmt.Println(len(regions))
}

func knotHash(input []int, rounds int) [16]int {
	input = append(input, []int{17, 31, 73, 47, 23}...)

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
		}
	}

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

	return dense
}

func CountAdj(x int, y int, grid [][]byte, char byte) (adj int) {
	if x+1 < len(grid) && grid[x+1][y] == char {
		adj++
	}
	if y+1 < len(grid[x]) && grid[x][y+1] == char {
		adj++
	}
	if x > 0 && grid[x-1][y] == char {
		adj++
	}
	if y > 0 && grid[x][y-1] == char {
		adj++
	}

	return adj
}

func MarkAdj(x int, y int, grid *[][]byte, char byte, mark byte) {
	if x+1 < len(*grid) && (*grid)[x+1][y] == char {
		(*grid)[x+1][y] = mark
		MarkAdj(x+1, y, grid, char, mark)
	}
	if y+1 < len((*grid)[x]) && (*grid)[x][y+1] == char {
		(*grid)[x][y+1] = mark
		MarkAdj(x, y+1, grid, char, mark)
	}
	if x > 0 && (*grid)[x-1][y] == char {
		(*grid)[x-1][y] = mark
		MarkAdj(x-1, y, grid, char, mark)
	}
	if y > 0 && (*grid)[x][y-1] == char {
		(*grid)[x][y-1] = mark
		MarkAdj(x, y-1, grid, char, mark)
	}
}
