package power_of_four

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{
			n:        16,
			expected: true,
		},
		{
			n:        5,
			expected: false,
		},
		{
			n:        1,
			expected: true,
		},
		{
			n:        0,
			expected: false,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, isPowerOfFour(tt.n))
		})
	}
}
