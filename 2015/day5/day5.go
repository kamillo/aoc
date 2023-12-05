package main

import (
	"fmt"
	"github.com/kamillo/aoc/fileutil"
	"strings"
)

func main() {
	nice := 0
	for _, line := range fileutil.GetLines("input.txt") {
		if countVowels(line) >= 3 && countDoubledRune(line) >= 1 && !containsForbidden(line) {
			nice++
		}
	}
	fmt.Println("Part 1: ", nice)

	nice = 0
	for _, line := range fileutil.GetLines("input.txt") {
		if containsPair(line) && containsRepeatingWithSeparator(line) {
			nice++
			fmt.Println(line)
		}
	}
	fmt.Println("Part 2: ", nice)
}

func countVowels(text string) (count int) {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for _, c := range text {
		if _, ok := vowels[c]; ok {
			count++
		}
	}
	return count
}

func countDoubledRune(text string) (count int) {
	var prev rune
	for _, c := range text {
		if prev == c {
			count++
		}
		prev = c
	}
	return count
}

func containsForbidden(text string) bool {
	return strings.Contains(text, "ab") ||
		strings.Contains(text, "cd") ||
		strings.Contains(text, "pq") ||
		strings.Contains(text, "xy")
}

func containsPair(text string) bool {
	var prev, praprev rune
	pairs := make(map[string]int)
	for i, c := range text {
		praprev = prev
		prev = c
		if i >= 1 {
			pair := string(praprev) + string(prev)
			if index, ok := pairs[pair]; ok && index != i-1 {
				fmt.Println("found ", pair, text)
				return true
			}
			pairs[pair] = i
		}
	}
	return false
}

func containsRepeatingWithSeparator(text string) bool {
	var prev, praprev rune
	for _, c := range text {
		if praprev == c {
			return true
		}
		praprev = prev
		prev = c
	}
	return false
}
