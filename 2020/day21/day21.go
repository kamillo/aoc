package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"sort"
	"strings"
)

func main() {
	lines := fileutil.GetLines("input.txt")

	allergensMap := make(map[string]map[string]bool)
	ingredientsTimes := make(map[string]int)
	allIngredients := make(map[string]bool)
	for _, line := range lines {
		allergensIndex := strings.Index(line, "(")
		allergens := strings.Split(line[allergensIndex+10:len(line)-1], ", ")
		ingredients := strings.Split(line[:allergensIndex-1], " ")

		for _, i := range ingredients {
			ingredientsTimes[i]++
			allIngredients[i] = true
		}

		for _, a := range allergens {
			if _, ok := allergensMap[a]; !ok {
				allergensMap[a] = make(map[string]bool)
				for _, i := range ingredients {
					allergensMap[a][i] = true
				}
			} else {
				for ing := range allergensMap[a] {
					found := false
					for _, i := range ingredients {
						if allergensMap[a][ing] && ing == i {
							found = true
							break
						}
					}
					if !found {
						delete(allergensMap[a], ing)
					}
				}
			}
		}
	}

	part1 := 0
	goodIngredients := make(map[string]bool)
	for i := range allIngredients {
		bad := false
		for a := range allergensMap {
			if v, ok := allergensMap[a][i]; ok {
				bad = v
				break
			}
		}
		if !bad {
			goodIngredients[i] = true
		}
	}
	for i := range goodIngredients {
		part1 += ingredientsTimes[i]
	}
	fmt.Println("Part 1:", part1)

	difference := func(set map[string]bool, set2 map[string]interface{}) map[string]interface{} {
		n := make(map[string]interface{})

		for k, _ := range set {
			if v, exists := set2[k]; !exists {
				n[k] = v
			}
		}
		return n
	}

	sortedKeys := func(m map[string]interface{}) []string {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		return keys
	}

	valuesSortedByKeys := func(m map[string]interface{}) []string {
		keys := make([]string, 0, len(m))
		for _, v := range sortedKeys(m) {
			keys = append(keys, fmt.Sprintf("%v", m[v]))
		}
		return keys
	}

	taken := make(map[string]interface{})
	items := make(map[string]interface{})
	found := true
	for found {
		for k, v := range allergensMap {
			diff := difference(v, taken)
			if len(diff) == 1 {
				i := sortedKeys(diff)[0]
				items[k] = i
				taken[i] = true
				found = true
				break
			} else {
				found = false
			}
		}
	}

	fmt.Println("Part 2: ", strings.Join(valuesSortedByKeys(items), ","))
}
