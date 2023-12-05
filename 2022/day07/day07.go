package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Dir struct {
	parent string
	size   int
}

func main() {
	lines := utils.GetLines("input.txt")

	fs := map[string]*Dir{}
	currentDir := "/"
	fs["/"] = &Dir{"", 0}

	updateSize := func(dir string, size int) {
		parent := fs[dir].parent
		for parent != "" {
			fs[parent].size += size
			parent = fs[parent].parent
		}
	}

	for _, line := range lines {
		fields := strings.Fields(line)

		if strings.HasPrefix(line, "$") {
			if fields[1] == "cd" {
				dir := fields[2]

				if dir == ".." {
					currentDir = fs[currentDir].parent

				} else if dir != "/" {
					newDir := currentDir + "/" + dir
					if currentDir == "/" {
						newDir = currentDir + dir
					}

					fs[newDir] = &Dir{currentDir, 0}
					currentDir = newDir
				}
			}

		} else {
			size := 0
			name := ""
			if _, err := fmt.Sscanf(line, "%d %s", &size, &name); err == nil {
				fs[currentDir].size += size
				updateSize(currentDir, size)
			}

		}

	}

	sum := 0
	for _, v := range fs {
		if v.size <= 100000 {
			sum += v.size
		}
	}

	fmt.Println("Part 1:", sum)

	min := math.MaxInt64
	for _, v := range fs {
		if fs["/"].size-v.size <= 70000000-30000000 {
			if v.size < min {
				min = v.size
			}
		}
	}

	fmt.Println("Part 2:", min)
}
