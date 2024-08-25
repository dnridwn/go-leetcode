package find_words_containing_character

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWordsContaining(t *testing.T) {
	testCases := []struct {
		words    []string
		x        byte
		expected []int
	}{
		{
			words:    []string{"leet", "code"},
			x:        'e',
			expected: []int{0, 1},
		},
		{
			words:    []string{"abc", "bcd", "aaaa", "cbc"},
			x:        'a',
			expected: []int{0, 2},
		},
		{
			words:    []string{"abc", "bcd", "aaaa", "cbc"},
			x:        'z',
			expected: []int{},
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, findWordsContaining(tc.words, tc.x))
		})
	}
}
