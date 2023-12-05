package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"math"
)

type Moon struct {
	position [3]int
	velocity [3]int
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	lines := fileutil.GetLines("input.txt")
	//lines = []string{
	//	"<x=-1, y=0, z=2>",
	//	"<x=2, y=-10, z=-7>",
	//	"<x=4, y=-8, z=8>",
	//	"<x=3, y=5, z=-1>",
	//}

	moons := []Moon{}
	for _, line := range lines {
		x, y, z := 0, 0, 0
		_, _ = fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moons = append(moons, Moon{
			position: [3]int{x, y, z},
			velocity: [3]int{0, 0, 0},
		})

	}
	fmt.Println(moons)
	moonsInitial := make([]Moon, len(moons))
	copy(moonsInitial, moons)

	found := make(map[int]int)
	for i := 0; ; i++ {
		for m := range moons {
			for n := m + 1; n < len(moons); n++ {
				for j := 0; j < 3; j++ {
					if moons[m].position[j] > moons[n].position[j] {
						moons[m].velocity[j]--
						moons[n].velocity[j]++
					} else if moons[m].position[j] < moons[n].position[j] {
						moons[m].velocity[j]++
						moons[n].velocity[j]--
					}
				}
			}
		}

		foundX := true
		foundY := true
		foundZ := true
		for m := range moons {
			for j := 0; j < 3; j++ {
				moons[m].position[j] += moons[m].velocity[j]
			}
			foundX = foundX && moons[m].position[0] == moonsInitial[m].position[0]
			foundY = foundY && moons[m].position[1] == moonsInitial[m].position[1]
			foundZ = foundZ && moons[m].position[2] == moonsInitial[m].position[2]
		}
		if foundX && found['x'] == 0 {
			found['x'] = i + 2
		}
		if foundY && found['y'] == 0 {
			found['y'] = i + 2
		}
		if foundZ && found['z'] == 0 {
			found['z'] = i + 2
		}

		if i == 999 {
			fmt.Println(moons)
			energy := 0
			for m := range moons {
				pot := 0
				kin := 0
				for j := 0; j < 3; j++ {
					pot += int(math.Abs(float64(moons[m].position[j])))
					kin += int(math.Abs(float64(moons[m].velocity[j])))
				}
				energy += pot * kin
			}
			fmt.Println("Part 1: ", energy)
		}

		if len(found) == 3 {
			break
		}
	}
	fmt.Println("Part 2: ", LCM(found['x'], found['y'], found['z']))

}
