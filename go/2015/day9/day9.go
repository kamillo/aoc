package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"math"
)

func main() {
	locations := map[string]bool{}
	distances := map[string]int{}
	for _, line := range utils.GetLines("input.txt") {
		loc1, loc2, dist := "", "", 0
		fmt.Sscanf(line, "%s to %s = %d", &loc1, &loc2, &dist)
		locations[loc1] = true
		locations[loc2] = true
		distances[loc1+loc2] = dist
		distances[loc2+loc1] = dist
	}

	locs := make([]interface{}, 0, len(locations))
	for k := range locations {
		locs = append(locs, k)
	}
	fmt.Println(len(locs))
	//locs := []interface{} {"a", "b", "c", "aa", "bb", "cc", "zz", "xx"}
	perm := utils.HeapPermutation(locs)
	minDist := math.MaxInt64
	maxDist := 0
	for _, p := range perm {
		distSum := 0
		for i := range p {
			if i+1 >= len(p) {
				break
			}
			if dist, ok := distances[p[i].(string)+p[i+1].(string)]; ok {
				distSum += dist
			}
		}
		if distSum < minDist {
			minDist = distSum
		}
		if distSum > maxDist {
			maxDist = distSum
		}
	}
	fmt.Println("Part 1: ", minDist)
	fmt.Println("Part 2: ", maxDist)
}
