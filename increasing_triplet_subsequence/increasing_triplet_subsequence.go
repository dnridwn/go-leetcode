package increasing_triplet_subsequence

import "math"

func increasingTriplet(nums []int) bool {
	x := nums[0]
	y := math.MaxInt32

	for i := 1; i < len(nums); i++ {
		if x >= nums[i] {
			x = nums[i]
		} else if y >= nums[i] {
			y = nums[i]
		} else {
			return true
		}
	}

	return false
}
