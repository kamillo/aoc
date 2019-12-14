package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/2019/intcode"
	"github.com/kamillo/aoc/fileutil"
)

func main() {
	lines := fileutil.GetLines("input.txt")
	//lines = []string{"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"}
	//lines = []string{"1102,34915192,34915192,7,4,7,99,0"}
	//lines = []string{"104,1125899906842624,99"}

	splitted := strings.Split(lines[0], ",")
	ints := make([]int, len(splitted)*100)
	for i := 0; i < len(splitted); i++ {
		value, _ := strconv.Atoi(splitted[i])
		ints[i] = value
	}

	intCode := intcode.Make(ints)
	intCode.Put([]int{1})
	fmt.Println(intCode.Get())
	intCode = intcode.Make(ints)
	intCode.Put([]int{2})
	fmt.Println(intCode.Get())
}
