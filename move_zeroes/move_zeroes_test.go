package move_zeroes

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
		{
			nums:     []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			moveZeroes(tt.nums)
			assert.Equal(t, tt.expected, tt.nums)
		})
	}
}
