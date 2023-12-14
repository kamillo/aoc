package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	platform := [][]rune{}
	for _, line := range lines {
		platform = append(platform, []rune(line))
	}

	fmt.Println("Part 1:", part1(platform))

	platform = [][]rune{}
	for _, line := range lines {
		platform = append(platform, []rune(line))
	}

	cache := map[string]utils.Point2D{}
	loads := []int{}

	cycleCount := 0
	first := 0
	for i := 0; i < 1000000000; i++ {
		platformStr, load := cycle(platform)
		loads = append(loads, load)

		fmt.Println(load)
		if _, ok := cache[platformStr]; ok {
			first = cache[platformStr].Y
			second := i

			cycleCount = second - first
			break
		}

		cache[platformStr] = utils.NewPoint2D(load, i)
	}

	fmt.Println("Part 2:", loads[first+(1000000000-first-1)%cycleCount])
}

func cycle(platform [][]rune) (string, int) {

	platform = move(platform, moveUp)
	platform = move(platform, moveLeft)

	newPlatform := platform
	for y := len(platform) - 1; y >= 0; y-- {
		for x := 0; x < len(platform[y]); x++ {
			if platform[y][x] == 'O' {
				newPlatform[y][x] = '.'
				x, y = moveDown(platform, x, y)
				newPlatform[y][x] = 'O'
			}
		}
		platform = newPlatform
	}

	newPlatform = platform
	for y := 0; y < len(platform); y++ {
		for x := len(platform[y]) - 1; x >= 0; x-- {
			if platform[y][x] == 'O' {
				newPlatform[y][x] = '.'
				x, y = moveRight(platform, x, y)
				newPlatform[y][x] = 'O'
			}
		}
		platform = newPlatform
	}

	ret := ""
	sum := 0
	for i, line := range platform {
		for _, tile := range line {
			if tile == 'O' {
				sum += len(platform) - i
			}
		}
		ret += string(line)
	}

	return ret, sum

}

func part1(platform [][]rune) int {
	platform = move(platform, moveUp)

	sum := 0
	for i, line := range platform {
		for _, tile := range line {
			if tile == 'O' {
				sum += len(platform) - i
			}
		}
	}

	return sum

}

func move(platform [][]rune, direction func(platform [][]rune, x, y int) (int, int)) [][]rune {
	newPlatform := platform

	for y := 0; y < len(platform); y++ {
		for x := 0; x < len(platform[y]); x++ {
			if platform[y][x] == 'O' {
				newPlatform[y][x] = '.'
				x, y = direction(platform, x, y)
				newPlatform[y][x] = 'O'
			}
		}
		platform = newPlatform
	}

	return platform
}

func moveUp(platform [][]rune, x, y int) (int, int) {
	for y > 0 && platform[y-1][x] == '.' {
		y--
	}
	return x, y
}

func moveDown(platform [][]rune, x, y int) (int, int) {
	for y < len(platform)-1 && platform[y+1][x] == '.' {
		y++
	}
	return x, y
}

func moveLeft(platform [][]rune, x, y int) (int, int) {
	for x > 0 && platform[y][x-1] == '.' {
		x--
	}
	return x, y
}

func moveRight(platform [][]rune, x, y int) (int, int) {
	for x < len(platform[y])-1 && platform[y][x+1] == '.' {
		x++
	}
	return x, y
}
