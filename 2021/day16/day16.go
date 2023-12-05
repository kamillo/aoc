package main

import (
	"fmt"
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

	readPacket(0, b)

	fmt.Println("Part 1: ", versionSum)
}

func readPacket(start int, b string) uint64 {
	version, _ := strconv.ParseUint(b[start:start+3], 2, 64)
	typeID, _ := strconv.ParseUint(b[start+3:start+6], 2, 64)
	//fmt.Println("version : ", version, "type : ", typeID)

	versionSum += int(version)

	ptr := start + 6
	if typeID == 4 {
		literal := ""
		for i := ptr; i < len(b)-5; i += 5 {
			group := b[i+1 : i+5]
			literal += group
			ptr += 5
			if b[i] == '0' {
				break
			}
		}
	} else {
		lengthTypeID, _ := strconv.ParseUint(string(b[ptr]), 2, 64)
		ptr++

		if lengthTypeID == 0 {
			length, _ := strconv.ParseInt(b[ptr:ptr+15], 2, 64)
			ptr += 15
			for length > 0 {
				read := readPacket(ptr, b)
				length -= int64(read)
				ptr += int(read)
			}
		} else {
			length, _ := strconv.ParseUint(b[ptr:ptr+11], 2, 64)
			ptr += 11
			for i := 0; i < int(length); i++ {
				read := readPacket(ptr, b)
				ptr += int(read)
			}
		}
	}

	return uint64(ptr - start)
}
