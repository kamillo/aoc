package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "cxdnnyjw"
	// input := "abc"

	password := ""
	password2 := [8]string{}
	pass2Count := 0
	for i := 0; ; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(input, i))))
		if strings.HasPrefix(hash, "00000") {
			if len(password) < 8 {
				password += string(hash[5])
			}
			pos, err := strconv.Atoi(string(hash[5]))
			if pos < 8 && err == nil && password2[pos] == "" {
				password2[pos] = string(hash[6])
				pass2Count++
			}

			if pass2Count == 8 {
				break
			}
		}
	}
	fmt.Println("Part1:", password)
	fmt.Println("Part2:", strings.Join(password2[:], ""))
}
