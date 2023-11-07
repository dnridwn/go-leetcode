// https://leetcode.com/problems/daily-temperatures/

package daily_temperatures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	testCases := []struct {
		name         string
		temperatures []int
		expected     []int
	}{
		{
			name:         "Test Daily Temperatures #1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "Test Daily Temperatures #2",
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		{
			name:         "Test Daily Temperatures #3",
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, dailyTemperatures(tc.temperatures))
		})
	}
}
