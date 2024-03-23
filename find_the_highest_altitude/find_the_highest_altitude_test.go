package find_the_highest_altitude

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestAltitude(t *testing.T) {
	tc := []struct {
		gain     []int
		expected int
	}{
		{
			gain:     []int{-5, 1, 5, 0, -7},
			expected: 1,
		},
		{
			gain:     []int{-4, -3, -2, -1, 4, 3, 2},
			expected: 0,
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, largestAltitude(tt.gain))
		})
	}
}
