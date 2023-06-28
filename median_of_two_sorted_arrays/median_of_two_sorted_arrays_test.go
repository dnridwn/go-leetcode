package median_of_two_sorted_arrays

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: float64(2),
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: float64(2.5),
		},
		{
			nums1:    []int{2, 2, 4, 4},
			nums2:    []int{2, 2, 4, 4},
			expected: float64(3),
		},
		{
			nums1:    []int{},
			nums2:    []int{1, 2, 3, 4, 5},
			expected: float64(3),
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, findMedianSortedArrays(tt.nums1, tt.nums2))
		})
	}
}
