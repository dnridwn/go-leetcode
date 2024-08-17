package maximum_average_subarray_i

import "math"

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxAverage := float64(sum) / float64(k)

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		maxAverage = math.Max(maxAverage, float64(sum)/float64(k))
	}

	return maxAverage
}
