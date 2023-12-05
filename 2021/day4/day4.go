package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Number struct {
	v       int
	checked bool
}

type Board struct {
	grid [5][5]Number
	win  bool
}

func main() {
	boards := []Board{}
	input := utils.GetLines("input.txt")

	numbers := utils.StringsToInts(strings.Split(input[0], ","))
	input = input[2:]

	board := Board{}
	y := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		for x, v := range strings.Fields(line) {
			if i, err := strconv.Atoi(v); err == nil {
				board.grid[x][y] = Number{i, false}
			}
		}

		y++

		if y%5 == 0 {
			y = 0
			boards = append(boards, board)
			board = Board{}
		}
	}

	check := func() int {
		winners := 0
		for _, n := range numbers {
			for b := range boards {
				if boards[b].win {
					continue
				}
				for y := range boards[b].grid {
					for x := range boards[b].grid[y] {
						if boards[b].grid[x][y].v == n {
							boards[b].grid[x][y].checked = true

							if checkBoard(boards[b], x, y) {
								boards[b].win = true
								winners++

								if winners == 0 {
									fmt.Println("Part 1: ", calculateResult(boards[b])*n)
								}

								if winners == len(boards)-1 {
									fmt.Println("Part 2: ", calculateResult(boards[b])*n)
								}
							}
						}
					}
				}
			}
		}

		return 0
	}

	check()
}

func checkBoard(board Board, x int, y int) (bingo bool) {
	bingo = true
	for yy := range board.grid {
		bingo = bingo && board.grid[x][yy].checked
	}

	if !bingo {
		bingo = true
		for xx := range board.grid[y] {
			bingo = bingo && board.grid[xx][y].checked
		}
	}

	return
}

func calculateResult(board Board) (sum int) {
	for y := range board.grid {
		for x := range board.grid[y] {
			if !board.grid[x][y].checked {
				sum += board.grid[x][y].v
			}
		}
	}
	return
}
