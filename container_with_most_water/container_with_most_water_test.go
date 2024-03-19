package container_with_most_water

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	tc := []struct {
		heigth   []int
		expected int
	}{
		{
			heigth:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			heigth:   []int{1, 1},
			expected: 1,
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, maxArea(tt.heigth))
		})
	}
}
