package detect_capital

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCapital(t *testing.T) {
	testCases := []struct {
		word     string
		expected bool
	}{
		{
			word:     "USA",
			expected: true,
		},
		{
			word:     "FlaG",
			expected: false,
		},
		{
			word:     "test",
			expected: true,
		},
		{
			word:     "g",
			expected: true,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, detectCapitalUse(tt.word))
		})
	}
}
