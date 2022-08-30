package sort

import (
	"math/rand"
	"time"
)

func quickSortInt(a []int32) []int32 {
	rand.Seed(time.Now().UnixNano()) // don't forget to import "math/rand" and "time" package

	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSortInt(a[:left])
	quickSortInt(a[left+1:])

	return a
}

func quickSortByte(a []byte) []byte {
	rand.Seed(time.Now().UnixNano()) // don't forget to import "math/rand" and "time" package

	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSortByte(a[:left])
	quickSortByte(a[left+1:])

	return a
}
