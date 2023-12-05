package utils

import (
	"strconv"
	"strings"
)

func SliceRotateBool(image [][]bool) [][]bool {
	size := len(image)
	newImage := make([][]bool, size)

	for y := 0; y < size; y++ {
		newImage[y] = make([]bool, size)
	}

	for y := range image {
		for x := range image[y] {
			newImage[size-x-1][y] = image[y][x]
		}
	}

	return newImage
}

func SliceFlipBool(image [][]bool) [][]bool {
	size := len(image)
	newImage := make([][]bool, size)

	for i := 0; i < size; i++ {
		newImage[i] = make([]bool, size)
	}

	for y := range image {
		for x := range image[y] {
			newImage[y][size-x-1] = image[y][x]
		}
	}

	return newImage
}

func SliceToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}

func Sum(array []interface{}) int {
	result := 0
	for _, v := range array {
		if i, err := strconv.Atoi(v.(string)); err == nil {
			result += i
		}
	}
	return result
}

func JoinInts(elems []int, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(strconv.Itoa(elems[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.Itoa(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(strconv.Itoa(s))
	}
	return b.String()
}
