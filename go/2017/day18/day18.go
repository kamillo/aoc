package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kamillo/aoc/utils"
)

type Pair struct{ Id, Result int }

func main() {
	lines := utils.GetLines("input.txt")
	fmt.Println("Part 1:", checkFreq(lines))

	c1 := make(chan int, 10000)
	c2 := make(chan int, 10000)
	done := make(chan Pair, 2)
	//fmt.Println("Part 1:", prog(utils.GetLines("test2.txt")))

	var wg sync.WaitGroup

	wg.Add(2)
	go prog(0, lines, c1, c2, done, &wg)
	go prog(1, lines, c2, c1, done, &wg)

	wg.Wait()
	close(done)

	for p := range done {
		if p.Id == 1 {
			fmt.Println("Part 2: ", p.Result)
		}
	}
}

func checkFreq(lines []string) int {
	registers := map[string]int{}
	lastFreq := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fields := strings.Fields(line)

		b := 0
		if len(fields) > 2 {
			b = registers[fields[2]]
			if x, error := strconv.Atoi(fields[2]); error == nil {
				b = x
			}
		}

		switch fields[0] {
		case "snd":
			lastFreq = registers[fields[1]]
		case "set":
			registers[fields[1]] = b
		case "add":
			registers[fields[1]] += b
		case "mul":
			registers[fields[1]] *= b
		case "mod":
			registers[fields[1]] %= b
		case "rcv":
			if registers[fields[1]] != 0 {
				return lastFreq
			}
		case "jgz":
			if registers[fields[1]] > 0 {
				i += b
				i--
			}
		}

	}

	return lastFreq
}

func prog(id int, lines []string, inQueue, outQueue chan int, done chan Pair, wg *sync.WaitGroup) {
	defer wg.Done()
	registers := map[string]int{}
	registers["p"] = id
	cnt := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fields := strings.Fields(line)

		b := 0
		if len(fields) > 2 {
			b = registers[fields[2]]
			if x, error := strconv.Atoi(fields[2]); error == nil {
				b = x
			}
		}

		switch fields[0] {
		case "snd":
			cnt++
			outQueue <- registers[fields[1]]
			// fmt.Println("Prog -> ", id, registers[fields[1]])
		case "set":
			registers[fields[1]] = b
		case "add":
			registers[fields[1]] += b
		case "mul":
			registers[fields[1]] *= b
		case "mod":
			registers[fields[1]] %= b
		case "rcv":
			select {
			case registers[fields[1]] = <-inQueue:
				// fmt.Println("Prog <- ", id, registers[fields[1]])

			case <-time.After(1 * time.Second):
				// fmt.Println("timeout ", id, cnt)
				done <- Pair{id, cnt}
				return
			}
		case "jgz":
			a := 0
			if len(fields) > 2 {
				a = registers[fields[1]]
				if x, error := strconv.Atoi(fields[1]); error == nil {
					a = x
				}
			}
			if a > 0 {
				i += b
				i--
			}
		}

	}
}
