package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"math"
)

func main() {
	lines := utils.GetLines("input.txt")

	// Part 1
	ship := utils.NewPoint2D(0, 0)
	directionsMap := map[rune]int{'E': 0, 'S': 1, 'W': 2, 'N': 3}
	directionsArr := []rune{'E', 'S', 'W', 'N'}
	heading := 'E'

	var move func(utils.Point2D, rune, int) utils.Point2D
	move = func(ship utils.Point2D, dir rune, value int) utils.Point2D {
		switch dir {
		case 'N':
			ship.Y -= value
		case 'S':
			ship.Y += value
		case 'W':
			ship.X -= value
		case 'E':
			ship.X += value
		case 'F':
			ship = move(ship, heading, value)
		case 'L':
			dirNum := value / 90
			index := modWrap(directionsMap[heading]-dirNum, len(directionsArr))
			heading = directionsArr[index]
		case 'R':
			dirNum := value / 90
			index := modWrap(directionsMap[heading]+dirNum, len(directionsArr))
			heading = directionsArr[index]
		}
		return ship
	}

	for _, line := range lines {
		var dir rune
		value := 0
		fmt.Sscanf(line, "%c%d", &dir, &value)

		ship = move(ship, dir, value)
	}
	fmt.Println(math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y)))

	// Part 2
	ship.X, ship.Y = 0, 0
	waypoint := utils.NewPoint2D(10, -1)
	heading = 'E'

	for _, line := range lines {
		var dir rune
		num := 0
		fmt.Sscanf(line, "%c%d", &dir, &num)

		switch dir {
		case 'N':
			waypoint.Y -= num
		case 'S':
			waypoint.Y += num
		case 'W':
			waypoint.X -= num
		case 'E':
			waypoint.X += num
		case 'F':
			ship.X += num * waypoint.X
			ship.Y += num * waypoint.Y
		case 'R':
			angle := num / 90
			for ; angle > 0; angle-- {
				waypoint.X, waypoint.Y = -waypoint.Y, waypoint.X
			}
		case 'L':
			angle := num / 90
			for ; angle > 0; angle-- {
				waypoint.X, waypoint.Y = waypoint.Y, -waypoint.X
			}
		}
	}
	fmt.Println(math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y)))
}

func modWrap(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
