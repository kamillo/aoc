package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// GetLines - Get lines from file as a Slice
func GetLines(fileName string) (lines []string) {
	// fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return
}

func GetLinesAsInts(fileName string) (lines []int) {
	for _, line := range GetLines(fileName) {
		if i, err := strconv.Atoi(line); err == nil {
			lines = append(lines, i)
		}
	}

	return
}

func GetLinesAsInterface(fileName string) (lines []interface{}) {
	for _, line := range GetLines(fileName) {
		lines = append(lines, line)
	}

	return
}

func GetLinesAs2dArray(fileName string) (result [][]byte) {
	for _, line := range GetLines(fileName) {
		result = append(result, []byte(line))
	}

	return
}
