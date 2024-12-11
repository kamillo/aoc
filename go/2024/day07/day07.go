package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type operation string
const (
  mul = "*"
  add = "+"
)

func main() {
  lines := utils.GetLines("input.txt")
  // lines := utils.GetLines("test.txt")

  part1 := 0  
  part2 := 0

  // variationsCache := map[int][]string{}
  equations := []int{}
  for _, line := range lines {
    split := strings.Split(line, ": ")
    k := utils.JustAtoi(split[0])
    equations = utils.ToIntArr(split[1], " ")

    part1 += calc(k, equations, false)
    part2 += calc(k, equations, true)
  }

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}

func calc(k int, equations []int, part2 bool) int {
  ret := 0
  variationsCache := map[int][]string{}
    opLen := len(equations) - 1
    operations := []string{"+", "*"}
  if part2 {
    operations = append(operations, "|")
  }
    if _, ok := variationsCache[opLen]; !ok {
      variationsCache[opLen] = utils.Variations(operations, opLen)
    }

    for _, variation := range variationsCache[opLen] {
      res := equations[0]
      for i := 1; i < len(equations); i++ {
        if variation[i - 1] == '*' {
          res *= equations[i]
        } else if variation[i - 1] == '+'{
          res += equations[i]
        } else if part2 {
          res = utils.JustAtoi(fmt.Sprintf("%d%d", res, equations[i]))
        }
      }

      if res == k {
        // fmt.Println(v)
        ret += k
        break
      }
    }

  return ret
}
