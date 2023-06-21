// https://leetcode.com/problems/single-number-iii/

package single_number_iii

func singleNumber(nums []int) []int {
	seenCount := make(map[int]int)
	for _, n := range nums {
		seenCount[n] += 1
	}

	result := []int{}
	for k, v := range seenCount {
		if v == 1 {
			result = append(result, k)
		}
	}

	return result
}
