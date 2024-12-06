package main

import (
  "fmt"
  "image"

  "github.com/kamillo/aoc/utils"
)

type direction image.Point

var (
  up    = image.Pt(0, -1)
  down  = image.Pt(0, 1)
  left  = image.Pt(-1, 0)
  right = image.Pt(1, 0)
)

var lab [][]byte
func main() {
  lab = utils.GetLinesAs2dArray("input.txt")
  // lab = utils.GetLinesAs2dArray("test.txt")

  part1 := 0  
  part2 := 0

  var guard image.Point 
  for y := range lab {
    for x := range lab[y] {
      if lab[y][x] == '^' {
        guard = image.Pt(x, y)
        break
      }
    }
  } 

  visited := map[image.Point]direction{}
  visited[guard] = direction(up)

  _, visited = walk(guard)

  for k, _ := range visited {
    lab[k.Y][k.X] = '#'

    if cycle, _ := walk(guard); cycle {
      part2++
    }

    lab[k.Y][k.X] = '.'
  }

  part1 = len(visited)

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}

func turnRight(d image.Point) image.Point {
  switch d {
  case up:
    return right
  case down:
    return left
  case left:
    return up
  case right:
    return down
  }

  return d
}

func walk(guard image.Point) (bool, map[image.Point]direction) {
  visited := map[image.Point]direction{}
  dir := up
  visited[guard] = direction(dir)
  next := guard.Add(dir)
  cycle := false

  for ; next.X >= 0 && next.X < len(lab[0]) && next.Y >= 0 && next.Y < len(lab); {
    for {
      if d, ok := visited[next]; ok && d == direction(dir) {
        cycle = true
        break
      } 

      if lab[next.Y][next.X] != '#' {
        guard = next

        visited[guard] = direction(dir)
        break
      } 

      dir = turnRight(dir)
      next = guard.Add(dir)
    }

    if cycle {
      break
    }

    next = guard.Add(dir)
  }

  return cycle, visited
}
