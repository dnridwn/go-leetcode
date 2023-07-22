// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array

package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	positions := []int{-1, -1}
	for k, n := range nums {
		if n == target {
			if positions[0] == -1 {
				if k > positions[0] {
					positions[0] = k
				}
			}

			if k > positions[1] {
				positions[1] = k
			}
		}
	}

	return positions
}
