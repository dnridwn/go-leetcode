package search_2d_matrix_ii

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch2DMatrixII(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			matrix:   [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target:   5,
			expected: true,
		},
		{
			matrix:   [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target:   20,
			expected: false,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, searchMatrix(tt.matrix, tt.target))
		})
	}
}
