package rotate

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	tc := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			matrix:   [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			expected: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, rotate(tt.matrix))
		})
	}
}
