package utils

import "math/bits"

func HeapPermutation(a []interface{}) [][]interface{} {
	var permutations [][]interface{}
	var generate func([]interface{}, int)

	generate = func(a []interface{}, size int) {
		if size == 1 {
			A := make([]interface{}, len(a))
			copy(A, a)
			permutations = append(permutations, A)
		}
		for i := 0; i < size; i++ {
			generate(a, size-1)
			if size%2 == 1 {
				a[0], a[size-1] = a[size-1], a[0]
			} else {
				a[i], a[size-1] = a[size-1], a[i]
			}
		}
	}
	generate(a, len(a))
	return permutations
}

func Combinations(set []interface{}, n int) (subsets [][]interface{}) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []interface{}

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

func CombinationsWithRepetitions(n int, lst []string) [][]string {
	if n == 0 {
		return [][]string{nil}
	}
	if len(lst) == 0 {
		return nil
	}
	r := CombinationsWithRepetitions(n, lst[1:])
	for _, x := range CombinationsWithRepetitions(n-1, lst) {
		r = append(r, append(x, lst[0]))
	}
	return r
}
