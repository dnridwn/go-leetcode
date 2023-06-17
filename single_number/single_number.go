// https://leetcode.com/problems/single-number/

package single_number

func singleNumber(nums []int) int {
	result := 0
	for i := range nums {
		result = result ^ nums[i]
	}
	return result
}
