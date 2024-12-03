package main

import (
  "fmt"
  "strings"

  "github.com/kamillo/aoc/utils"
)

func main() {
  lines := utils.GetLines("input.txt")
  // lines := utils.GetLines("test2.txt")

  part1 := 0  
  part2 := 0

  do := true
  for _, line := range lines {
    index := 0

    for {
      i := strings.Index(line[index:], "mul(")
     
      if i < 0 {
        break
      }

      index += i + 3
      
      ido := strings.LastIndex(line[:index], "don't()")
      if ido > -1 {
        do = false
      }

      if i = strings.LastIndex(line[:index], "do()"); i > -1 && i > ido {
        do = true
      }

      x, y := 0, 0
      
      if c, ok := fmt.Sscanf(line[index:], "(%d,%d)", &x, &y); ok == nil && c == 2 {
        part1 += x * y

        if do {
          part2 += x * y
        }
      }
    }
  }

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}
