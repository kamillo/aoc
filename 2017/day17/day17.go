package main

import (
	"container/ring"
	"fmt"
)

func main() {
	rotate := 345
	iterations := 2017

	r := ring.New(1)
	r.Value = 0
	r = r.Next()

	for i := 1; i <= iterations; i++ {
		for j := 0; j < rotate; j++ {
			r = r.Next()
		}

		s := ring.New(1)
		s.Value = i

		r = r.Link(s)
		r = r.Prev()
	}

	fmt.Println("Part 1:", r.Next().Value.(int))

	p := 0
	next := 0
	for i := 1; i <= 50000000; i++ {
		p = ((rotate + p) % i) + 1
		if p == 1 {
			next = i
		}
	}

	fmt.Println("Part 2: ", next)
}
