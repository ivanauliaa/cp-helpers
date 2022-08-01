package golang

import (
	"math"
)

func getMinMax(arr *[]int32) (int32, int32) {
	min, max := int32(math.MaxInt32), int32(0)

	for _, val := range *arr {
		if val > max {
			max = val
		}

		if val < min {
			min = val
		}
	}

	return min, max
}
