package main

import (
  "fmt"
  "image"
  "github.com/kamillo/aoc/utils"
)

func main() {
  city := utils.GetLinesAs2dArray("input.txt")
  // city := utils.GetLinesAs2dArray("test.txt")

  part1 := 0  
  // part2 := 0

  antenas := map[byte][]image.Point{}

  for y := range city {
    for x, b := range city[y] {
      if b != '.' {
        antenas[b] = append(antenas[b], image.Pt(x, y))
      }
    }
  }

  antinodes := map[image.Point]bool{}
  for _, points := range antenas {
    for _, p := range points {
      for _, q := range points {
        if p != q {
          a, b := FindAntinodes(p, q)
 
          if CheckBounds(len(city[0]), len(city), a) && city[a.Y][a.X] != '#' { 
            part1++
            city[a.Y][a.X] = '#'
          }

          if CheckBounds(len(city[0]), len(city), b) && city[b.Y][b.X] != '#' {
            part1++
            city[b.Y][b.X] = '#'
          }
        }
      }
    }
  }

  antinodes = map[image.Point]bool{}
  for _, points := range antenas {
    for _, p := range points {
      for _, q := range points {
        if p != q {
          for _, a := range FindAllAntinodes(len(city[0]), len(city), p , q) {
            antinodes[a] = true
          }
        }
      }
    }
  }

  for ant := range antinodes {
    city[ant.X][ant.Y] = '#'
  } 

  for _, y := range city {
    for _, x := range y {
      fmt.Printf("%c", x)
    }
    fmt.Println()
  }


  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", len(antinodes))
}

func CheckBounds(maxX, maxY int, p image.Point) bool {
  return p.X >= 0 && p.X < maxX && p.Y >= 0 && p.Y < maxY
}

func FindAntinodes(p1, p2 image.Point) (image.Point, image.Point) {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	p3 := image.Point{X: p1.X - dx, Y: p1.Y - dy} 
	p4 := image.Point{X: p2.X + dx, Y: p2.Y + dy}

	return p3, p4
}

func FindAllAntinodes(maxX, maxY int, p1, p2 image.Point) ([]image.Point) {
  dx := p2.X - p1.X
	dy := p2.Y - p1.Y

  ret := []image.Point{}
  pa := p1 

  for ; CheckBounds(maxX, maxY, pa);  { 
    ret = append(ret, pa)
    pa = image.Point{X: pa.X - dx, Y: pa.Y - dy} 
  } 

  pa = p1 
  for ; CheckBounds(maxX, maxY, pa); {
    ret = append(ret, pa)
    pa = image.Point{X: pa.X + dx, Y: pa.Y + dy} 
  }

  return ret
}

