package length_of_last_word

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		words    string
		expected int
	}{
		{
			words:    "Hello World",
			expected: 5,
		},
		{
			words:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			words:    "luffy is still joyboy",
			expected: 6,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, lengthOfLastWord(tt.words))
		})
	}
}
