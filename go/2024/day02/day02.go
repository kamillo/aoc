package main

import (
  "fmt"
  "math"

  "github.com/kamillo/aoc/utils"
)

func main() {
  lines := utils.GetLines("input.txt")

  part1 := 0  
  part2 := 0

  for _, s := range lines {
    line := utils.ToIntArr(s, " ")
    ok := false

    ok = checkLevels(line)
    if ok {
      part1++
      part2++

    } else {

      for i := 0; i < len(line); i++ {
        newLine := []int{}
        newLine = append(newLine, line...)

        ok = checkLevels(utils.DeleteAtIndex(newLine, i))

        if ok {
          part2++
          break;
        }
      }
    }
  }

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}

func isSafe(a, b int, inc bool) bool {
  abs := int(math.Abs(float64(a) - float64(b))) 
  return (inc == (a < b)) && abs > 0 && abs <=3 
}

func checkLevels(line []int) bool {
  ok := false
  inc := line[0] < line[1]

  for i := 0; i < len(line)-1; i++ {
    ok = isSafe(line[i], line[i+1], inc)

    if !ok {
      break
    }
  }

  return ok
}
