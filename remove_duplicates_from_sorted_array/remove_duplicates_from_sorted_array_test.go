// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package remove_duplicates_from_sorted_array

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tc := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 2},
			expected: 2,
		},
		{
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
		{
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			nums:     []int{1, 1, 1},
			expected: 1,
		},
		{
			nums:     []int{-1, 0, 0, 0, 0, 3, 3},
			expected: 3,
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, removeDuplicates(tt.nums))
		})
	}
}
