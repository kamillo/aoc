package utils

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
