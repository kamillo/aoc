package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	input := "yzbqklnj"

	for i := 0; ; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(input, i))))
		if strings.HasPrefix(hash, "00000") {
			fmt.Println("Part 1: ", i)
			break
		}
	}

	for i := 0; ; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(input, i))))
		if strings.HasPrefix(hash, "000000") {
			fmt.Println("Part 2: ", i)
			break
		}
	}
}
