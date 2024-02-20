package decode_the_message

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesAndSpaces(t *testing.T) {
	tc := []struct {
		chars    []rune
		expected []rune
	}{
		{
			chars:    []rune{'t', 'h', 'e', ' ', 'q', 'u', 'i', 'c', 'k', ' ', 'b', 'r', 'o', 'w', 'n', ' ', 'f', 'o', 'x', ' ', 'j', 'u', 'm', 'p', 's', ' ', 'o', 'v', 'e', 'r', ' ', 't', 'h', 'e', ' ', 'l', 'a', 'z', 'y', ' ', 'd', 'o', 'g'},
			expected: []rune{'t', 'h', 'e', 'q', 'u', 'i', 'c', 'k', 'b', 'r', 'o', 'w', 'n', 'f', 'x', 'j', 'm', 'p', 's', 'v', 'l', 'a', 'z', 'y', 'd', 'g'},
		},
		{
			chars:    []rune{'e', 'l', 'j', 'u', 'x', 'h', 'p', 'w', 'n', 'y', 'r', 'd', 'g', 't', 'q', 'k', 'v', 'i', 's', 'z', 'c', 'f', 'm', 'a', 'b', 'o'},
			expected: []rune{'e', 'l', 'j', 'u', 'x', 'h', 'p', 'w', 'n', 'y', 'r', 'd', 'g', 't', 'q', 'k', 'v', 'i', 's', 'z', 'c', 'f', 'm', 'a', 'b', 'o'},
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, removeDuplicatesAndSpaces(tt.chars))
		})
	}
}

func TestMapCharsToKey(t *testing.T) {
	tc := []struct {
		chars    []rune
		expected map[rune]int
	}{
		{
			chars:    []rune{'a', 'b', 'c'},
			expected: map[rune]int{'a': 0, 'b': 1, 'c': 2},
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, mapCharsToKey(tt.chars))
		})
	}
}

func TestDecodeMessage(t *testing.T) {
	tc := []struct {
		key      string
		message  string
		expected string
	}{
		{
			key:      "the quick brown fox jumps over the lazy dog",
			message:  "vkbs bs t suepuv",
			expected: "this is a secret",
		},
		{
			key:      "eljuxhpwnyrdgtqkviszcfmabo",
			message:  "zwx hnfx lqantp mnoeius ycgk vcnjrdb",
			expected: "the five boxing wizards jump quickly",
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, decodeMessage(tt.key, tt.message))
		})
	}
}
