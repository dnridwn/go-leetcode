package single_number_iii

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumberII(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 1, 3, 2, 5},
			expected: []int{3, 5},
		},
		{
			nums:     []int{-1, 0},
			expected: []int{-1, 0},
		},
		{
			nums:     []int{0, 1},
			expected: []int{0, 1},
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, singleNumber(tt.nums))
		})
	}
}
