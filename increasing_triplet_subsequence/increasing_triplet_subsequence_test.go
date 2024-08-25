package increasing_triplet_subsequence

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTriplet(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			nums:     []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			nums:     []int{2, 1, 5, 0, 4, 6},
			expected: true,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, increasingTriplet(tc.nums))
		})
	}
}
