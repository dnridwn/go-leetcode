package valid_palindrome

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tc := []struct {
		s        string
		expected bool
	}{
		{
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			s:        "race a car",
			expected: false,
		},
		{
			s:        "ab_a",
			expected: true,
		},
		{
			s:        "0P",
			expected: false,
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, isPalindrome(tt.s))
		})
	}
}
