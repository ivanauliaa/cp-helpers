package golang

import "math"

func minInt(arr *[]int32) int32 {
	min := int32(math.MaxInt32) // don't forget to import "math" package

	for _, val := range *arr {
		if val < min {
			min = val
		}
	}

	return min
}
