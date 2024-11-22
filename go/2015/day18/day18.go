package main

import (
	"bytes"
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strings"
)

func main() {
	lights := utils.GetLinesAs2dArray("input.txt")
	run(lights, false)
	fmt.Println("Part 1: ", strings.Count(string(bytes.Join(lights, []byte{})), "#"))

	lights = utils.GetLinesAs2dArray("input.txt")
	run(lights, true)
	fmt.Println("Part 2: ", strings.Count(string(bytes.Join(lights, []byte{})), "#"))
}

func run(lights [][]byte, part2 bool) {
	if part2 {
		lights[0][len(lights[0])-1] = '#'
		lights[0][0] = '#'
		lights[len(lights)-1][0] = '#'
		lights[len(lights)-1][len(lights[0])-1] = '#'
	}
	for i := 0; i < 100; i++ {
		newLights := make([][]byte, len(lights))
		copy(newLights, lights)
		for light := range lights {
			newLights[light] = make([]byte, len(lights[0]))
			copy(newLights[light], lights[light])
		}

		for x := range newLights {
			for y := range newLights[x] {
				switch newLights[x][y] {
				case '#':
					adj := utils.CountAdj(x, y, newLights, '#')
					if adj != 3 && adj != 2 {
						lights[x][y] = '.'
					}
				case '.':
					if utils.CountAdj(x, y, newLights, '#') == 3 {
						lights[x][y] = '#'
					}
				}
			}
		}

		if part2 {
			lights[0][len(lights[0])-1] = '#'
			lights[0][0] = '#'
			lights[len(lights)-1][0] = '#'
			lights[len(lights)-1][len(lights[0])-1] = '#'
		}
	}
}
