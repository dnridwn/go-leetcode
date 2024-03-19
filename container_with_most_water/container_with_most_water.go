package container_with_most_water

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		h := 0
		if height[left] > height[right] {
			h = height[right]
		} else {
			h = height[left]
		}

		w := right - left
		area := h * w

		if area > maxArea {
			maxArea = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}
