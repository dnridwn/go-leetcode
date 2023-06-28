package sorting_the_sentence

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSentence(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{
			s:        "is2 sentence4 This1 a3",
			expected: "This is a sentence",
		},
		{
			s:        "Myself2 Me1 I4 and3",
			expected: "Me Myself and I",
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, sortSentence(tt.s))
		})
	}
}
