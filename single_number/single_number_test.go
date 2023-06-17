package single_number

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, singleNumber(tt.nums))
		})
	}
}
