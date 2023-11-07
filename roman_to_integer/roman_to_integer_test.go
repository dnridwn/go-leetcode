// https://leetcode.com/problems/roman-to-integer/

package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Test Roman to Integer #1",
			s:        "III",
			expected: 3,
		},
		{
			name:     "Test Roman to Integer #2",
			s:        "LVIII",
			expected: 58,
		},
		{
			name:     "Test Roman to Integer #3",
			s:        "MCMXCIV",
			expected: 1994,
		},
		{
			name:     "Test Roman to Integer #4",
			s:        "IX",
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, romanToInt(tc.s))
		})
	}
}
