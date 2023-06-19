package my_pow

import (
	"math"
)

func myPow(x float64, n int) float64 {
	return math.Round((math.Pow(x, float64(n)) * 1e5)) / 1e5
}
