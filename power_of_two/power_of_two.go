// https://leetcode.com/problems/power-of-two/submissions/978374936/

package power_of_two

func isPowerOfTwo(n int) bool {
	result := 1
	for result < n {
		result *= 2
	}

	return result == n
}
