// https://leetcode.com/problems/single-number-ii/

package single_number_ii

func singleNumber(nums []int) int {
	seenCount := make(map[int]int)
	for _, n := range nums {
		seenCount[n] += 1
	}

	for k, v := range seenCount {
		if v == 1 {
			return k
		}
	}

	return -1
}
