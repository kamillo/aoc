package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"os"
	"strings"
)

func main() {
	wide := 25
	tall := 6
	lines := utils.GetLines(os.Args[1])
	splitted := strings.Split(lines[0], "")

	layersCount := len(splitted) / (wide * tall)
	layerSize := wide * tall
	layers := make([][]string, layersCount)

	minZeros := wide * tall
	resultPart1 := 0

	image := splitted[:layerSize]

	for i := range layers {
		layers[i] = splitted[i*layerSize : (i+1)*layerSize]
		blackCount, whiteCount, transparentCount := 0, 0, 0

		for j := range layers[i] {
			switch layers[i][j] {
			case "0":
				blackCount++
			case "1":
				whiteCount++
			case "2":
				transparentCount++
			}

			if image[j] == "2" {
				image[j] = layers[i][j]
			}
		}

		if blackCount < minZeros {
			minZeros = blackCount
			resultPart1 = whiteCount * transparentCount
		}
	}

	fmt.Println("Part 1: ", resultPart1)
	fmt.Println("Part 2: ")
	for i := range image {
		if image[i] == "0" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(image[i])
		}
		if (i+1)%wide == 0 {
			fmt.Println()
		}
	}
}
