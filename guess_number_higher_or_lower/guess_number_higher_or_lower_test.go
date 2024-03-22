package guess_number_higher_or_lower

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuessNumber(t *testing.T) {
	tc := []struct {
		n        int
		pick     int
		expected int
	}{
		{
			n:        10,
			pick:     6,
			expected: 6,
		},
		{
			n:        1,
			pick:     1,
			expected: 1,
		},
		{
			n:        2,
			pick:     1,
			expected: 1,
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			pickN = tt.pick
			assert.Equal(t, tt.expected, guessNumber(tt.n))
		})
	}
}
