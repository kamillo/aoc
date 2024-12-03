package main

import (
  "fmt"
  "image"

  "github.com/kamillo/aoc/utils"
)

func main() {
  cross := utils.GetLinesAs2dArray("input.txt")
  // cross := utils.GetLinesAs2dArray("test.txt")

  part1 := 0  
  part2 := 0

  for x := range cross {
    for y := range cross[x] {
      if cross[x][y] == 'X' {
        for _, d := range utils.GetAdj(x, y, cross, 'M', false) {
          if CheckLetter(image.Pt(x + d.X, y + d.Y), d, cross, 'A') {
            if CheckLetter(image.Pt(x + 2 * d.X, y + 2 * d.Y), d, cross, 'S') {
              part1++
            }
          }
        }
      }

      if cross[x][y] == 'A' {
        adjM := utils.GetAdj(x, y, cross, 'M', true)
        adjS := utils.GetAdj(x, y, cross, 'S', true)

        if len(adjM) == 2 && len(adjS) == 2 {
          if adjM[0].X == adjM[1].X || adjM[0].Y == adjM[1].Y {
            part2++
          }
        }
      }

    }

  }

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}

func CheckLetter(p image.Point, d image.Point, grid [][]byte, char byte) bool {
  if p.X + d.X >= 0 && p.Y + d.Y >= 0 && p.Y + d.Y < len(grid[p.X]) && p.X + d.X < len(grid) && grid[p.X + d.X][p.Y + d.Y] == char {
    return true
  }

  return false
}
