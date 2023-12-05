package main

import (
	"bytes"
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"strings"
	"time"
)

func main() {
	clear()
	lines := fileutil.GetLines("input.txt")
	seats := make([][]byte, len(lines))
	for s, line := range lines {
		seats[s] = []byte(line)
	}

	newSeats := make([][]byte, len(lines))
	copy(newSeats, seats)
	for seat := range seats {
		newSeats[seat] = make([]byte, len(lines[0]))
		copy(newSeats[seat], seats[seat])
	}
	runSimulation(newSeats, countAdj, 4)
	fmt.Println("Part 1: ", strings.Count(string(bytes.Join(newSeats, []byte{})), "#"))

	newSeats = make([][]byte, len(lines))
	copy(newSeats, seats)
	for seat := range seats {
		newSeats[seat] = make([]byte, len(lines[0]))
		copy(newSeats[seat], seats[seat])
	}
	runSimulation(newSeats, countVisible, 5)
	fmt.Println("Part 2: ", strings.Count(string(bytes.Join(newSeats, []byte{})), "#"))
}

func runSimulation(newSeats [][]byte, countFunc func(int, int, [][]byte) int, emptyLimit int) {
	for i := 0; i < 200; i++ {
		seats := make([][]byte, len(newSeats))
		toPrint := ""
		copy(seats, newSeats)
		for seat := range newSeats {
			seats[seat] = make([]byte, len(newSeats[0]))
			copy(seats[seat], newSeats[seat])
		}

		for x := range seats {
			for y := range seats[x] {
				switch seats[x][y] {
				case 'L':
					if countFunc(x, y, seats) == 0 {
						newSeats[x][y] = '#'
					}

					break
				case '#':
					if countFunc(x, y, seats) >= emptyLimit {
						newSeats[x][y] = 'L'
					}
					break
				}
			}
			toPrint += string(seats[x]) + "\n"
		}
		//debugPrint(toPrint, i)
	}
}

func countAdj(x int, y int, seats [][]byte) (adj int) {
	if x+1 < len(seats) && seats[x+1][y] == '#' {
		adj++
	}
	if y+1 < len(seats[x]) && seats[x][y+1] == '#' {
		adj++
	}
	if x+1 < len(seats) && y+1 < len(seats[x]) && seats[x+1][y+1] == '#' {
		adj++
	}
	if x > 0 && seats[x-1][y] == '#' {
		adj++
	}
	if x > 0 && y+1 < len(seats[x]) && seats[x-1][y+1] == '#' {
		adj++
	}
	if y > 0 && seats[x][y-1] == '#' {
		adj++
	}
	if y > 0 && x+1 < len(seats) && seats[x+1][y-1] == '#' {
		adj++
	}
	if x > 0 && y > 0 && seats[x-1][y-1] == '#' {
		adj++
	}

	return adj
}

func countVisible(x int, y int, seats [][]byte) (adj int) {
	checkSeat := func(seat byte, adj int) (bool, int) {
		if seat == 'L' {
			return true, adj
		}
		if seat == '#' {
			return true, adj + 1
		}
		return false, adj
	}
	var stop bool
	for xx := x + 1; xx < len(seats); xx++ {
		if stop, adj = checkSeat(seats[xx][y], adj); stop {
			break
		}
	}
	for yy := y + 1; yy < len(seats[x]); yy++ {
		if stop, adj = checkSeat(seats[x][yy], adj); stop {
			break
		}
	}
	for yy, xx := y+1, x+1; yy < len(seats[x]) && xx < len(seats); yy, xx = yy+1, xx+1 {
		if stop, adj = checkSeat(seats[xx][yy], adj); stop {
			break
		}
	}
	for xx := x - 1; xx >= 0; xx-- {
		if stop, adj = checkSeat(seats[xx][y], adj); stop {
			break
		}
	}

	for xx, yy := x-1, y+1; xx >= 0 && yy < len(seats[x]); xx, yy = xx-1, yy+1 {
		if stop, adj = checkSeat(seats[xx][yy], adj); stop {
			break
		}
	}

	for yy := y - 1; yy >= 0; yy-- {
		if stop, adj = checkSeat(seats[x][yy], adj); stop {
			break
		}
	}

	for yy, xx := y-1, x+1; yy >= 0 && xx < len(seats); yy, xx = yy-1, xx+1 {
		if stop, adj = checkSeat(seats[xx][yy], adj); stop {
			break
		}
	}

	for yy, xx := y-1, x-1; yy >= 0 && xx >= 0; yy, xx = yy-1, xx-1 {
		if stop, adj = checkSeat(seats[xx][yy], adj); stop {
			break
		}
	}
	return adj
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func debugPrint(toPrint string, i int) {
	if i%2 == 0 {
		clear()
		toPrint = strings.ReplaceAll(toPrint, "L", "⬛️️")
		toPrint = strings.ReplaceAll(toPrint, "#", "🟪")
		toPrint = strings.ReplaceAll(toPrint, ".", "️🟨")
		fmt.Print(toPrint)
		time.Sleep(time.Millisecond * 150)
	}
}
