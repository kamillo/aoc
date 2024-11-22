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

	enchanced := image
	for i := 0; i < 2; i++ {
		enchanced = enchance(enchanced, algorithm)
	}
	fmt.Println("Part 1:", count(enchanced))

	enchanced = image
	for i := 0; i < 50; i++ {
		enchanced = enchance(enchanced, algorithm)
	}
	fmt.Println("Part 2:", count(enchanced))
}

func count(image []string) (count int) {
	for _, line := range image {
		//fmt.Println(line)
		for _, c := range line {
			if c == '#' {
				count++
			}
		}
	}

	return
}

var outsideVal = "0"

func enchance(input []string, algorithm string) (image []string) {
	for y := -1; y < len(input)+1; y++ {
		line := ""
		for x := -1; x < len(input[0])+1; x++ {
			pixels := getPixels(x, y, input)
			line += string(algorithm[toDec(pixels)])
		}

		image = append(image, line)
	}

	if outsideVal == "0" {
		outsideVal = toString(algorithm[0])
	} else {
		outsideVal = toString(algorithm[len(algorithm)-1])
	}

	return
}

func toString(c byte) string {
	if c == '#' {
		return "1"
	}

	return "0"
}

func getPixels(targetX int, targetY int, image []string) (line string) {

	for y := targetY - 1; y < targetY+2; y++ {
		for x := targetX - 1; x < targetX+2; x++ {
			if y < 0 || x < 0 || y >= len(image) || x >= len(image[0]) {
				line += outsideVal
			} else {
				line += toString(image[y][x])
			}
		}
	}

	return
}

func toDec(bin string) int {
	v, _ := strconv.ParseInt(bin, 2, 64)
	return int(v)
}
