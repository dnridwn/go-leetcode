package top_k_frequent_elements

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	tc := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.ElementsMatch(t, tt.expected, topKFrequent(tt.nums, tt.k))
		})
	}
}
