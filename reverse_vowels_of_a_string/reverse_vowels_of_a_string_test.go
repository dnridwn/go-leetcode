package reverse_vowels_of_a_string

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tc := []struct {
		s        string
		expected string
	}{
		{
			s:        "hello",
			expected: "holle",
		},
		{
			s:        "leetcode",
			expected: "leotcede",
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, reverseVowels(tt.s))
		})
	}
}
