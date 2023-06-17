// https://leetcode.com/problems/two-sum/

package two_sum

func twoSum(nums []int, target int) []int {
	for i1, n1 := range nums {
		for i2, n2 := range nums {
			if i1 != i2 {
				if (n1 + n2) == target {
					return []int{i1, i2}
				}
			}
		}
	}
	return nil
}
