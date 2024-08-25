package find_peak_element

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPeakElement(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 1},
			expected: 2,
		},
		{
			nums:     []int{1, 2, 1, 3, 5, 6, 4},
			expected: 5,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, findPeakElement(tc.nums))
		})
	}
}
