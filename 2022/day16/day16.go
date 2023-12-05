package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Valve struct {
	tunels []string
	flow   int
}

func main() {
	lines := utils.GetLines("input.txt")

	valves := map[string]Valve{}

	dist := map[string]map[string]int{}

	getDist := func(i, j string) int {
		if v, ok := dist[i][j]; ok {
			return v
		} else {
			return 99999999
		}
	}

	for _, line := range lines {
		name := ""
		rate := 0
		t := ""
		var tunnels []string

		if _, err := fmt.Sscanf(line, "Valve %s has flow rate=%d; tunnels lead to valves %s", &name, &rate, &t); err == nil {
			tunnels = strings.Split(line, ", ")
			tunnels[0] = tunnels[0][len(tunnels[0])-2:]
		} else if _, err := fmt.Sscanf(line, "Valve %s has flow rate=%d; tunnel leads to valve %s", &name, &rate, &t); err == nil {
			tunnels = []string{t}
		} else {
			panic("scanf")
		}

		valves[name] = Valve{tunels: tunnels, flow: rate}
		dist[name] = map[string]int{}
		dist[name][name] = 0

		for _, t := range tunnels {
			dist[name][t] = 1
		}

		for k := range valves {
			for i := range valves {
				for j := range valves {
					dist[i][j] = int(math.Min(float64(getDist(i, j)), float64(getDist(i, k)+getDist(k, j))))
				}
			}
		}
	}

	var openValve func(i string, t int, remaining map[string]bool) int
	openValve = memoize(func(i string, t int, remaining map[string]bool) int {
		res := 0
		for j := range remaining {
			next_t := t - getDist(i, j) - 1
			if next_t >= 0 {
				nextRemaining := make(map[string]bool, len(remaining))
				for k, v := range remaining {
					nextRemaining[k] = v
				}
				delete(nextRemaining, j)
				res = utils.MaxInArray([]int{res, valves[j].flow*next_t + openValve(j, next_t, nextRemaining)})
			}
		}

		return res
	})

	remaining := map[string]bool{}

	for n, v := range valves {
		if v.flow > 0 {
			remaining[n] = true
		}
	}

	fmt.Println("Part 1:", openValve("AA", 30, remaining))

	memoizeCache = make(map[string]int)
	var openValve2 func(i string, t int, remaining map[string]bool, elephant bool) int
	openValve2 = memoize2(func(i string, t int, remaining map[string]bool, elephant bool) int {
		res := 0
		if elephant {
			res = openValve2("AA", 26, remaining, false)
		}

		for j := range remaining {
			next_t := t - getDist(i, j) - 1
			if next_t >= 0 {
				nextRemaining := make(map[string]bool, len(remaining))
				for k, v := range remaining {
					nextRemaining[k] = v
				}
				delete(nextRemaining, j)
				res = utils.MaxInArray([]int{res, valves[j].flow*next_t + openValve2(j, next_t, nextRemaining, elephant)})
			}
		}

		return res
	})

	fmt.Println("Part 2:", openValve2("AA", 26, remaining, true))

}

var memoizeCache = make(map[string]int)

func memoize(f func(i string, t int, remaining map[string]bool) int) func(i string, t int, remaining map[string]bool) int {
	return func(i string, t int, remaining map[string]bool) int {
		key := getHash(i, t, remaining)
		if v, ok := memoizeCache[key]; ok {
			return v
		}

		result := f(i, t, remaining)
		memoizeCache[key] = result

		return result
	}
}

func memoize2(f func(i string, t int, remaining map[string]bool, elephant bool) int) func(i string, t int, remaining map[string]bool, elephant bool) int {
	return func(i string, t int, remaining map[string]bool, elephant bool) int {
		key := getHash2(i, t, remaining, elephant)
		if v, ok := memoizeCache[key]; ok {
			return v
		}

		result := f(i, t, remaining, elephant)
		memoizeCache[key] = result

		return result
	}
}

func getHash(i string, t int, m map[string]bool) string {
	return fmt.Sprintf("%s-%d-%v", i, t, m)
}

func getHash2(i string, t int, m map[string]bool, elephant bool) string {
	return fmt.Sprintf("%s-%d-%v-%v", i, t, m, elephant)
}
