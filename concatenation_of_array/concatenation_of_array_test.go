package concatenation_of_array

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConcatenation(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 1},
			expected: []int{1, 2, 1, 1, 2, 1},
		},
		{
			nums:     []int{1, 3, 2, 1},
			expected: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, getConcatenation(tt.nums))
		})
	}
}
