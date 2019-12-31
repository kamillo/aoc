package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
)

const m = 119315717514047
const n = 101741582076661
const position = 2020

type calc func(int64, int64) int64

func expBySqrt(f calc, unit int64, a int64, b int64) int64 {
	for r := unit; ; {

		if b == 0 {
			return r
		}
		if b&1 == 1 {
			r = f(r, a)
		}
		b >>= 1
		a = f(a, a)
	}
}

func add(a int64, b int64) int64 { return (m + (a+b)%m) % m }       // +  (mod m)
func mul(a int64, b int64) int64 { return expBySqrt(add, 0, a, b) } // *  (mod m)
func pow(a int64, b int64) int64 { return expBySqrt(mul, 1, a, b) } // ** (mod m)

func main() {
	lines := fileutil.GetLines("input.txt")

	deck := make([]int, 10007)
	for i := range deck {
		deck[i] = i
	}

	for _, line := range lines {
		newDeck := make([]int, len(deck))
		var value int
		if line == "deal into new stack" {
			for i, j := 0, len(deck)-1; i < j; i, j = i+1, j-1 {
				deck[i], deck[j] = deck[j], deck[i]
			}
			newDeck = deck
		}
		if n, _ := fmt.Sscanf(line, "deal with increment %d", &value); n > 0 {
			newPosition := 0
			for i := range deck {
				newDeck[newPosition] = deck[i]
				newPosition = (newPosition + value) % len(deck)
			}
		}
		if n, _ := fmt.Sscanf(line, "cut %d", &value); n > 0 {
			if value < 0 {
				newDeck = deck[len(deck)+value:]
				newDeck = append(newDeck, deck[0:len(deck)+value]...)
			} else {
				newDeck = deck[value:]
				newDeck = append(newDeck, deck[0:value]...)
			}
		}
		deck = newDeck
	}

	for i := range deck {
		if deck[i] == 2019 {
			fmt.Println("Part 1: ", i)
			break
		}
	}

	var k, b, x int64
	k, b, x = 1, 0, 0
	for _, line := range lines {
		var value int64
		if line == "deal into new stack" {
			k = add(0, -k)
			b = add(-1, -b)
		}
		if n, _ := fmt.Sscanf(line, "deal with increment %d", &value); n > 0 {
			x = value
			k = mul(k, x)
			b = mul(b, x)
		}
		if n, _ := fmt.Sscanf(line, "cut %d", &value); n > 0 {
			b = add(b, -value)
		}
		//fmt.Println(k, b, x)
	}
	x = mul(b, pow(k-1, m-2)) // compute (λ c => k*c + b)**-n and feed it position
	fmt.Println("Part 2: ", add(mul(add(x, position), pow(pow(k, m-2), n)), -x))

}
