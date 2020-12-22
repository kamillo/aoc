package main

import (
	"fmt"
	"github.com/kamillo/aoc/2019/intcode"
	"github.com/kamillo/aoc/utils"
	"math"
)

func Left(p utils.Point2D) utils.Point2D {
	return utils.NewPoint2D(p.X-1, p.Y)
}
func Right(p utils.Point2D) utils.Point2D {
	return utils.NewPoint2D(p.X+1, p.Y)
}
func Top(p utils.Point2D) utils.Point2D {
	return utils.NewPoint2D(p.X, p.Y+1)
}
func Bottom(p utils.Point2D) utils.Point2D {
	return utils.NewPoint2D(p.X, p.Y-1)
}

func main() {
	lines := utils.GetLines("input.txt")
	ints := intcode.ParseInput(lines[0])
	intCode := intcode.Make(ints)

	area := make(map[utils.Point2D]int)
	distances := make(map[utils.Point2D]float64)

	var res int
	var found utils.Point2D
	dist := 0.0
	dir := 2
	iter := 0
	pos := utils.NewPoint2D(500, 500)

	for {
		intCode.Put([]int{dir})
		res, _ = intCode.Get()
		np := pos
		switch dir {
		case 1:
			np.Y--
		case 2:
			np.Y++
		case 3:
			np.X--
		case 4:
			np.X++
		}

		area[np] = res
		switch res {
		case 0:
			switch dir {
			case 1:
				dir = 4
			case 2:
				dir = 3
			case 3:
				dir = 1
			case 4:
				dir = 2
			}
		case 1:
			distances[pos] = dist
			dist++
			if _, ok := distances[np]; ok {
				dist = math.Min(dist, distances[np])
			}
			pos = np
			switch dir {
			case 4:
				dir = 1
			case 3:
				dir = 2
			case 1:
				dir = 3
			case 2:
				dir = 4
			}
		case 2:
			distances[pos] = dist
			dist++
			if _, ok := distances[np]; ok {
				dist = math.Min(dist, distances[np])
			}
			pos = np
			found = pos
			// break
		}
		if iter > 0 && pos.X == 500 && pos.Y == 500 {
			break
		}
		iter++
	}
	fmt.Println("Part 1: ", distances[found])

	needsAir := true
	steps := 0
	area[found] = 3
	for needsAir {
		needsAir = false
		areaTemp := make(map[utils.Point2D]int)
		for k, v := range area {
			if v == 1 {
				needsAir = true
				neighbors := []utils.Point2D{Top(k), Bottom(k), Left(k), Right(k)}
				for _, n := range neighbors {
					if area[n] == 3 {
						v = 3
					}
				}
			}
			areaTemp[k] = v
		}
		steps++
		area = areaTemp
	}
	fmt.Println("Part 2: ", steps)
}
