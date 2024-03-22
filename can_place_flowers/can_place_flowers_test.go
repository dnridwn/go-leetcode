package can_place_flowers

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	tc := []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, canPlaceFlowers(tt.flowerbed, tt.n))
		})
	}
}
