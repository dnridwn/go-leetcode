package flipping_an_image

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipAndInvertImage(t *testing.T) {
	testCases := []struct {
		image    [][]int
		expected [][]int
	}{
		{
			image:    [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			expected: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			image:    [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			expected: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, flipAndInvertImage(tt.image))
		})
	}
}
