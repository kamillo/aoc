package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/kamillo/aoc/fileutil"
)

func check(p1 Point, p2 Point, p3 Point) bool {
	dxc := p1.x - p2.x
	dyc := p1.y - p2.y

	dxl := p3.x - p2.x
	dyl := p3.y - p2.y

	cross := dxc*dyl - dyc*dxl

	if cross == 0 {
		if math.Abs(float64(dxl)) >= math.Abs(float64(dyl)) {
			if dxl > 0 {
				return p2.x <= p1.x && p1.x <= p3.x
			} else {
				return p3.x <= p1.x && p1.x <= p2.x
			}
		} else {
			if dyl > 0 {
				return p2.y <= p1.y && p1.y <= p3.y
			} else {
				return p3.y <= p1.y && p1.y <= p2.y
			}
		}
	}

	return false
}

type Point struct {
	x int
	y int
}

func main() {
	lines := fileutil.GetLines("input.txt")
	//lines = []string{
	//	"......#.#." ,
	//	"#..#.#....",
	//	"..#######.",
	//	".#.#.###..",
	//	".#..#.....",
	//	"..#....#.#",
	//	"#..#....#.",
	//	".##.#..###",
	//	"##...#..#.",
	//	".#....####",
	//}
	//lines = []string {
	//	".#....#####...#..",
	//	"##...##.#####..##",
	//	"##...#...#.#####.",
	//	"..#.....X...###..",
	//	"..#.#.....#....##",
	//}
	//lines = []string {
	//	".#..##.###...#######",
	//	"##.############..##.",
	//	".#.######.########.#",
	//	".###.#######.####.#.",
	//	"#####.##.#.##.###.##",
	//	"..#####..#.#########",
	//	"####################",
	//	"#.####....###.#.#.##",
	//	"##.#################",
	//	"#####.##.###..####..",
	//	"..######..##.#######",
	//	"####.##.####...##..#",
	//	".#####..#.######.###",
	//	"##...#.##########...",
	//	"#.##########.#######",
	//	".####.#.###.###.#.##",
	//	"....##.##.###..#####",
	//	".#.#.###########.###",
	//	"#.#.#.#####.####.###",
	//	"###.##.####.##.#..##",
	//}

	pointsMap := make(map[Point][]Point)
	points := []Point{}

	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			if c == "#" {
				points = append(points, Point{x, y})
			}
		}
	}

	max := Point{}
	for _, p1 := range points {
		for _, p2 := range points {
			if p1 == p2 {
				continue
			}
			found := false
			for _, p3 := range points {
				if p1 == p3 || p2 == p3 {
					continue
				}
				if check(p3, p1, p2) {
					found = true
					break
				}
			}
			if !found {
				pointsMap[p1] = append(pointsMap[p1], p2)
			}
		}

		if len(pointsMap[p1]) > len(pointsMap[max]) {
			max = p1
		}
	}

	fmt.Println("Part 1: ", max, len(pointsMap[max]))

	type kv struct {
		point Point
		angle float64
	}

	var closePoints []kv
	for _, k := range pointsMap[max] {
		angle := math.Atan2(float64(k.y-max.y), float64(k.x-max.x)) - math.Atan2(float64(0-max.y), float64(max.x-max.x))

		angle = angle * (180.0 / math.Pi)
		if angle < 0.0 {
			angle = 360.0 + angle
		}
		closePoints = append(closePoints, kv{k, angle})
	}

	sort.Slice(closePoints, func(i, j int) bool {
		return closePoints[i].angle < closePoints[j].angle
	})

	fmt.Println("Part 2: ", closePoints[199], closePoints[199].point.x*100+closePoints[199].point.y)
}
