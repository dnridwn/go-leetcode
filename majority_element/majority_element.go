// https://leetcode.com/problems/majority-element/

package majority_element

func majorityElement(nums []int) int {
	uniqueNums := unique(nums)
	highestCount := 0
	highestElem := 0
	for _, v := range uniqueNums {
		count := count(nums, v)
		if count > highestCount {
			highestCount = count
			highestElem = v
		}
	}

	return highestElem
}

func unique(nums []int) []int {
	uniqueNums := []int{}
	for _, v := range nums {
		if count(uniqueNums, v) == 0 {
			uniqueNums = append(uniqueNums, v)
		}
	}
	return uniqueNums
}

func count(nums []int, n int) int {
	total := 0
	for _, v := range nums {
		if v == n {
			total += 1
		}
	}
	return total
}
