package nth_tribonacci_number

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTribonnaci(t *testing.T) {
	tc := []struct {
		n        int
		expected int
	}{
		{
			n:        4,
			expected: 4,
		},
		{
			n:        25,
			expected: 1389537,
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, tribonacci(tt.n))
		})
	}
}
