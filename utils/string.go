package utils

import (
	"strconv"
	"strings"
)

func ToIntArr(s string, sep string) (ret []int) {
	for _, v := range strings.Split(s, sep) {
		if i, err := strconv.Atoi(v); err == nil {
			ret = append(ret, i)
		}
	}

	return
}

func ToIntSet(s string, sep string) (ret map[int]bool) {
	ret = map[int]bool{}
	for _, v := range strings.Split(s, sep) {
		if i, err := strconv.Atoi(v); err == nil {
			ret[i] = true
		}
	}

	return
}

func Reverse(s string) (rs string) {
	for _, r := range s {
		rs = string(r) + rs
	}
	return
}
