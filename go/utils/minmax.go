package utils

// Min - Find minimal value from argumet variables
func Min(vn ...int) (ret int) {
	for i, value := range vn {
		if i == 0 || value < ret {
			ret = value
		}
	}

	return
}

// Max - Find minimal value from argumet variables
func Max(vn ...int) (ret int) {
	for i, value := range vn {
		if i == 0 || value > ret {
			ret = value
		}
	}

	return
}

// MinInArray - Find minimal value from array
func MinInArray(v []int) (ret int) {
	for i, value := range v {
		if i == 0 || value < ret {
			ret = value
		}
	}

	return
}

// MaxInArray - Find minimal value from array
func MaxInArray(v []int) (ret int) {
	for i, value := range v {
		if i == 0 || value > ret {
			ret = value
		}
	}

	return
}

type CompareFunc func(interface{}, interface{}) bool

// MaxInAnyArray - Find maximal value from array
func MaxInAnyArray(v []interface{}, fn CompareFunc) (ret interface{}) {
	for i, value := range v {
		if i == 0 || fn(value, ret) {
			ret = value
		}
	}

	return
}

// Gcd - Greatest common divisor
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Lcm - Least common multiple
func Lcm(a, b int, integers ...int) int {
	result := a * b / Gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}

	return result
}
