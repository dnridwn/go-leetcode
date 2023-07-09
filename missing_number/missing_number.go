// https://leetcode.com/problems/missing-number/

package missing_number

func missingNumber(nums []int) int {
	existence := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		existence[nums[i]] = true
	}

	for i := 0; i <= len(nums); i++ {
		if _, presence := existence[i]; !presence {
			return i
		}
	}

	return -1
}
