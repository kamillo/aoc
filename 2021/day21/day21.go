package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	// Player 1 starting position: 4
	// Player 2 starting position: 10
	p1 := 4
	p2 := 10

	fmt.Println("Part 1: ", part1(p1, p2))
	fmt.Println("Part 2: ", part2(p1, p2))
}

func part1(p1, p2 int) int {
	score1 := 0
	score2 := 0
	rolls := 0

	for i := 1; score1 < 1000 && score2 < 1000; {
		p1 = (p1 + nextThree(&i)) % 10
		if p1 == 0 {
			p1 = 10
		}
		score1 += p1

		rolls += 3
		if score1 >= 1000 {
			break
		}

		p2 = (p2 + nextThree(&i)) % 10
		if p2 == 0 {
			p2 = 10
		}
		score2 += p2
		rolls += 3
	}

	return utils.Min(score1, score2) * rolls
}

func nextThree(x *int) (sum int) {
	for i := 0; i < 3; i++ {
		sum += *x
		*x++
		if *x > 100 {
			*x = *x - 100
		}
	}
	return
}

type state struct {
	p1, p2         int
	score1, score2 int
}

func advance(pos int, roll int) int {
	pos += roll
	for pos > 10 {
		pos -= 10
	}
	return pos
}

var states = make(map[state][]int)

func part2(p1, p2 int) int {
	rolls := []int{}
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				rolls = append(rolls, r1+r2+r3)
			}
		}
	}

	s := state{p1: p1, p2: p2, score1: 0, score2: 0}
	var run func(s state) []int
	run = func(s state) []int {
		if s.score1 >= 21 {
			return []int{1, 0}
		}
		if s.score2 >= 21 {
			return []int{0, 1}
		}
		if v, ok := states[s]; ok {
			return v
		}
		win := []int{0, 0}
		for _, roll := range rolls {
			p1 := advance(s.p1, roll)
			s1 := s.score1 + p1

			w := run(state{
				p1:     s.p2,
				score1: s.score2,
				p2:     p1,
				score2: s1,
			})
			win[0] += w[1]
			win[1] += w[0]
		}
		states[s] = win
		return win
	}

	scores := run(s)
	return utils.Max(scores[0], scores[1])
}
