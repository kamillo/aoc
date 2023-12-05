package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"

	"github.com/kamillo/aoc/utils"
)

type Packets []any

func main() {
	lines := utils.GetLines("input.txt")

	sum := 0
	for i := 0; i < len(lines); i += 3 {
		line1 := lines[i]
		line2 := lines[i+1]

		var t1, t2 any
		json.Unmarshal([]byte(line1), &t1)
		json.Unmarshal([]byte(line2), &t2)

		if compare(t1, t2) < 0 {
			sum += (i / 3) + 1
		}
	}

	fmt.Println("Part 1:", sum)

	packets := Packets{}
	lines = append(lines, "[[2]]", "[[6]]")
	for _, line := range lines {
		var t any
		json.Unmarshal([]byte(line), &t)
		if t != nil {
			packets = append(packets, t)
		}
	}

	sort.Sort(packets)
	res := 1
	for i, x := range packets {
		s := fmt.Sprintf("%v", x)
		if s == "[[2]]" || s == "[[6]]" {
			res *= i + 1
		}
	}
	fmt.Println("Part 2:", res)
}

func compare(a, b interface{}) int {
	aList := false
	bList := false

	switch a.(type) {
	case float64:
		aList = false
	case []interface{}:
		aList = true
	}

	switch b.(type) {
	case float64:
		bList = false
	case []interface{}:
		bList = true
	}

	if !aList && !bList {
		return int(a.(float64) - b.(float64))
	}

	if aList && bList {
		aLen := len(a.([]interface{}))
		bLen := len(b.([]interface{}))

		for i := 0; i < int(math.Min(float64(aLen), float64(bLen))); i++ {
			res := compare(a.([]interface{})[i], b.([]interface{})[i])
			if res != 0 {
				return res
			}
		}
		return aLen - bLen
	}

	if !aList && bList {
		return compare([]interface{}{a}, b)
	}

	if aList && !bList {
		return compare(a, []interface{}{b})
	}

	fmt.Println(aList, bList)

	panic("should not happen")
}

func (a Packets) Len() int           { return len(a) }
func (a Packets) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Packets) Less(i, j int) bool { return compare(a[i], a[j]) < 0 }
