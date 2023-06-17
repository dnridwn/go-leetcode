package first_unique_character_in_a_string

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqueCharacterInAString(t *testing.T) {
	testCases := []struct {
		words    string
		expected int
	}{
		{
			words:    "leetcode",
			expected: 0,
		},
		{
			words:    "loveleetcode",
			expected: 2,
		},
		{
			words:    "aabb",
			expected: -1,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, firstUniqChar(tt.words))
		})
	}
}
