package main

import (
	"fmt"
)

func main() {
	card := 6270530
	door := 14540258

	loop := 0
	for i := 1; i != card; i = i * 7 % 20201227 {
		loop++
	}

	fmt.Println(loop)
	p1 := 1
	for i := 0; i < loop; i++ {
		p1 = p1 * door % 20201227
	}
	fmt.Println(p1)

	loop = 0
	for i := 1; i != door; i = i * 7 % 20201227 {
		loop++
	}

	fmt.Println(loop)
	p2 := 1
	for i := 0; i < loop; i++ {
		p2 = p2 * card % 20201227
	}
	fmt.Println(p1)
}
