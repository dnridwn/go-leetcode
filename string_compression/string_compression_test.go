// https://leetcode.com/problems/string-compression/

package string_compression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	testCases := []struct {
		name          string
		chars         []byte
		expected      int
		expectedBytes []byte
	}{
		{
			name:          "",
			chars:         []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected:      6,
			expectedBytes: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			name:          "",
			chars:         []byte{'a'},
			expected:      1,
			expectedBytes: []byte{'a'},
		},
		{
			name:          "",
			chars:         []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected:      4,
			expectedBytes: []byte{'a', 'b', '1', '2'},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, compress(tc.chars))
		})
	}
}
