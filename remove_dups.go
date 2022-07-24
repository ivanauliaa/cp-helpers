package golang

func removeIntArrDups(arr *[]int32) []int32 {
	mapper := map[int32]bool{}
	new := []int32{}

	for _, a := range *arr {
		if !mapper[a] {
			mapper[a] = true
			new = append(new, a)
		}
	}

	return new
}
