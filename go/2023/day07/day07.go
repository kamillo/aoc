package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Cards map[rune]int

type Hand struct {
	cards string
	bid   int
	rank  int
}

func main() {
	lines := utils.GetLines("input.txt")

	cards1 := utils.MakeIndexed("23456789TJQKA")
	cards2 := utils.MakeIndexed("J23456789TQKA")

	fmt.Println("Part 1:", totalWinnings(lines, cards1, getRank1))
	fmt.Println("Part 2:", totalWinnings(lines, cards2, getRank2))
}

func totalWinnings(lines []string, cards Cards, getRank func(s string) int) int {
	hands := make([]Hand, len(lines))

	for i, line := range lines {
		cards := strings.Split(line, " ")[0]
		bid := utils.JustAtoi(strings.Split(line, " ")[1])
		rank := getRank(cards)
		hands[i] = Hand{cards, bid, rank}
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rank == hands[j].rank {
			for x := 0; x < len(hands[i].cards); x++ {
				if cards[rune(hands[i].cards[x])] != cards[rune(hands[j].cards[x])] {
					return cards[rune(hands[i].cards[x])] < cards[rune(hands[j].cards[x])]
				}
			}
		}
		return hands[i].rank < hands[j].rank
	})

	sum := 0
	for i, h := range hands {
		sum += h.bid * (i + 1)
	}

	return sum
}

func getRank1(s string) int {
	max := 0
	unique := make(map[rune]int, len(s))
	for _, r := range s {
		unique[r]++
		if unique[r] > max {
			max = unique[r]
		}
	}

	switch max {
	case 1:
		return 0
	case 2:
		if len(unique) == 4 {
			return 1
		}
		return 2
	case 3:
		if len(unique) == 3 {
			return 3
		}
		return 4
	case 4:
		return 5
	default:
		return 6
	}
}

func getRank2(s string) int {
	max := 0
	unique := make(map[rune]int, len(s))
	for _, r := range s {
		unique[r]++
		if unique[r] > max && r != 'J' {
			max = unique[r]
		}
	}

	max += unique['J']
	delete(unique, 'J')

	switch max {
	case 1:
		return 0
	case 2:
		if len(unique) == 4 {
			return 1
		}
		return 2
	case 3:
		if len(unique) == 3 {
			return 3
		}
		return 4
	case 4:
		return 5
	default:
		return 6
	}
}
