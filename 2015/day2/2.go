package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/kamillo/aoc/fileutil"
	"github.com/kamillo/aoc/mathutils"
)

func main() {
	paper, ribbon := 0, 0
	lines := fileutil.GetLines(os.Args[1])

	for _, line := range lines {
		// LxWxH
		// 2*l*w + 2*w*h + 2*h*l

		splitted := strings.Split(line, "x")
		l, _ := strconv.Atoi(splitted[0])
		w, _ := strconv.Atoi(splitted[1])
		h, _ := strconv.Atoi(splitted[2])
		dimens := []int{l, w, h}
		sort.Ints(dimens)

		paper += 2*l*w + 2*w*h + 2*h*l + mathutils.Min(l*w, w*h, h*l)
		ribbon += l*w*h + 2*dimens[0] + 2*dimens[1]
	}

	fmt.Println("Part one: ", paper)
	fmt.Println("Part two: ", ribbon)
}
