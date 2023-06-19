package my_pow

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyPow(t *testing.T) {
	testCases := []struct {
		x        float64
		n        int
		expected float64
	}{
		{
			x:        2.00000,
			n:        10,
			expected: 1024.00000,
		},
		{
			x:        2.10000,
			n:        3,
			expected: 9.26100,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, myPow(tt.x, tt.n))
		})
	}
}
