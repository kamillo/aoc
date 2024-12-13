package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/kamillo/aoc/utils"
)


func main() {
  lines := utils.GetLines("input.txt")
  // lines := utils.GetLines("test.txt")

  buttonA, buttonB, prize := image.Point{}, image.Point{}, image.Point{}


  part1 := 0
  part2 := 0

  for _, line := range lines {
    if strings.HasPrefix(line, "Button A") {
      fmt.Sscanf(line, "Button A: X+%d, Y+%d", &buttonA.X, &buttonA.Y)

    } else if strings.HasPrefix(line, "Button B") {
      fmt.Sscanf(line, "Button B: X+%d, Y+%d", &buttonB.X, &buttonB.Y)

    } else if strings.HasPrefix(line, "Prize") {
      fmt.Sscanf(line, "Prize: X=%d, Y=%d", &prize.X, &prize.Y)

      if tokens := calcTokens(buttonA, buttonB, prize, 0); tokens != 0 {
        part1 += tokens
      }

      if tokens := calcTokens(buttonA, buttonB, prize, 10000000000000); tokens != 0 {
        part2 += tokens
      }
    }
  }

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}


func calcTokens(A, B, P image.Point, offset int) int {
  prize := image.Pt(P.X + offset, P.Y + offset) 
  det := A.X * B.Y - A.Y * B.X
  a := (prize.X * B.Y - prize.Y * B.X) / det 
  b := (A.X * prize.Y - A.Y * prize.X) / det 

  if A.X * a + B.X * b == prize.X && A.Y * a + B.Y * b == prize.Y {
    return a * 3 + b
  }

  return 0
}
