// https://leetcode.com/problems/search-insert-position/

package search_insert_position

func searchInsert(nums []int, target int) int {
	for k, v := range nums {
		if v == target || v > target {
			return k
		}
	}
	return len(nums)
}
