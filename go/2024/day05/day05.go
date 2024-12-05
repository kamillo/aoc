package main

import (
	"fmt"
	"sort"

	"github.com/kamillo/aoc/utils"
)


type set map[int]bool
type Order []int

var rules = map[int]set{}

func (a Order) Len() int           { return len(a) }
func (a Order) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Order) Less(i, j int) bool { return rules[a[i]][a[j]] }

func main() {
  lines := utils.GetLines("input.txt")
  // lines := utils.GetLines("test.txt")

  rules = map[int]set{}
  part1 := 0  
  part2 := 0

  pages := false

  for _, line := range lines {
    if !pages {
      l, r := 0, 0

      if c, ok := fmt.Sscanf(line, "%d|%d", &l, &r); c != 2 || ok != nil {
        pages = true
      } 

      if ok := rules[l]; ok == nil {
        rules[l] = set{}
      }

      rules[l][r] = true

    } else {
      order := Order(utils.ToIntArr(line, ","))

      if sort.IsSorted(order) {
        part1 += order[len(order)/2]

      } else {
        sort.Sort(order)
        part2 += order[len(order)/2]
      }
    }
  }

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}



