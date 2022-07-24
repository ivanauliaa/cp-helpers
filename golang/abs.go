package golang

func abs(num int32) int32 {
	if num < 0 {
		return num * -1
	}

	return num
}
