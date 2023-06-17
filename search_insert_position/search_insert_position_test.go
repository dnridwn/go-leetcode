package search_insert_position

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsertPosition(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, searchInsert(tt.nums, tt.target))
		})
	}
}
