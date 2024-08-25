package jewels_and_stones

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumJewelsInStones(t *testing.T) {
	testCases := []struct {
		jewels   string
		stones   string
		expected int
	}{
		{
			jewels:   "aA",
			stones:   "aAAbbbb",
			expected: 3,
		},
		{
			jewels:   "z",
			stones:   "ZZ",
			expected: 0,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, numJewelsInStones(tc.jewels, tc.stones))
		})
	}
}
