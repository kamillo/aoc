package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	binaryLights := [1000][1000]bool{}
	analogLights := [1000][1000]int{}
	for _, line := range lines {
		upperLeft, bottomRight := utils.Point2D{0, 0}, utils.Point2D{0, 0}
		action := ""

		if _, err := fmt.Sscanf(line, "%s %s %d,%d through %d,%d", &action, &action, &upperLeft.X, &upperLeft.Y, &bottomRight.X, &bottomRight.Y); err != nil {
			fmt.Sscanf(line, "%s %d,%d through %d,%d", &action, &upperLeft.X, &upperLeft.Y, &bottomRight.X, &bottomRight.Y)
		}

		for x := upperLeft.X; x <= bottomRight.X; x++ {
			for y := upperLeft.Y; y <= bottomRight.Y; y++ {
				switch action {
				case "toggle":
					binaryLights[x][y] = !binaryLights[x][y]
					analogLights[x][y] += 2
				case "on":
					binaryLights[x][y] = true
					analogLights[x][y]++
				case "off":
					binaryLights[x][y] = false
					analogLights[x][y]--
					if analogLights[x][y] < 0 {
						analogLights[x][y] = 0
					}
				}
			}
		}
	}

	lit := 0
	brightness := 0
	for x := range binaryLights {
		for y := range binaryLights {
			brightness += analogLights[x][y]
			if binaryLights[x][y] {
				lit++
			}
		}
	}

	fmt.Println(lit)
	fmt.Println(brightness)
}
