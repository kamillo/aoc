package main

import (
  "fmt"
  "image"

  "github.com/kamillo/aoc/utils"
)


func main() {
  grid := utils.GetLinesAs2dArray("input.txt")
  // grid := utils.GetLinesAs2dArray("test.txt")

  part1 := 0 
  part2 := 0

  graph := utils.Graph[Node]{}
  heads := []Node{}
  tails := []Node{}
  for y := range grid {
    for x := range grid[y] {

      node := Node(image.Pt(x, y))

      if grid[y][x] == '0' {
        heads = append(heads, node)
      }

      if grid[y][x] == '9' {
        tails = append(tails, node)
      }

      for _, p := range utils.GetPerpendicularAdj(x, y, grid, func(char byte) bool { 
        return char - grid[y][x] == 1 
      }) {
        graph[node] = append(graph[node], Node(p))
      }
    }
  }

  for _, head := range heads {
    paths := utils.BFSPathFinder[Node](graph, head, func(n Node) bool { return utils.Contains(tails, n) })
    paths2 := utils.DFSPathFinder[Node](graph, head, func(n Node) bool { return utils.Contains(tails, n) })
    part1 += len(paths)
    part2 += len(paths2)
  }

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}

type Node image.Point


