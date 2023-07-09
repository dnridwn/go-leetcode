// https://leetcode.com/problems/sum-of-unique-elements/

package sum_of_unique_elements

func sumOfUnique(nums []int) int {
	seen := make(map[int]bool, len(nums))
	for _, num := range nums {
		v, presence := seen[num]
		if !presence {
			seen[num] = true
			continue
		}

		if v {
			seen[num] = false
		}
	}

	sum := 0
	for k, v := range seen {
		if !v {
			continue
		}

		sum += k
	}

	return sum
}
