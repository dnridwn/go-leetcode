// https://leetcode.com/problems/text-justification/

package text_justification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	testCases := []struct {
		name     string
		words    []string
		maxWidth int
		expected []string
	}{
		{
			name:     "Test Full Justify #1",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			expected: []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			name:     "Test Full Justify #2",
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			expected: []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, fullJustify(tc.words, tc.maxWidth))
		})
	}
}
