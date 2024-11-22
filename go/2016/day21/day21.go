package main

import (
	"container/list"
	"container/ring"
	"fmt"

	"github.com/kamillo/aoc/utils"
)

func main() {
	input := "abcdefgh"
	lines := utils.GetLines("input.txt")

	scramble := func(lines []string) string {
		letters := ring.New(len(input))
		for i := 0; i < letters.Len(); i++ {
			letters.Value = int(input[i])
			letters = letters.Next()
		}

		for _, line := range lines {
			a, b := 0, 0

			if _, err := fmt.Sscanf(line, "swap position %d with position %d", &a, &b); err == nil {
				letters = swapPosition(letters, a, b)

			} else if _, err := fmt.Sscanf(line, "swap letter %c with letter %c", &a, &b); err == nil {
				letters = swap(letters, a, b)

			} else if _, err := fmt.Sscanf(line, "rotate left %d step", &a); err == nil {
				letters = spinLeft(letters, a)

			} else if _, err := fmt.Sscanf(line, "rotate right %d step", &a); err == nil {
				letters = spinRight(letters, a)

			} else if _, err := fmt.Sscanf(line, "rotate based on position of letter %c", &a); err == nil {
				i := indexOf(letters, a)
				if i >= 4 {
					i++
				}
				letters = spinRight(letters, i+1)

			} else if _, err := fmt.Sscanf(line, "reverse positions %d through %d", &a, &b); err == nil {
				letters = reverse(letters, a, b)

			} else if _, err := fmt.Sscanf(line, "move position %d to position %d", &a, &b); err == nil {
				letters = move(letters, a, b)
			}

			letters.Do(func(p interface{}) {
				fmt.Printf("%c", p.(int))
			})
			fmt.Println()
		}
		ret := ""
		letters.Do(func(p interface{}) {
			ret += fmt.Sprintf("%c", p.(int))
		})

		return ret
	}

	fmt.Println("Part 1: ", scramble(lines))

	input = "fbgdceah"
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	fmt.Println("Part 2: ", unscramble(lines, input))
}

func unscramble(lines []string, input string) string {
	letters := ring.New(len(input))
	for i := 0; i < letters.Len(); i++ {
		letters.Value = int(input[i])
		letters = letters.Next()
	}

	for _, line := range lines {
		a, b := 0, 0

		if _, err := fmt.Sscanf(line, "swap position %d with position %d", &a, &b); err == nil {
			letters = swapPosition(letters, b, a)

		} else if _, err := fmt.Sscanf(line, "swap letter %c with letter %c", &a, &b); err == nil {
			letters = swap(letters, b, a)

		} else if _, err := fmt.Sscanf(line, "rotate left %d step", &a); err == nil {
			letters = spinRight(letters, a)

		} else if _, err := fmt.Sscanf(line, "rotate right %d step", &a); err == nil {
			letters = spinLeft(letters, a)

		} else if _, err := fmt.Sscanf(line, "rotate based on position of letter %c", &a); err == nil {
			i := indexOf(letters, a)
			rot := 0
			switch i {
			case 0:
				rot = 7
			case 1:
				rot = 7
			case 2:
				rot = 2
			case 3:
				rot = 6
			case 4:
				rot = 1
			case 5:
				rot = 5
			case 6:
				rot = 0
			case 7:
				rot = 4
			}
			letters = spinRight(letters, rot)
		} else if _, err := fmt.Sscanf(line, "reverse positions %d through %d", &a, &b); err == nil {
			letters = reverse(letters, a, b)

		} else if _, err := fmt.Sscanf(line, "move position %d to position %d", &a, &b); err == nil {
			letters = move(letters, b, a)
		}

		// letters.Do(func(p interface{}) {
		// 	fmt.Printf("%c", p.(int))
		// })
		// fmt.Println()
	}
	ret := ""
	letters.Do(func(p interface{}) {
		ret += fmt.Sprintf("%c", p.(int))
	})

	return ret
}

func move(ring *ring.Ring, a, b int) *ring.Ring {
	l := list.New()

	var elA *list.Element
	var elB *list.Element

	for i := 0; i < ring.Len(); i++ {
		x := l.PushBack(ring.Value.(int))
		if i == a {
			elA = x
		}
		if i == b {
			elB = x
		}
		ring = ring.Next()
	}

	// ring.Do(func(p interface{}) {
	// 	fmt.Printf("%c", p.(int))
	// })

	// fmt.Println()

	el := l.Remove(elA)
	if a < b {
		l.InsertAfter(el, elB)
	} else {
		l.InsertBefore(el, elB)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		ring.Value = e.Value
		ring = ring.Next()
	}

	// ring.Do(func(p interface{}) {
	// 	fmt.Printf("%c", p.(int))
	// })

	// fmt.Println()

	return ring
	// ret := ring
	// ring = ring.Prev()

	// tmp := ring
	// for i := 0; i < a; i++ {
	// 	ring = ring.Next()
	// }
	// removed := ring.Unlink(1)
	// ring = tmp

	// tmp = ring
	// for i := 0; i < b; i++ {
	// 	ring = ring.Next()
	// }

	// ring.Link(removed)

	// ring = ret
	// if b < a {
	// 	ring = ring.Prev()
	// }

	// return ring

}

func reverse(ring *ring.Ring, a, b int) *ring.Ring {
	ret := ring
	for i := 0; i < a; i++ {
		ring = ring.Next()
	}

	letters := []int{}
	for i := 0; i <= b-a; i++ {
		letters = append(letters, ring.Value.(int))
		ring = ring.Next()
	}

	for i := 0; i <= b-a; i++ {
		ring = ring.Prev()
		ring.Value = letters[i]
	}

	return ret
}

func spinLeft(ring *ring.Ring, i int) *ring.Ring {
	ring = ring.Move(i)
	return ring
}

func spinRight(ring *ring.Ring, i int) *ring.Ring {
	ring = ring.Move(-i)
	return ring
}

func swapPosition(ring *ring.Ring, a, b int) *ring.Ring {
	max := utils.Max(a, b)
	aP := ring
	bP := ring
	ret := ring
	for i := 0; i <= max; i++ {
		if i == a {
			aP = ring
		}

		if i == b {
			bP = ring
		}

		ring = ring.Next()
	}

	tmp := aP.Value
	aP.Value = bP.Value
	bP.Value = tmp

	return ret
}

func swap(ring *ring.Ring, a, b int) *ring.Ring {
	aP := ring
	bP := ring
	ret := ring
	for i := 0; i < ring.Len(); i++ {
		if ring.Value == int(a) {
			aP = ring
		}

		if ring.Value == int(b) {
			bP = ring
		}

		ring = ring.Next()
	}

	tmp := aP.Value
	aP.Value = bP.Value
	bP.Value = tmp

	return ret
}

func indexOf(ring *ring.Ring, letter int) int {
	ret := ring
	for i := 0; i < ring.Len(); i++ {
		if ring.Value == letter {
			ring = ret
			return i
		}

		ring = ring.Next()
	}

	return 0
}

// ebcda ebcda
// edcba edcba
// abcde abcde
// bcdea bcdea
// bdeac bdeac
// abdec bdeca
// ecabd abdec
// decab decab
