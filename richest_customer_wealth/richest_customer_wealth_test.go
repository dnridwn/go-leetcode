package richest_customer_wealth

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumWealth(t *testing.T) {
	testCases := []struct {
		accounts [][]int
		expected int
	}{
		{
			accounts: [][]int{{1, 2, 3}, {3, 2, 1}},
			expected: 6,
		},
		{
			accounts: [][]int{{1, 5}, {7, 3}, {3, 5}},
			expected: 10,
		},
		{
			accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			expected: 17,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, maximumWealth(tc.accounts))
		})
	}
}
