package sum_of_unique_elements

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfUnique(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 2},
			expected: 4,
		},
		{
			nums:     []int{1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, sumOfUnique(tt.nums))
		})
	}
}
