package majority_element

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, majorityElement(tt.nums))
		})
	}
}
