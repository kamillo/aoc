package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	// salt := "abc"
	salt := "cuanljph"

	oneTimePad := func(part2 bool) int {
		trios := map[int]rune{}
		quintets := map[rune][]int{}

		key := 0

		for i := 0; i < 100000; i++ {
			key = 0
			salted := fmt.Sprintf("%s%d", salt, i)
			hash := fmt.Sprintf("%x", md5.Sum([]byte(salted)))

			if part2 {
				for j := 0; j < 2016; j++ {
					hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
				}
			}

			prev := '#'
			same := 0

			for _, c := range hash {
				if c == prev {
					same++
					if same == 2 {
						//fmt.Println(hash, string(c))
						if _, exist := trios[i]; !exist {
							trios[i] = c
						}
					}
					if same == 4 {
						quintets[c] = append(quintets[c], i)
					}
				} else {
					same = 0
				}

				prev = c
			}
		}

		// fmt.Println(trios)

		for k := 0; k < 100000; k++ {
			if v, ok := trios[k]; ok {
				ids, exist := quintets[v]
				if exist {
					for _, id := range ids {
						if id-k > 0 && id-k < 1000 {
							key++
							break
						}
					}
				}

				if key == 64 {
					return k
				}
			}
		}

		return 0
	}

	fmt.Println("Part 1:", oneTimePad(false))
	fmt.Println("Part 2:", oneTimePad(true))
}
