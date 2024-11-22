package main

import (
	"fmt"
	"math"

	"github.com/kamillo/aoc/utils"
)

type Particle struct {
	position utils.PointD3D
	velocity utils.PointD3D
	acc      utils.PointD3D
	collided bool
}

func main() {

	particles := []Particle{}

	for _, line := range utils.GetLines("input.txt") {
		p := Particle{}
		//p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
		fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &(p.position.X), &(p.position.Y), &(p.position.Z), &(p.velocity.X), &(p.velocity.Y), &(p.velocity.Z), &(p.acc.X), &(p.acc.Y), &(p.acc.Z))

		particles = append(particles, p)
	}

	particlesCopy := make([]Particle, len(particles))
	copy(particlesCopy, particles)

	for i := 0; i < 1000; i++ { // pretty long run, right?
		positions := map[utils.PointD3D]int{}
		for p := range particles {
			if particles[p].collided {
				continue
			}

			particles[p].velocity.X += particles[p].acc.X
			particles[p].velocity.Y += particles[p].acc.Y
			particles[p].velocity.Z += particles[p].acc.Z

			particles[p].position.X += particles[p].velocity.X
			particles[p].position.Y += particles[p].velocity.Y
			particles[p].position.Z += particles[p].velocity.Z

			if k, ok := positions[particles[p].position]; ok {
				particles[p].collided = true
				particles[k].collided = true
			}
			positions[particles[p].position] = p
		}
	}

	min0 := math.MaxInt
	part0 := 0
	collided := 0
	for p := range particles {
		if distance0(particles[p].position) < min0 {
			min0 = distance0(particles[p].position)
			part0 = p
		}

		if particles[p].collided {
			collided++
		}
	}

	fmt.Println("Part 1:", part0)
	fmt.Println("Part 2:", len(particles)-collided)
}

func distance(a, b utils.PointD3D) int {
	return int(
		math.Abs(float64(a.X)-float64(b.X)) +
			math.Abs(float64(a.Y)-float64(b.Y)) +
			math.Abs(float64(a.Z)-float64(b.Z)))
}

func distance0(p2 utils.PointD3D) int {
	return distance(utils.NewPointD3D(0, 0, 0), p2)
}
