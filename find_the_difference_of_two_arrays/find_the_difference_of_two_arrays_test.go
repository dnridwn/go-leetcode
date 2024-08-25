package find_the_difference_of_two_arrays

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDifference(t *testing.T) {
	testCases := []struct {
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{
			nums1:    []int{1, 2, 3},
			nums2:    []int{2, 4, 6},
			expected: [][]int{{1, 3}, {4, 6}},
		},
		{
			nums1:    []int{1, 2, 3, 3},
			nums2:    []int{1, 1, 2, 2},
			expected: [][]int{{3}, []int(nil)},
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, findDifference(tc.nums1, tc.nums2))
		})
	}
}
