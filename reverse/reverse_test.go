package reverse

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		x        int
		expected int
	}{
		{
			x:        123,
			expected: 321,
		},
		{
			x:        -123,
			expected: -321,
		},
		{
			x:        120,
			expected: 21,
		},
		{
			x:        1534236469,
			expected: 0,
		},
		{
			x:        -2147483648,
			expected: 0,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, reverse(tt.x))
		})
	}
}
