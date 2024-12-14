package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/kamillo/aoc/utils"
)

type Region struct {
  perimeter int
  points map[image.Point]bool
}

func main() {
  garden := utils.GetLinesAs2dArray("input.txt")
  // garden := utils.GetLinesAs2dArray("test.txt")

  part1 := 0  
  part2 := 0

  regions := map[byte][]Region{}

  for y := range garden {
    for x := range garden {
      s := garden[y][x]
      point := image.Pt(x, y)
      if _, ok := regions[s]; !ok {
        regions[s] = []Region{}
        regions[s] = append(regions[s], Region{0,  map[image.Point]bool{point: true}})

        regions[s][0].points = findAdjs(point, garden, func(char byte) bool { return char == s }, regions[s][0].points)
        for p := range regions[s][0].points {
          regions[s][0].perimeter += calcPerimeter(p.X, p.Y, garden, s)
        }

      } else {

        exist := false 

        for _, r := range regions[s] {
          if ok := r.points[point]; ok {
            exist = true
            break
          }
        }

        if !exist {
          regions[s] = append(regions[s], Region{0, map[image.Point]bool{image.Pt(x, y): true}})

          last := len(regions[s]) - 1

          // fmt.Printf("adding2 %c %v\n", s, regions[s])
          regions[s][last].points = findAdjs(point, garden, func(char byte) bool { return char == s }, regions[s][last].points)
          for p := range regions[s][last].points {
            regions[s][last].perimeter += calcPerimeter(p.X, p.Y, garden, s)
          }
        }
      }
    } 
  }


  for _, v := range regions {
    for _, r := range v {
      // fmt.Printf("%c: %d * %d\n", k, len(r.points), r.perimeter)
      part1 += len(r.points) * r.perimeter
      part2 += len(r.points) * len(getEdges(garden, r.points))
      // fmt.Printf("- %c %d * %d %v\n", k , len(r.points), len(getEdges(garden, r.points)), getEdges(garden, r.points))
    }

  }

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)

}

func calcPerimeter(x, y int, grid [][]byte, char byte) int {
  count := utils.CountPerpendicularAdj(y, x, grid, char)

  return 4 - count
}


func findAdjs(point image.Point, grid [][]byte, condition func(char byte) bool, adj map[image.Point]bool) map[image.Point]bool {
  adjs := utils.GetPerpendicularAdj(point.X, point.Y, grid, condition)  
  for _, a := range adjs {
    if ok := adj[a]; !ok {
      adj[a] = true 
      adj = findAdjs(a, grid, condition, adj)
    }
  }

  return adj

}


func getEdges(garden [][]byte, region  map[image.Point]bool) []int {
  edges := []int{}

  for y := range garden {
    groups := []int{}

    for x := range garden[y] {
      if !region[image.Pt(x, y)] {
        continue
      }

      if y == 0 || garden[y - 1][x] != garden[y][x] {
        if len(groups) > 0 && garden[y][x-1] == garden[y][x] && (x == 0 || y == 0 || (x > 0 && y > 0 && garden[y-1][x-1] != garden[y][x])) {
          groups[len(groups) - 1]++
        } else {
          groups = append(groups, 1)
        }
      }
    }

    edges = append(edges, groups...)
   }

  for y := range slices.Backward(garden) {
    groups := []int{}

    for x := range garden[y] {
      if !region[image.Pt(x, y)] {
        continue
      }

      if y + 1 >= len(garden) || garden[y+1][x] != garden[y][x] {
        if len(groups) > 0 && garden[y][x-1] == garden[y][x] && ( x == 0 || y + 1 == len(garden) || (x > 0 && y + 1 < len(garden) && garden[y+1][x-1] != garden[y][x])) {
          groups[len(groups) - 1]++
        } else {
          groups = append(groups, 1)
        }
      }
    }

    edges = append(edges, groups...)
  }

  for x := range garden[0] {
    groups := []int{}

    for y := range garden {
      if !region[image.Pt(x, y)] {
        continue
      }

      if x == 0 || garden[y][x-1] != garden[y][x] {
        if len(groups) > 0 && garden[y-1][x] == garden[y][x] && (x == 0 || y == 0 || (x > 0 && y > 0 && garden[y-1][x-1] != garden[y][x])) {
          groups[len(groups)-1]++

        } else {
          groups = append(groups, 1)
        }
      }
    }

    edges = append(edges, groups...)
  }

  for x := range slices.Backward(garden[0]) {
    groups := []int{}

    for y := range garden {
      if !region[image.Pt(x, y)] {
        continue
      }

      if x + 1 >= len(garden[0]) || garden[y][x+1] != garden[y][x] {
        if len(groups) > 0 && garden[y-1][x] == garden[y][x] && (x + 1 ==  len(garden[0]) || y == 0 || (x + 1 < len(garden[0]) && y > 0 && garden[y-1][x+1] != garden[y][x])) {
          groups[len(groups)-1]++

        } else {
          groups = append(groups, 1)
        }
      }
    }

    edges = append(edges, groups...)
  }



  return edges
}


func GetPerpendicularAdj(x int, y int, grid [][]byte, condition func(char byte) bool) []image.Point {
	adj := []image.Point{}
  if y+1 < len(grid) && condition(grid[y+1][x]) {
    adj = append(adj, image.Pt(x, y + 1))
  }
  if x+1 < len(grid[x]) && condition(grid[y][x+1]) {
    adj = append(adj, image.Pt(x + 1, y))
  }
  if y > 0 && condition(grid[y-1][x]) {
    adj = append(adj, image.Pt(x, y - 1))
  }
  if x > 0 && condition(grid[y][x-1]) {
    adj = append(adj, image.Pt(x - 1, y))
  }
  return adj
}
