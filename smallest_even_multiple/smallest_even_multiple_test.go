package smallest_even_multiple

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestEvenMultiple(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{
			n:        5,
			expected: 10,
		},
		{
			n:        6,
			expected: 6,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, smallestEvenMultiple(tc.n))
		})
	}
}
