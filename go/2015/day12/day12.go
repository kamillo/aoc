package main

import (
	"encoding/json"
	"fmt"
	"github.com/kamillo/aoc/utils"
	"io"
	"strings"
)

func main() {
	lines := utils.GetLines("input.txt")

	fmt.Println("Part 1: ", sumNumbers(lines[0]))

	b := []byte(lines[0])
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	_, m = removeRedFromMap(m)
	//fmt.Println(m)
	j, _ := json.Marshal(m)
	fmt.Println(string(j))
	fmt.Println("Part 2: ", sumNumbers(string(j)))
}

func removeRedFromMap(m map[string]interface{}) (bool, map[string]interface{}) {
	for k, v := range m {
		switch v.(type) {
		case []interface{}:
			m[k] = removeRedFromArray(v.([]interface{}))
		case map[string]interface{}:
			if ok, _ := removeRedFromMap(v.(map[string]interface{})); ok {
				delete(m, k)
			}
		case string:
			if v.(string) == "red" {
				return true, m
				//delete(m, k)
			}
		}
	}

	return false, m
}

func removeRedFromArray(m []interface{}) []interface{} {
	z := make([]interface{}, 0)
	for _, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			if ok, _ := removeRedFromMap(v.(map[string]interface{})); !ok {
				z = append(z, v)
			}
		case []interface{}:
			z = append(z, removeRedFromArray(v.([]interface{})))
		default:
			z = append(z, v)
		}
	}

	return z
}

func sumNumbers(jsonString string) int {
	dec := json.NewDecoder(strings.NewReader(jsonString))
	sum := float64(0)
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		switch t.(type) {
		case float64:
			sum += t.(float64)
		default:
			break
		}
	}

	return int(sum)
}
