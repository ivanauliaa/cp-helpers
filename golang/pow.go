package golang

import "math"

func pow(base, power int64) int64 {
	return int64(math.Pow(float64(base), float64(power))) // don't forget to import "math" package
}
