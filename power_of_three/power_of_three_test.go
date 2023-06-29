package power_of_three

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerOfThree(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{
			n:        27,
			expected: true,
		},
		{
			n:        0,
			expected: false,
		},
		{
			n:        -1,
			expected: false,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, isPowerOfThree(tt.n))
		})
	}
}
