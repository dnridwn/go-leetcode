package contains_duplicate

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, containsDuplicate(tt.nums))
		})
	}
}

func BenchmarkContainsDuplicate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		testCases := []struct {
			nums     []int
			expected bool
		}{
			{
				nums:     []int{1, 2, 3, 1},
				expected: true,
			},
			{
				nums:     []int{1, 2, 3, 4},
				expected: false,
			},
			{
				nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
				expected: true,
			},
		}

		for k, tt := range testCases {
			t.Run(strconv.Itoa(k), func(t *testing.B) {
				assert.Equal(t, tt.expected, containsDuplicate(tt.nums))
			})
		}
	}
}
