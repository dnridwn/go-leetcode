package add_digits

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDigits(t *testing.T) {
	testCases := []struct {
		num      int
		expected int
	}{
		{
			num:      38,
			expected: 2,
		},
		{
			num:      0,
			expected: 0,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, addDigits(tt.num))
		})
	}
}
