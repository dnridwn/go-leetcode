// https://leetcode.com/problems/shuffle-the-array/

package shuffle_the_array

func shuffle(nums []int, n int) []int {
	shuffledNums := make([]int, 0)
	for i := 0; i < n; i++ {
		shuffledNums = append(shuffledNums, nums[i], nums[n+i])
	}
	return shuffledNums
}
