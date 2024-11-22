package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	// targetStart := utils.NewPoint2D(20, -5)
	// targetEnd := utils.NewPoint2D(30, -10)
	targetStart := utils.NewPoint2D(150, -86)
	targetEnd := utils.NewPoint2D(193, -136)

	fire := func(velocity utils.Point2D) (utils.Point2D, bool) {
		probe := utils.NewPoint2D(0, 0)
		max := utils.NewPoint2D(0, 0)

		for {
			probe.X += velocity.X
			probe.Y += velocity.Y

			if velocity.X == 0 {
				velocity.X = 0
			} else if velocity.X > 0 {
				velocity.X--
			} else {
				velocity.X++
			}

			velocity.Y--

			if max.Y < probe.Y {
				max.Y = probe.Y
			}

			if velocity.Y < 0 && probe.Y < targetEnd.Y {
				return utils.NewPoint2D(0, 0), false
			}

			if probe.X >= targetStart.X && probe.X <= targetEnd.X && probe.Y <= targetStart.Y && probe.Y >= targetEnd.Y {
				return max, true
			}
		}
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
