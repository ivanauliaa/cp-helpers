package max

func maxInt(arr *[]int32) int32 {
	max := int32(0)

	for _, val := range *arr {
		if val > max {
			max = val
		}
	}

	return max
}
