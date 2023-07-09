package missing_number

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 0, 1},
			expected: 2,
		},
		{
			nums:     []int{0, 1},
			expected: 2,
		},
		{
			nums:     []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, missingNumber(tt.nums))
		})
	}
}
