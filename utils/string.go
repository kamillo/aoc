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
