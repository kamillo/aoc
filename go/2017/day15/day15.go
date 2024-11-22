package main

import "fmt"

func main() {
	factorA := 16807
	factorB := 48271

	// startA := 65
	// startB := 8921
	startA := 873
	startB := 583

	iterations := 40000000

	genA := generator(startA, factorA, iterations)
	genB := generator(startB, factorB, iterations)

	cnt := 0
	for i := 0; i < iterations; i++ {
		if genA[i] == genB[i] {
			cnt++
		}
	}

	fmt.Println("Part 1:", cnt)

	iterations = 5000000
	genA = generator2(startA, factorA, iterations, 4)
	genB = generator2(startB, factorB, iterations, 8)

	cnt = 0
	for i := 0; i < iterations; i++ {
		if genA[i] == genB[i] {
			cnt++
		}
	}

	fmt.Println("Part 2:", cnt)
}

func generator(start, factor, iterations int) (ret []int) {
	prev := start

	for len(ret) < iterations {
		prev = (prev * factor) % 2147483647
		ret = append(ret, prev&0x0000ffff)
	}

	return
}

func generator2(start, factor, iterations, multiples int) (ret []int) {
	prev := start

	for len(ret) < iterations {
		prev = (prev * factor) % 2147483647
		if prev%multiples == 0 {
			ret = append(ret, prev&0x0000ffff)
		}
	}

	return
}
