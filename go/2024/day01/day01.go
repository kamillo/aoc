package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	
  part1 := 0
  part2 := 0

  digits1 := []int{}
  digits2 := []int{}
  counts := map[int]int{}

  for _, line := range lines {
    x, y := 0,0;
    
    fmt.Sscanf(string(line), "%d %d", &x, &y)
		digits1 = append(digits1, x)
    digits2 = append(digits2, y)

    counts[y] ++
	}

  sort.Ints(digits1)
  sort.Ints(digits2)

  for i := range digits1 {
    part1 += int(math.Abs(float64(digits1[i]) - float64(digits2[i])))
    part2 += digits1[i] * counts[digits1[i]] 
  }

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}
