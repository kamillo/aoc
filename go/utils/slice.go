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

func SumInts(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Product(array []interface{}) int {
	result := 0
	for _, v := range array {
		if i, err := strconv.Atoi(v.(string)); err == nil {
			if result != 0 {
				result *= i
			} else {
				result = i
			}
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

func AnyToInt(in []interface{}) []int {
	result := make([]int, len(in))
	for x, v := range in {
		if i, err := strconv.Atoi(v.(string)); err == nil {
			result[x] = i
		}
	}

	return result
}

func StringsToInts(in []string) []int {
	result := make([]int, len(in))
	for x, v := range in {
		if i, err := strconv.Atoi(v); err == nil {
			result[x] = i
		}
	}

	return result
}

func AnyToString(in []interface{}) []string {
	result := make([]string, len(in))
	for x, v := range in {
		result[x] = v.(string)
	}

	return result
}

func SliceRotateString(image []string) []string {
	size := len(image)
	newImage := make([][]byte, size)

	for y := 0; y < size; y++ {
		newImage[y] = make([]byte, size)
	}

	for y := range image {
		for x := range image[y] {
			newImage[size-x-1][y] = image[y][x]
		}
	}

	ret := []string{}
	for y := range newImage {
		ret = append(ret, string(newImage[y]))
	}
	return ret
}

func SliceFlipHString(image []string) []string {
	size := len(image)
	newImage := make([]string, size)

	for y := range image {
		line := make([]byte, size)
		for x := range image[y] {
			line[size-x-1] = image[y][x]
		}
		newImage[y] = string(line)
	}

	return newImage
}

func SliceFlipVString(image []string) []string {
	newImage := make([]string, len(image))
	copy(newImage, image)

	for i, j := 0, len(newImage)-1; i < j; i, j = i+1, j-1 {
		newImage[i], newImage[j] = newImage[j], newImage[i]
	}

	return newImage
}

func DeleteAtIndex[T any](slice []T, index int) []T {
  return append(slice[:index], slice[index+1:]...)
}
