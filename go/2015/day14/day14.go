package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type reindeer struct {
	speed    int
	flyTime  int
	restTime int
	points   int
	distance int
	seconds  int
	flying   bool
}

func main() {
	time := 2503
	reindeers := make([]reindeer, 0)
	maxDist := 0
	for _, line := range utils.GetLines("input.txt") {
		split := strings.Split(line, " ")
		deer := reindeer{
			speed:    justAtoi(split[3]),
			flyTime:  justAtoi(split[6]),
			restTime: justAtoi(split[13]),
			flying:   true,
		}
		reindeers = append(reindeers, deer)

		times := time / (deer.flyTime + deer.restTime)
		rest := int(math.Min(float64(time-times*(deer.flyTime+deer.restTime)), float64(deer.flyTime)))
		dist := times*deer.speed*deer.flyTime + rest*deer.speed
		if dist > maxDist {
			maxDist = dist
		}
	}

	fmt.Println("Part 1: ", maxDist)

	for i := 0; i < time; i++ {
		max := 0
		for i := range reindeers {
			r := &reindeers[i]
			update(r)

			if r.distance > max {
				max = r.distance
			}
		}
		for i := range reindeers {
			if reindeers[i].distance == max {
				reindeers[i].points++
			}
		}
	}

	s := make([]interface{}, len(reindeers))
	for i, v := range reindeers {
		s[i] = v
	}

	fmt.Println("Part 2: ", utils.MaxInAnyArray(s, func(i1, i2 interface{}) bool {
		return i1.(reindeer).points > i2.(reindeer).points
	}).(reindeer).points)
}

func justAtoi(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func update(r *reindeer) {
	if r.flying {
		if r.seconds+1 <= r.flyTime {
			r.seconds++
			r.distance += r.speed
		} else {
			r.seconds = 1
			r.flying = false
		}
	} else {
		if r.seconds+1 <= r.restTime {
			r.seconds++
		} else {
			r.seconds = 1
			r.flying = true
			r.distance += r.speed
		}
	}
}
