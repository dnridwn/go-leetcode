// https://leetcode.com/problems/sqrtx/

package sqrt_x

func mySqrt(x int) int {
	lastResult := x
	times := 1
	for i := 1; i <= x; i++ {
		if i%2 == 1 {
			lastResult = lastResult - i
			if lastResult == 0 {
				return times
			} else if lastResult < 0 {
				return times - 1
			}
			times += 1
		}
		if i == x {
			if i%2 == 1 {
				return times
			} else {
				return times - 1
			}
		}
	}
	return 0
}
