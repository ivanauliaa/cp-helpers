package golang

func sumBigInt(arr *[]int32) int64 {
	sum := int64(0)

	for _, val := range *arr {
		sum += int64(val)
	}

	return sum
}

func sumInt(arr *[]int32) int32 {
	sum := int32(0)

	for _, val := range *arr {
		sum += val
	}

	return sum
}
