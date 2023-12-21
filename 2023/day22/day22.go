package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type BrickLine struct {
	start  utils.PointD3D
	end    utils.PointD3D
	pilars []int
}

type FloorBrick struct {
	z     int
	index int
}

func main() {
	lines := utils.GetLines("input.txt")
	bricks := []BrickLine{}

	maxX, maxY, maxZ := 0, 0, 0

	for _, line := range lines {
		start := utils.ToIntArr(strings.Split(line, "~")[0], ",")
		end := utils.ToIntArr(strings.Split(line, "~")[1], ",")

		startPoint := utils.PointD3D{start[0], start[1], start[2]}
		endPoint := utils.PointD3D{end[0], end[1], end[2]}

		bricks = append(bricks, BrickLine{startPoint, endPoint, []int{}})

		maxX = utils.Max(maxX, startPoint.X, endPoint.X)
		maxY = utils.Max(maxY, startPoint.Y, endPoint.Y)
		maxZ = utils.Max(maxZ, startPoint.Z, endPoint.Z)
	}

	ground := make([][]FloorBrick, maxY+1)
	for i := range ground {
		ground[i] = make([]FloorBrick, maxX+1)
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].start.Z < bricks[j].start.Z
	})

	bricks, _ = fallDown(bricks, ground)

	// for b := range bricks {
	// 	fmt.Println(bricks[b].start, bricks[b].end)
	// }

	count := 0
	sum := 0
	for i := range bricks {
		ground := make([][]FloorBrick, maxY+1)
		for i := range ground {
			ground[i] = make([]FloorBrick, maxX+1)
		}

		newBricks := make([]BrickLine, len(bricks))
		copy(newBricks, bricks)
		newBricks = append(newBricks[:i], newBricks[i+1:]...)

		_, fallen := fallDown(newBricks, ground)
		sum += fallen
		if fallen == 0 {
			count++
		}
	}

	fmt.Println("Part 1:", count)
	fmt.Println("Part 2:", sum)
}

func fallDown(bricks []BrickLine, ground [][]FloorBrick) ([]BrickLine, int) {
	count := 0
	for i, brick := range bricks {
		minZ1, minZ2 := 0, 0
		p1 := bricks[i].start
		p2 := bricks[i].end

		if p1.X == p2.X && p1.Z == p2.Z {
			for y := p1.Y; y <= p2.Y; y++ {
				if ground[p1.X][y].z > minZ1 {
					minZ1 = ground[p1.X][y].z
				}
			}

			for y := p1.Y; y <= p2.Y; y++ {
				if ground[p1.X][y].z == minZ1 {
					bricks[i].pilars = append(bricks[i].pilars, ground[p1.X][y].index)
				}
			}

			minZ1, minZ2 = minZ1+1, minZ1+1

			for y := p1.Y; y <= p2.Y; y++ {
				ground[p1.X][y].z = minZ1
				ground[p1.X][y].index = i
			}

		} else if p1.Y == p2.Y && p1.Z == p2.Z {
			for x := p1.X; x <= p2.X; x++ {
				if ground[x][p1.Y].z > minZ1 {
					minZ1 = ground[x][p1.Y].z
				}
			}

			for x := p1.X; x <= p2.X; x++ {
				if ground[x][p1.Y].z == minZ1 {
					bricks[i].pilars = append(bricks[i].pilars, ground[x][p1.Y].index)
				}
			}

			minZ1, minZ2 = minZ1+1, minZ1+1

			for x := p1.X; x <= p2.X; x++ {
				ground[x][p1.Y].z = minZ1
				ground[x][p1.Y].index = i
			}

			// if x and y coordinates are same, block elongates in z dir
		} else if p1.X == p2.X && p1.Y == p2.Y {
			bricks[i].pilars = append(bricks[i].pilars, ground[brick.start.X][brick.start.Y].index)
			minZ1 = ground[p1.X][p1.Y].z + 1
			minZ2 = minZ1 + p2.Z - p1.Z
			ground[p1.X][p1.Y].z = minZ2
			ground[p1.X][p1.Y].index = i
		}

		if bricks[i].start.Z != minZ1 || bricks[i].end.Z != minZ2 {
			count++
			bricks[i].start.Z = minZ1
			bricks[i].end.Z = minZ2
		}
	}

	return bricks, count
}

func printBricks(bricks []BrickLine, maxX, maxY, maxZ int) {
	fmt.Println("X")
	for z := maxZ; z > 0; z-- {
		for x := 0; x <= maxX; x++ {
			br := false
			for i, brick := range bricks {
				if brick.start.Z <= z && brick.end.Z >= z && brick.start.X <= x && brick.end.X >= x {
					fmt.Print(i)
					br = true
				}
			}
			if !br {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println("Y")
	for z := maxZ; z > 0; z-- {
		for y := 0; y <= maxY; y++ {
			br := false
			for i, brick := range bricks {
				if brick.start.Z <= z && brick.end.Z >= z && brick.start.Y <= y && brick.end.Y >= y {
					fmt.Print(i)
					br = true
				}
			}
			if !br {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
