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
