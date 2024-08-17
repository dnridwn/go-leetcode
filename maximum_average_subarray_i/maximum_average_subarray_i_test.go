package maximum_average_subarray_i

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75,
		},
		{
			nums:     []int{5},
			k:        1,
			expected: 5.0,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, findMaxAverage(tc.nums, tc.k))
		})
	}
}
