package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	n := len(nums)

	l := make([]int, n)
	r := make([]int, n)
	answer := make([]int, n)

	l[0] = 1
	for i := 1; i < n; i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	r[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		answer[i] = l[i] * r[i]
	}

	return answer
}
