package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	sum := 0
	northpoleObjectId := 0

	for _, line := range lines {
		checksum := line[len(line)-6 : len(line)-1]
		line = line[:len(line)-7]

		split := strings.Split(line, "-")
		id, _ := strconv.Atoi(split[len(split)-1])
		split = split[:len(split)-1]

		chars := map[string]int{}
		for _, s := range split {
			for _, c := range s {
				chars[string(c)]++
			}
		}

		sorted := utils.SortMapStringInt(chars, true)

		calculatedChecksum := ""
		for _, s := range sorted {
			calculatedChecksum += s.Key
			if len(calculatedChecksum) == 5 {
				break
			}
		}
		//fmt.Println(line, checksum, calculatedChecksum, id, split, chars)

		if calculatedChecksum == checksum {
			sum += id

			decrypted := strings.Map(func(r rune) rune {
				if r != ' ' {
					return 'a' + rune(r-'a'+rune(id))%26
				}
				return r
			}, strings.Join(split, " "))

			if decrypted == "northpole object storage" {
				northpoleObjectId = id
				break
			}
		}
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", northpoleObjectId)
}
