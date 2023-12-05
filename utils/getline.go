package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
