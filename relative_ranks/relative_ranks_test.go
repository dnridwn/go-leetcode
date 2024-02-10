package relative_ranks

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRelativeRanks(t *testing.T) {
	tc := []struct {
		score    []int
		expected []string
	}{
		{
			score:    []int{5, 4, 3, 2, 1},
			expected: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			score:    []int{10, 3, 8, 9, 4},
			expected: []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, findRelativeRanks(tt.score))
		})
	}
}
