package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"math"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	startX float64
	startY float64
	endX   float64
	endY   float64
	length float64
}

func lineLineIntersection(line1 Line, line2 Line) (x float64, y float64, ok bool) {
	s1X := line1.endX - line1.startX
	s1Y := line1.endY - line1.startY
	s2X := line2.endX - line2.startX
	s2Y := line2.endY - line2.startY

	s, t := 0.0, 0.0
	if (-s2X*s1Y + s1X*s2Y) != 0 {
		s = (-s1Y*(line1.startX-line2.startX) + s1X*(line1.startY-line2.startY)) / (-s2X*s1Y + s1X*s2Y)
		t = (s2X*(line1.startY-line2.startY) - s2Y*(line1.startX-line2.startX)) / (-s2X*s1Y + s1X*s2Y)
	} else {
		return 0, 0, false
	}

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		// Collision detected
		x = math.Round(line1.startX + (t * s1X))
		y = math.Round(line1.startY + (t * s1Y))
		return x, y, true
	}

	return 0, 0, false // No collision
}

func main() {
	lines := utils.GetLines(os.Args[1])
	//lines := []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}
	//lines := []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	//lines := []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	cables := make([][]Line, len(lines))

	for i, line := range lines {
		paths := strings.Split(line, ",")
		cables[i] = make([]Line, len(paths))

		x, y := 0.0, 0.0
		for j, path := range paths {
			var line Line
			value, _ := strconv.ParseFloat(path[1:], 64)

			switch path[0] {
			case 'R':
				line = Line{x, y, x + value, y, value}
			case 'L':
				line = Line{x, y, x - value, y, value}
			case 'U':
				line = Line{x, y, x, y + value, value}
			case 'D':
				line = Line{x, y, x, y - value, value}
			}

			x, y = line.endX, line.endY

			cables[i][j] = line
		}
		//fmt.Println()
	}

	minDist := 0.0
	minPath := 0.0
	for i := range cables[0] {
		for j := range cables[1] {
			x, y, ok := lineLineIntersection(cables[0][i], cables[1][j])
			if ok {
				path := 0.0
				dist := math.Abs(x) + math.Abs(y)

				if dist < minDist || minDist == 0 {
					minDist = dist
				}

				for _, cable := range cables[0][:i] {
					path += cable.length
				}
				path += math.Abs(x-cables[0][i].startX) + math.Abs(y-cables[0][i].startY)

				for _, cable := range cables[1][:j] {
					path += cable.length
				}
				path += math.Abs(x-cables[1][j].startX) + math.Abs(y-cables[1][j].startY)

				if path < minPath || minPath == 0 {
					minPath = path
				}
				//fmt.Println("path: ", path)
				//fmt.Println(x, y, dist)
			}
		}
	}

	fmt.Println("Part 1: ", minDist)
	fmt.Println("Part 2: ", minPath)
}
