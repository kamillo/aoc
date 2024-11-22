package main

import (
	"fmt"
	"github.com/kamillo/aoc/utils"
	"strconv"
	"strings"
)

func main() {
	persons := map[string]bool{}
	happiness := map[string]int{}

	for _, line := range utils.GetLines("input.txt") {
		split := strings.Split(line, " ")
		person1 := split[0]
		person2 := split[len(split)-1]
		person2 = strings.Trim(person2, ".")
		happy, _ := strconv.Atoi(split[3])
		if split[2] == "lose" {
			happy = -happy
		}

		persons[person1] = true
		happiness[person1+person2] = happy
	}

	fmt.Println("Part 1: ", getOptimal(persons, happiness))

	for p := range persons {
		happiness["me"+p] = 0
		happiness[p+"me"] = 0
	}
	persons["me"] = true
	fmt.Println("Part 2: ", getOptimal(persons, happiness))
}

func getOptimal(persons map[string]bool, happiness map[string]int) int {
	personsArray := make([]interface{}, 0, len(persons))
	for k, _ := range persons {
		personsArray = append(personsArray, k)
	}

	perm := utils.HeapPermutation(personsArray)
	max := 0
	for p1 := range perm {
		sum := 0
		for p2 := range perm[p1] {
			p3 := p2 + 1
			if p3 >= len(perm[p1]) {
				p3 = 0
			}

			sum += happiness[perm[p1][p2].(string)+perm[p1][p3].(string)]
			sum += happiness[perm[p1][p3].(string)+perm[p1][p2].(string)]
		}
		if sum > max {
			max = sum
		}
	}

	return max
}
