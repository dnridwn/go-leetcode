// https://leetcode.com/problems/power-of-four

package power_of_four

func isPowerOfFour(n int) bool {
	result := 1
	for result < n {
		result *= 4
	}

	return result == n
}
