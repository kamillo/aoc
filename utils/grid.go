package utils

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
