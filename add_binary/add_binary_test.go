package add_binary

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBn(t *testing.T) {
	testCases := []struct {
		a        string
		b        string
		expected string
	}{
		{
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			a:        "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			b:        "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			expected: "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, addBinary(tt.a, tt.b))
		})
	}
}
