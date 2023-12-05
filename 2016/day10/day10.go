package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

type Bot struct {
	Id     int
	Values []int

	High map[string]int
	Low  map[string]int
}

func main() {
	lines := utils.GetLines("input.txt")
	bots := map[int]Bot{}
	outputs := map[int][]int{}

	for _, line := range lines {
		v, b := 0, 0

		if _, ok := fmt.Sscanf(line, "value %d goes to bot %d", &v, &b); ok == nil {
			if _, exist := bots[b]; exist {
				bot := bots[b]
				bot.Values = append(bot.Values, v)
				bots[b] = bot
			} else {
				bots[b] = Bot{b, []int{v}, map[string]int{}, map[string]int{}}
			}

		} else {
			l, h := 0, 0
			out1, out2 := "", ""
			if _, ok := fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &b, &out1, &l, &out2, &h); ok == nil {
				if _, exist := bots[b]; !exist {
					bots[b] = Bot{b, []int{}, map[string]int{}, map[string]int{}}
				}

				bot := bots[b]
				bot.Low[out1] = l
				bot.High[out2] = h
				bots[b] = bot
			}
		}
	}

	run := 1
	for run > 0 {
		run = 0
		for i, bbb := range bots {
			if len(bbb.Values) == 2 {
				run++
				high, low := bbb.Values[0], bbb.Values[1]
				if bbb.Values[1] > bbb.Values[0] {
					high = bbb.Values[1]
					low = bbb.Values[0]
				}

				if (high == 61 || low == 61) && (high == 17 || low == 17) {
					fmt.Println("Part 1:", bbb)
				}

				for k, bot := range bbb.High {
					if k == "bot" {
						otherBot := bots[bot]
						otherBot.Values = append(bots[bot].Values, high)
						bots[bot] = otherBot
					} else {
						outputs[bot] = append(outputs[bot], high)
					}
				}

				for k, bot := range bbb.Low {
					if k == "bot" {
						otherBot := bots[bot]
						otherBot.Values = append(bots[bot].Values, low)
						bots[bot] = otherBot

					} else {
						outputs[bot] = append(outputs[bot], low)
					}
				}

				bbb.Values = []int{}
				bots[i] = bbb
			}
		}
	}

	fmt.Println("Part 2:", outputs[0][0]*outputs[1][0]*outputs[2][0])
}
