package mathutils

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
