package find_the_difference_of_two_arrays

func findDifference(nums1 []int, nums2 []int) [][]int {
	var (
		nm1, nm2       map[int]bool = make(map[int]bool), make(map[int]bool)
		cNums1, cNums2 []int
		un1, un2       []int
	)

	for _, n := range nums1 {
		if _, ok := nm1[n]; ok {
			continue
		}

		nm1[n] = true
		cNums1 = append(cNums1, n)
	}

	for _, n := range nums2 {
		if _, ok := nm2[n]; ok {
			continue
		}

		nm2[n] = true
		cNums2 = append(cNums2, n)
	}

	for _, n := range cNums1 {
		if _, ok := nm2[n]; !ok {
			un1 = append(un1, n)
		}
	}

	for _, n := range cNums2 {
		if _, ok := nm1[n]; !ok {
			un2 = append(un2, n)
		}
	}

	return [][]int{un1, un2}
}
