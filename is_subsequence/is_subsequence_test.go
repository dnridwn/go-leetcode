package is_subsequence

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tc := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, isSubsequence(tt.s, tt.t))
		})
	}
}
