package equal_row_and_column_pairs

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualPairs(t *testing.T) {
	testCases := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid:     [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
			expected: 1,
		},
		{
			grid:     [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
			expected: 3,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, equalPairs(tc.grid))
		})
	}
}
