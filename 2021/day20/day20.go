package main

import (
	"fmt"
	"strconv"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")
	algorithm := lines[0]
	image := lines[2:]

	enchanced := enchance(image, algorithm)
	enchanced = enchance(enchanced, algorithm)
	count := 0
	for _, line := range enchanced {
		fmt.Println(line)
		for _, c := range line {
			if c == '#' {
				count++
			}
		}
	}

	fmt.Println(count)
}

func enchance(input []string, algorithm string) (image []string) {
	for y := -1; y < len(input)+1; y++ {
		line := ""
		for x := -1; x < len(input[0])+1; x++ {
			pixels := getPixels(x, y, input)
			line += string(algorithm[toDec(pixels)])
		}

		image = append(image, line)
	}

	return
}

func getPixels(targetX int, targetY int, image []string) (line string) {
	for y := targetY - 1; y <= targetY+1; y++ {
		for x := targetX - 1; x <= targetX+1; x++ {
			if y < 0 || x < 0 || y >= len(image) || x >= len(image[0]) {
				line += "0"
			} else {
				if image[y][x] == '.' {
					line += "0"
				} else {
					line += "1"
				}
			}
		}
	}

	return
}

func toDec(bin string) int {
	v, _ := strconv.ParseInt(bin, 2, 64)
	return int(v)
}
