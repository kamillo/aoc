package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/kamillo/aoc/utils"
)

var versionSum int = 0

func main() {
	lines := utils.GetLines("input.txt")
	line := lines[0]

	b := ""
	for _, c := range line {
		h, _ := strconv.ParseUint(string(c), 16, 64)
		b += fmt.Sprintf("%04b", h)
	}

	_, value := readPacket(0, b)

	fmt.Println("Part 1: ", versionSum)
	fmt.Println("Part 2: ", value)
}

func readPacket(start int, b string) (uint64, uint64) {
	version, _ := strconv.ParseUint(b[start:start+3], 2, 64)
	typeID, _ := strconv.ParseUint(b[start+3:start+6], 2, 64)

	value := uint64(0)
	if typeID == 1 {
		value = 1
	}

	versionSum += int(version)

	ptr := start + 6

	if typeID == 4 {
		literal := ""
		for i := ptr; i < len(b)-4; i += 5 {
			group := b[i+1 : i+5]
			literal += group
			ptr += 5
			if b[i] == '0' {
				break
			}
		}
		v, _ := strconv.ParseInt(literal, 2, 64)
		value = uint64(v)

	} else {
		lengthTypeID, _ := strconv.ParseUint(string(b[ptr]), 2, 64)
		ptr++

		if lengthTypeID == 0 {
			length, _ := strconv.ParseInt(b[ptr:ptr+15], 2, 64)
			ptr += 15
			min := uint64(math.MaxUint64)
			max := uint64(0)
			prev := -1
			for length > 0 {
				read, v := readPacket(ptr, b)
				length -= int64(read)
				ptr += int(read)
				switch typeID {
				case 0:
					value += uint64(v)
				case 1:
					value *= uint64(v)
				case 2:
					if v < min {
						min = v
					}
					value = uint64(min)
				case 3:
					if v > max {
						max = v
					}
					value = uint64(max)
				case 5:
					if prev != -1 && uint64(prev) > v {
						value = 1
					} else {
						value = 0
					}
					prev = int(v)
				case 6:
					if prev != -1 && uint64(prev) < v {
						value = 1
					} else {
						value = 0
					}
					prev = int(v)
				case 7:
					if prev != -1 && uint64(prev) == v {
						value = 1
					} else {
						value = 0
					}
					prev = int(v)
				}
			}

		} else {
			length, _ := strconv.ParseUint(b[ptr:ptr+11], 2, 64)
			ptr += 11
			min := uint64(math.MaxUint64)
			max := uint64(0)
			prev := -1
			for i := 0; i < int(length); i++ {
				read, v := readPacket(ptr, b)
				ptr += int(read)
				switch typeID {
				case 0:
					value += uint64(v)
				case 1:
					value *= uint64(v)
				case 2:
					if v < min {
						min = v
					}
					value = uint64(min)
				case 3:
					if v > max {
						max = v
					}
					value = uint64(max)
				case 5:
					if prev != -1 && uint64(prev) > v {
						value = 1
					} else {
						value = 0
					}
					prev = int(v)
				case 6:
					if prev != -1 && uint64(prev) < v {
						value = 1
					} else {
						value = 0
					}
					prev = int(v)
				case 7:
					if prev != -1 && uint64(prev) == v {
						value = 1
					} else {
						value = 0
					}
					prev = int(v)
				}
			}
		}
	}

	return uint64(ptr - start), value
}
