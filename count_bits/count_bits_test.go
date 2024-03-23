package count_bits

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	tc := []struct {
		n        int
		expected []int
	}{
		{
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, countBits(tt.n))
		})
	}
}
