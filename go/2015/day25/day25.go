package main

import "fmt"

func main() {
	a := 20151125
	m := 252533
	r := 33554393

	for i := 2; ; i++ {
		for x := 1; x <= i; x++ {
			row := i - (x - 1)
			col := x
			a = (a * m) % r

			if row == 2947 && col == 3029 {
				fmt.Println("Part 1: ", a)
				return
			}
		}
	}
}
