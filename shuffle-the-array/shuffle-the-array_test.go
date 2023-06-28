package shuffle_the_array

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	testCases := []struct {
		nums     []int
		n        int
		expected []int
	}{
		{
			nums:     []int{2, 5, 1, 3, 4, 7},
			n:        3,
			expected: []int{2, 3, 5, 4, 1, 7},
		},
		{
			nums:     []int{1, 2, 3, 4, 4, 3, 2, 1},
			n:        4,
			expected: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			nums:     []int{1, 1, 2, 2},
			n:        2,
			expected: []int{1, 2, 1, 2},
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, shuffle(tt.nums, tt.n))
		})
	}
}
