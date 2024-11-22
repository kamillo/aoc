package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Hailstone struct {
	px, py, pz, vx, vy, vz int
}

func main() {
	lines := utils.GetLines("input.txt")
	low := float64(200000000000000)
	max := float64(400000000000000)
	// lines := utils.GetLines("test.txt")
	// low := float64(7)
	// max := float64(27)

	vxx, vyy, vzz := make(map[int][]int), make(map[int][]int), make(map[int][]int)
	hailStoneA, hailStoneB := Hailstone{}, Hailstone{}

	count := 0
	for i := 0; i < len(lines); i++ {
		point := utils.ToIntArr(strings.Split(lines[i], " @ ")[0], ", ")
		velocity := utils.ToIntArr(strings.Split(lines[i], " @ ")[1], ", ")

		s1, i1 := calculateSlopeAndIntercept(point[0], point[1], velocity[0], velocity[1])

		if _, ok := vxx[velocity[0]]; !ok {
			vxx[velocity[0]] = make([]int, 0)
		}
		if _, ok := vyy[velocity[1]]; !ok {
			vyy[velocity[1]] = make([]int, 0)
		}
		if _, hz := vzz[velocity[2]]; !hz {
			vzz[velocity[2]] = make([]int, 0)
		}
		vxx[velocity[0]] = append(vxx[velocity[0]], point[0])
		vyy[velocity[1]] = append(vyy[velocity[1]], point[1])
		vzz[velocity[2]] = append(vzz[velocity[2]], point[2])

		for j := i + 1; j < len(lines); j++ {
			pointB := utils.ToIntArr(strings.Split(lines[j], " @ ")[0], ", ")
			velocityB := utils.ToIntArr(strings.Split(lines[j], " @ ")[1], ", ")

			if i == 0 {
				hailStoneA = Hailstone{point[0], point[1], point[2], velocity[0], velocity[1], velocity[2]}
				hailStoneB = Hailstone{pointB[0], pointB[1], pointB[2], velocityB[0], velocityB[1], velocityB[2]}
			}

			s2, v2 := calculateSlopeAndIntercept(pointB[0], pointB[1], velocityB[0], velocityB[1])
			if ok, pX, pY := utils.LineIntersection(s1, i1, s2, v2); ok {
				if checkSign(pX-float64(point[0]), float64(velocity[0])) && checkSign(pY-float64(point[1]), float64(velocity[1])) && checkSign(pX-float64(pointB[0]), float64(velocityB[0])) && checkSign(pY-float64(pointB[1]), float64(velocityB[1])) {
					if pX >= low && pX <= max && pY >= low && pY <= max {
						// fmt.Println("A:", lines[i])
						// fmt.Println("B:", lines[j])
						// fmt.Println(pX, pY)
						count++
					}
				}
			}
		}
	}

	fmt.Println("Part 1:", count)

	// // https://paulbourke.net/geometry/pointlineplane/
	rvx, rvy, rvz := getPossibleVelocity(vxx), getPossibleVelocity(vyy), getPossibleVelocity(vzz)
	//fmt.Println(rvx, rvy, rvz)

	mA := (float64(hailStoneA.vy) - rvy) / (float64(hailStoneA.vx) - rvx)
	mB := (float64(hailStoneB.vy) - rvy) / (float64(hailStoneB.vx) - rvx)
	cA := float64(hailStoneA.py) - (mA * float64(hailStoneA.px))
	cB := float64(hailStoneB.py) - (mB * float64(hailStoneB.px))
	xPos := (cB - cA) / (mA - mB)
	yPos := mA*xPos + cA
	time := (xPos - float64(hailStoneA.px)) / (float64(hailStoneA.vx) - rvx)
	zPos := float64(hailStoneA.pz) + (float64(hailStoneA.vz)-rvz)*time

	fmt.Println("Part 2:", int(xPos+yPos+zPos))
}

func getPossibleVelocity(vx map[int][]int) float64 {
	possvx := map[int]bool{}
	for v, d := range vx {
		if len(d) == 1 {
			continue
		}
		possible := map[int]bool{}
		for i := 0; i < len(d)-1; i++ {
			for j := i + 1; j < len(d); j++ {
				ddiff := utils.Abs(d[i] - d[j])
				for pv := -1000; pv <= 1000; pv++ {
					if pv != v && ddiff%(pv-v) == 0 {
						possible[pv] = true
					}
				}
			}
		}
		if len(possvx) == 0 {
			possvx = possible
		} else {
			possvx = intersect(possible, possvx)
		}
	}

	if len(possvx) != 1 {
		panic("not one possible velocity")
	}

	for k := range possvx {
		return float64(k)
	}

	return 0
}

func intersect(a, b map[int]bool) map[int]bool {
	res := map[int]bool{}
	for k := range a {
		if _, ok := b[k]; ok {
			res[k] = true
		}
	}
	return res
}

func checkSign(val1, val2 float64) bool {
	return (val1 > 0 && val2 > 0) || (val1 < 0 && val2 < 0)
}

func calculateSlopeAndIntercept(pX, pY, vX, vY int) (slope float64, intercept float64) {
	slope = float64(vY) / float64(vX)
	intercept = float64(pY) - slope*float64(pX)
	return
}
