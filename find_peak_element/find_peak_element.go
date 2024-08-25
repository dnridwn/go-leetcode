package find_peak_element

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	c := -1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[i] > nums[i+1] {
				c = i
			}
		} else if i == len(nums)-1 {
			if nums[i] > nums[i-1] {
				c = i
			}
		} else {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				c = i
			}
		}
	}

	return c
}
