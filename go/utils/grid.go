package utils

import (
	"image"
)

func CountAdj(x int, y int, grid [][]byte, char byte) (adj int) {
	if x+1 < len(grid) && grid[x+1][y] == char {
		adj++
	}
	if y+1 < len(grid[x]) && grid[x][y+1] == char {
		adj++
	}
	if x+1 < len(grid) && y+1 < len(grid[x]) && grid[x+1][y+1] == char {
		adj++
	}
	if x > 0 && grid[x-1][y] == char {
		adj++
	}
	if x > 0 && y+1 < len(grid[x]) && grid[x-1][y+1] == char {
		adj++
	}
	if y > 0 && grid[x][y-1] == char {
		adj++
	}
	if y > 0 && x+1 < len(grid) && grid[x+1][y-1] == char {
		adj++
	}
	if x > 0 && y > 0 && grid[x-1][y-1] == char {
		adj++
	}

	return adj
}


func CountPerpendicularAdj(x int, y int, grid [][]byte, char byte) (adj int) {
	if x+1 < len(grid) && grid[x+1][y] == char {
		adj++
	}
	if y+1 < len(grid[x]) && grid[x][y+1] == char {
		adj++
	}
	if x > 0 && grid[x-1][y] == char {
		adj++
	}
	if y > 0 && grid[x][y-1] == char {
		adj++
	}
	return adj
}


func GetAdj(x int, y int, grid [][]byte, char byte, diagOnly bool) []image.Point {
	adj := []image.Point{}
	if !diagOnly {
		if x+1 < len(grid) && grid[x+1][y] == char {
			adj = append(adj, image.Pt(1, 0))
		}
		if y+1 < len(grid[x]) && grid[x][y+1] == char {
			adj = append(adj, image.Pt(0, 1))
		}
		if x > 0 && grid[x-1][y] == char {
			adj = append(adj, image.Pt(-1, 0))
		}
		if y > 0 && grid[x][y-1] == char {
			adj = append(adj, image.Pt(0, -1))
		}
	}
	if x+1 < len(grid) && y+1 < len(grid[x]) && grid[x+1][y+1] == char {
		adj = append(adj, image.Pt(1, 1))
	}
	if x > 0 && y+1 < len(grid[x]) && grid[x-1][y+1] == char {
		adj = append(adj, image.Pt(-1, 1))
	}
	if y > 0 && x+1 < len(grid) && grid[x+1][y-1] == char {
		adj = append(adj, image.Pt(1, -1))
	}
	if x > 0 && y > 0 && grid[x-1][y-1] == char {
		adj = append(adj, image.Pt(-1, -1))
	}

	return adj
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
