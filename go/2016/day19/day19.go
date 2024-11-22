package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

const input = 3014387

// const input = 5

type Elf struct {
	Id   int
	Gift int
}

func main() {
	elfs := ring.New(input)
	for i := 0; i < input; i++ {
		elfs.Value = Elf{i + 1, 1}
		elfs = elfs.Next()
	}

	for {
		elfs.Unlink(1)
		elfs = elfs.Next()
		if elfs.Value.(Elf).Id == elfs.Next().Value.(Elf).Id {
			break
		}
	}
	fmt.Println("Part 1: ", elfs.Value.(Elf).Id)

	elfsA := list.New()
	elfsB := list.New()
	for i := 0; i < input; i++ {
		if i <= input/2 {
			elfsA.PushBack(Elf{i + 1, 1})
		} else {
			elfsB.PushBack(Elf{i + 1, 1})
		}
	}

	for elfsA.Len()+elfsB.Len() > 1 {
		if (elfsA.Len()+elfsB.Len())%2 != 0 {
			elfsA.Remove(elfsA.Back())
		} else {
			elfsB.Remove(elfsB.Front())
		}

		if elfsA.Len() > 0 {
			elfsB.PushBack(elfsA.Front().Value)
			elfsA.Remove(elfsA.Front())
		}

		if elfsB.Len() > 0 {
			elfsA.PushBack(elfsB.Front().Value)
			elfsB.Remove(elfsB.Front())
		}
	}

	fmt.Println("Part 2: ", elfsA.Front().Value.(Elf).Id)
}
