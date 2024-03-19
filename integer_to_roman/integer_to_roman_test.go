package integer_to_roman

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
	}{
		{
			num:      3,
			expected: "III",
		},
		{
			num:      58,
			expected: "LVIII",
		},
		{
			num:      1994,
			expected: "MCMXCIV",
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, intToRoman(tt.num))
		})
	}
}
