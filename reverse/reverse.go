package reverse

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	absX := int(math.Abs(float64(x)))
	absXInStr := strconv.Itoa(absX)
	newAbsXinStr := ""

	for _, xx := range absXInStr {
		newAbsXinStr = string(xx) + newAbsXinStr
	}

	if x < 0 {
		newAbsXinStr = "-" + newAbsXinStr
	}

	x64, _ := strconv.ParseInt(newAbsXinStr, 10, 64)

	if x64 > math.MaxInt32 || x64 < math.MinInt32 {
		return 0
	}

	return int(x64)
}
