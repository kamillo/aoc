package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

var offsets = []image.Point{
	image.Pt(1, 0),  // East
	image.Pt(0, 1),  // South
	image.Pt(-1, 0), // West
	image.Pt(0, -1), // North
}

func main() {
	lines := utils.GetLines("input.txt")

	fields := lines[:len(lines)-2]
	path := lines[len(lines)-1:][0]
	path = strings.ReplaceAll(path, "R", ";R;")
	path = strings.ReplaceAll(path, "L", ";L;")

	fmt.Println(path)
	for _, f := range fields {
		fmt.Println(f)
	}

	d := 0
	current := image.Pt(0, 0)
	for i := 0; i < len(fields[0]); i++ {
		if fields[0][i] != ' ' {
			current = image.Pt(i, 0)
			break
		}
	}

	for _, l := range strings.Split(path, ";") {
		if n, err := strconv.Atoi(l); err == nil {
			current = nextPoint(fields, current, offsets[d], n)

		} else {
			if l == "L" {
				d = utils.ModWrap((d - 1), 4)
			} else {
				d = utils.ModWrap((d + 1), 4)
			}
		}
	}

	fmt.Println(1000*(current.Y+1) + 4*(current.X+1) + d)

	d = 0
	for i := 0; i < len(fields[0]); i++ {
		if fields[0][i] != ' ' {
			current = image.Pt(i, 0)
			break
		}
	}

	for _, l := range strings.Split(path, ";") {
		if n, err := strconv.Atoi(l); err == nil {
			for i := 0; i < n; i++ {
				current, d = nextPoint2(fields, current, d)
			}

		} else {
			if l == "L" {
				d = utils.ModWrap((d - 1), 4)
			} else {
				d = utils.ModWrap((d + 1), 4)
			}
		}
	}
	fmt.Println(1000*(current.Y+1) + 4*(current.X+1) + d)

}

func nextPoint(field []string, point, offset image.Point, n int) image.Point {
	ret := point
	tmp := point

	for i := 0; i < n; i++ {
		tmp = tmp.Add(offset)
		tmp.Y = utils.ModWrap(tmp.Y, len(field))
		tmp.X = utils.ModWrap(tmp.X, len(field[tmp.Y]))

		for field[tmp.Y][tmp.X] == ' ' {
			tmp = tmp.Add(offset)
			tmp.Y = utils.ModWrap(tmp.Y, len(field))
			tmp.X = utils.ModWrap(tmp.X, len(field[tmp.Y]))
		}

		if field[tmp.Y][tmp.X] == '#' {
			return ret
		}

		ret = tmp
	}

	return ret
}

func nextPoint2(field []string, point image.Point, d int) (image.Point, int) {
	tmp := point.Add(offsets[d])
	nd := d

	if nd == 0 && inRange(tmp.Y, 0, 50) && tmp.X == 150 {
		nd, tmp.Y, tmp.X = 2, 149-tmp.Y, 99
	} else if nd == 0 && inRange(tmp.Y, 50, 100) && tmp.X == 100 {
		nd, tmp.Y, tmp.X = 3, 49, 50+tmp.Y
	} else if nd == 0 && inRange(tmp.Y, 100, 150) && tmp.X == 100 {
		nd, tmp.Y, tmp.X = 2, 149-tmp.Y, 149
	} else if nd == 0 && inRange(tmp.Y, 150, 200) && tmp.X == 50 {
		nd, tmp.Y, tmp.X = 3, 149, tmp.Y-100

	} else if nd == 1 && inRange(tmp.X, 0, 50) && tmp.Y == 200 {
		nd, tmp.Y, tmp.X = 1, 0, tmp.X+100
	} else if nd == 1 && inRange(tmp.X, 50, 100) && tmp.Y == 150 {
		nd, tmp.Y, tmp.X = 2, tmp.X+100, 49
	} else if nd == 1 && inRange(tmp.X, 100, 150) && tmp.Y == 50 {
		nd, tmp.Y, tmp.X = 2, tmp.X-50, 99

	} else if nd == 2 && inRange(tmp.Y, 0, 50) && tmp.X == 49 {
		nd, tmp.Y, tmp.X = 0, 149-tmp.Y, 0
	} else if nd == 2 && inRange(tmp.Y, 50, 100) && tmp.X == 49 {
		nd, tmp.Y, tmp.X = 1, 100, tmp.Y-50
	} else if nd == 2 && inRange(tmp.Y, 100, 150) && tmp.X == -1 {
		nd, tmp.Y, tmp.X = 0, 149-tmp.Y, 50
	} else if nd == 2 && inRange(tmp.Y, 150, 200) && tmp.X == -1 {
		nd, tmp.Y, tmp.X = 1, 0, tmp.Y-100

	} else if nd == 3 && inRange(tmp.X, 0, 50) && tmp.Y == 99 {
		nd, tmp.Y, tmp.X = 0, 50+tmp.X, 50
	} else if nd == 3 && inRange(tmp.X, 50, 100) && tmp.Y == -1 {
		nd, tmp.Y, tmp.X = 0, tmp.X+100, 0
	} else if nd == 3 && inRange(tmp.X, 100, 150) && tmp.Y == -1 {
		nd, tmp.Y, tmp.X = 3, 199, tmp.X-100
	}

	if field[tmp.Y][tmp.X] == '.' {
		return tmp, nd
	} else if field[tmp.Y][tmp.X] == '#' {
		return point, d
	}

	panic("should not happen")
}

func inRange(i, start, stop int) bool {
	return i >= start && i < stop
}
