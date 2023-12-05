package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	medicine := "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"
	//medicine := "HOHOHO"
	molecules := make(map[string]bool)

	for _, line := range utils.GetLines("input.txt") {
		split := strings.Split(line, " ")
		//strings.Replace(medicine, split[0], split[2], 1)
		index := 0
		for {
			i := strings.Index(medicine[index:], split[0])
			if i == -1 {
				break
			}
			index += i

			molecules[medicine[:index]+split[2]+medicine[index+len(split[0]):]] = true

			index += len(split[0])
			if index >= len(medicine) {
				break
			}
		}
	}
	// fmt.Println(molecules)
	fmt.Println("Part 1:", len(molecules))
	type Pair struct {
		a, b interface{}
	}
	transform := []Pair{}
	for _, line := range utils.GetLines("input.txt") {
		split := strings.Split(line, " ")
		transform = append(transform, Pair{a: split[0], b: split[2]})
	}

	cnt := 0
	for medicine != "e" {
		for _, v := range transform {
			if strings.Contains(medicine, v.b.(string)) {
				medicine = strings.Replace(medicine, v.b.(string), v.a.(string), 1)
				cnt++
			}
		}
	}

	fmt.Println("Part 2:", cnt)
}
