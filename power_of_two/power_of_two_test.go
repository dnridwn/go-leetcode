package power_of_two

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{
			n:        1,
			expected: true,
		},
		{
			n:        16,
			expected: true,
		},
		{
			n:        3,
			expected: false,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, isPowerOfTwo(tt.n))
		})
	}
}
