package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"strconv"
	"strings"
)

func main() {
	lines := fileutil.GetLines("input.txt")

	timestamp, _ := strconv.Atoi(lines[0])
	busses := strings.Split(lines[1], ",")
	minDelay, minId := 0, 0
	for _, id := range busses {
		if id == "x" {
			continue
		}

		id, _ := strconv.Atoi(id)
		depart := id
		for depart <= timestamp {
			depart += id
		}
		if minDelay == 0 || minDelay > depart-timestamp {
			minDelay = depart - timestamp
			minId = id
		}
	}
	fmt.Println("Part 1: ", minDelay*minId)

	timestamp = 0
	step := 1
	i := 0
	for i < len(busses) {
		if busses[i] == "x" {
			i++
		} else {
			id, _ := strconv.Atoi(busses[i])
			if (timestamp+i)%id == 0 {
				step *= id
				i++
			} else {
				timestamp += step
			}
		}
	}
	fmt.Println("Part 2: ", timestamp)
}
