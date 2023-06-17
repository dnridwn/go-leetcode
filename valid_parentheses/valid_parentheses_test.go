package valid_parentheses

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {
	testCases := []struct {
		strings  string
		expected bool
	}{
		{
			strings:  "()",
			expected: true,
		},
		{
			strings:  "()[]{}",
			expected: true,
		},
		{
			strings:  "(]",
			expected: false,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, isValid(tt.strings))
		})
	}
}
