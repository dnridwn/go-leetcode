package palindrome_number

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeNumber(t *testing.T) {
	testCases := []struct {
		number   int
		expected bool
	}{
		{
			number:   121,
			expected: true,
		},
		{
			number:   -121,
			expected: false,
		},
		{
			number:   10,
			expected: false,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, isPalindrome(tt.number))
		})
	}
}
