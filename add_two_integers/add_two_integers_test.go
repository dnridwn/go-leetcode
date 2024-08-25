package add_two_integers

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		num1     int
		num2     int
		expected int
	}{
		{
			num1:     12,
			num2:     5,
			expected: 17,
		},
		{
			num1:     -10,
			num2:     4,
			expected: -6,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, sum(tc.num1, tc.num2))
		})
	}
}
