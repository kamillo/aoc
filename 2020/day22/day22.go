package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"github.com/kamillo/aoc/utils"
	"strconv"
)

func main() {
	lines := fileutil.GetLines("input.txt")
	player1, player2 := parseDecks(lines)

	winner := game(player1, player2)
	p1 := 0
	for i := range winner {
		p1 += winner[i] * (len(winner) - i)
	}
	fmt.Println("Part 1:", p1)

	_, winner2 := gameOfRecursion(player1, player2)
	p2 := 0
	for i := range winner2 {
		p2 += winner2[i] * (len(winner2) - i)
	}
	fmt.Println("Part 2:", p2)
}

func game(player1 []int, player2 []int) []int {
	for len(player1) != 0 && len(player2) != 0 {
		if player1[0] > player2[0] {
			player1 = append(player1, player1[0], player2[0])
		} else {
			player2 = append(player2, player2[0], player1[0])
		}
		player2 = player2[1:]
		player1 = player1[1:]
	}
	if len(player1) > 0 {
		return player1
	} else {
		return player2
	}
}

func gameOfRecursion(p1 []int, p2 []int) (bool, []int) {
	player1 := make([]int, len(p1))
	player2 := make([]int, len(p2))
	copy(player1, p1)
	copy(player2, p2)
	games := map[string]bool{}

	for len(player1) != 0 && len(player2) != 0 {
		game := utils.SliceToString(player1, ",") + utils.SliceToString(player2, ",")
		if games[game] {
			return true, player1
		}
		games[game] = true

		if player1[0] <= len(player1)-1 && player2[0] <= len(player2)-1 {
			if p1wins, _ := gameOfRecursion(player1[1:player1[0]+1], player2[1:player2[0]+1]); p1wins {
				player1 = append(player1, player1[0], player2[0])
			} else {
				player2 = append(player2, player2[0], player1[0])
			}
		} else {
			if player1[0] > player2[0] {
				player1 = append(player1, player1[0], player2[0])
			} else {
				player2 = append(player2, player2[0], player1[0])
			}
		}
		player2 = player2[1:]
		player1 = player1[1:]
	}

	if len(player1) > len(player2) {
		return true, player1
	} else {
		return false, player2
	}
}

func parseDecks(lines []string) ([]int, []int) {
	player1 := make([]int, 0, len(lines)-3)
	player2 := make([]int, 0, len(lines)-3)

	first := true
	for _, line := range lines {
		if len(line) == 0 {
			first = false
		}
		if c, err := strconv.Atoi(line); err == nil {
			if first {
				player1 = append(player1, c)
			} else {
				player2 = append(player2, c)
			}
		}
	}

	return player1, player2
}
