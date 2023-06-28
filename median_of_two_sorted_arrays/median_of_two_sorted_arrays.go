// https://leetcode.com/problems/median-of-two-sorted-arrays

package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge both slices
	mergedNums := append(nums1, nums2...)

	if len(mergedNums) == 1 {
		return float64(mergedNums[0])
	}

	// sort
	for i := 0; i < len(mergedNums)-1; i++ {
		for j := i + 1; j < len(mergedNums); j++ {
			if mergedNums[i] > mergedNums[j] {
				mergedNums[i], mergedNums[j] = mergedNums[j], mergedNums[i]
			}
		}
	}

	// find median
	middleKey := len(mergedNums) / 2
	if len(mergedNums)%2 == 0 {
		return float64((mergedNums[middleKey-1] + mergedNums[middleKey])) / float64(2)
	}

	return float64(mergedNums[middleKey])
}
