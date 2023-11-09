// https://leetcode.com/problems/merge-strings-alternately/

package merge_strings_alternately

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		name     string
		w1       string
		w2       string
		expected string
	}{
		{
			name:     "Test Merge Alternately #1",
			w1:       "abc",
			w2:       "pqr",
			expected: "apbqcr",
		},
		{
			name:     "Test Merge Alternately #2",
			w1:       "ab",
			w2:       "pqrs",
			expected: "apbqrs",
		},
		{
			name:     "Test Merge Alternately #3",
			w1:       "abcd",
			w2:       "pq",
			expected: "apbqcd",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, mergeAlternately(tc.w1, tc.w2))
		})
	}
}
