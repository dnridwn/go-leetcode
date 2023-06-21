package single_number_ii

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumberII(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, 3, 2},
			expected: 3,
		},
		{
			nums:     []int{0, 1, 0, 1, 0, 1, 99},
			expected: 99,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, singleNumber(tt.nums))
		})
	}
}
