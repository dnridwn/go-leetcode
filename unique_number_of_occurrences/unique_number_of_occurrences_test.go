package unique_number_of_occurrences

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected bool
	}{
		{
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			arr:      []int{1, 2},
			expected: false,
		},
		{
			arr:      []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expected: true,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, uniqueOccurrences(tc.arr))
		})
	}
}
