package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	//target area: x=150..193, y=-136..-86
	//target area: x=20..30, y=-10..-5
	// targetStart := utils.NewPoint2D(20, -5)
	// targetEnd := utils.NewPoint2D(30, -10)

	targetStart := utils.NewPoint2D(150, -86)
	targetEnd := utils.NewPoint2D(193, -136)

	// velocity := utils.NewPoint2D(6, 9)

	fire := func(velocity utils.Point2D) (utils.Point2D, bool) {
		probe := utils.NewPoint2D(0, 0)
		max := utils.NewPoint2D(0, 0)
		for {
			probe.X += velocity.X
			probe.Y += velocity.Y

			if velocity.X > 0 {
				velocity.X--
			} else {
				velocity.X++
			}

			velocity.Y--

			if max.Y < probe.Y {
				max.Y = probe.Y
			}

			//fmt.Println(probe, velocity)

			if probe.X > targetEnd.X || probe.Y < targetEnd.Y {
				//fmt.Println("MISS ", max)
				return utils.NewPoint2D(0, 0), false
			}

			if probe.X >= targetStart.X && probe.X <= targetEnd.X && probe.Y <= targetStart.Y && probe.Y >= targetEnd.Y {
				//fmt.Println("HIT ", max)
				return max, true
			}
		}
		return max, false
	}

	max := 0
	count := 0
	for x := 0; x < targetEnd.X+1; x++ {
		for y := targetEnd.Y; y < 400; y++ {
			p, hit := fire(utils.NewPoint2D(x, y))
			if p.Y > max {
				max = p.Y
			}
			if hit {
				count++
			}
		}
	}

	fmt.Println(max, count)
}
