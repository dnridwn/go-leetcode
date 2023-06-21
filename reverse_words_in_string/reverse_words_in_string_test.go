package reverse_words_in_string

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWordsInString(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{
			s:        "the sky is blue",
			expected: "blue is sky the",
		},
		{
			s:        "  hello world  ",
			expected: "world hello",
		},
		{
			s:        "a good   example",
			expected: "example good a",
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, reverseWords(tt.s))
		})
	}
}
