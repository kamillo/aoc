package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	seeds := utils.ToIntArr(lines[0], " ")
	lines = lines[2:]

	part1 := func(s []int) []int {
		seeds := make([]int, len(seeds))
		copy(seeds, s)
		newSeeds := make([]int, len(seeds))
		for _, line := range lines {
			if strings.Contains(line, "map") {
				newSeeds = make([]int, len(seeds))
				copy(newSeeds, seeds)
			}

			if len(line) == 0 {
				copy(seeds, newSeeds)
			}

			mapping := utils.ToIntArr(line, " ")
			if len(mapping) > 0 {
				for i, seed := range seeds {
					if seed >= mapping[1] && seed < mapping[1]+mapping[2] {
						newSeeds[i] = mapping[0] + (seed - mapping[1])
					}
				}
			}
		}
		copy(seeds, newSeeds)

		return seeds
	}

	fmt.Println("Part 1:", utils.MinInArray(part1(seeds)))

	part2 := func(s []int) []int {
		seeds := make([]int, len(seeds))
		copy(seeds, s)
		newSeeds := make([]int, len(seeds))
		for _, line := range lines {
			if strings.Contains(line, "map") {
				newSeeds = make([]int, len(seeds))
				copy(newSeeds, seeds)
			}

			if len(line) == 0 {
				copy(seeds, newSeeds)
			}

			mapping := utils.ToIntArr(line, " ")
			if len(mapping) > 0 {
				for i := 0; i < len(seeds); i += 2 {
					dstBegin := mapping[0]
					// dstEnd := mapping[0] + mapping[2]
					mapSrcBeg := mapping[1]
					mapSrcEnd := mapping[1] + mapping[2]

					// <   (  >
					if seeds[i] >= mapping[1] && seeds[i] < mapSrcEnd {
						// <   (   )>
						if seeds[i]+seeds[i+1] <= mapSrcEnd {
							newSeeds[i] = dstBegin + (seeds[i] - mapping[1])

						} else {
							// <   (    >  )
							newSeeds[i] = dstBegin + (seeds[i] - mapping[1])
							newSeeds[i+1] = (mapSrcEnd) - seeds[i]

							newSeeds = append(newSeeds, mapSrcEnd)
							newSeeds = append(newSeeds, (seeds[i]+seeds[i+1])-(mapSrcEnd))

							seeds = append(seeds, mapSrcEnd)
							seeds = append(seeds, (seeds[i]+seeds[i+1])-(mapSrcEnd))
						}
						// (   <    ) >
					} else if seeds[i] < mapSrcBeg && seeds[i]+seeds[i+1] <= mapSrcEnd && seeds[i]+seeds[i+1] > mapSrcBeg {
						oldlen := newSeeds[i+1]
						newSeeds[i+1] = mapSrcBeg - seeds[i]
						seeds[i+1] = mapSrcBeg - seeds[i]

						newSeeds = append(newSeeds, dstBegin)
						newSeeds = append(newSeeds, oldlen-newSeeds[i+1])
						seeds = append(seeds, mapSrcBeg)
						seeds = append(seeds, oldlen-seeds[i+1])

						// (   <      > )
					} else if seeds[i] < mapSrcBeg && seeds[i]+seeds[i+1] > mapSrcEnd {
						oldlen := seeds[i+1]
						newSeeds[i+1] = mapSrcBeg - seeds[i]
						seeds[i+1] = mapSrcBeg - seeds[i]

						newSeeds = append(newSeeds, dstBegin)
						newSeeds = append(newSeeds, mapping[2])

						seeds = append(seeds, mapSrcBeg)
						seeds = append(seeds, mapping[2])

						newSeeds = append(newSeeds, mapSrcEnd)
						newSeeds = append(newSeeds, (seeds[i]+oldlen)-(mapSrcEnd))

						seeds = append(seeds, mapSrcEnd)
						seeds = append(seeds, (seeds[i]+oldlen)-(mapSrcEnd))
					}

				}
			}
		}
		copy(seeds, newSeeds)

		ret := []int{}
		for i := 0; i < len(seeds); i += 2 {
			ret = append(ret, seeds[i])
		}
		//fmt.Println(ret)
		return ret
	}

	fmt.Println("Part 2:", utils.MinInArray(part2(seeds)))
}
