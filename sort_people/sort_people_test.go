package sort_people

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortPeople(t *testing.T) {
	testCases := []struct {
		names    []string
		heights  []int
		expected []string
	}{
		{
			names:    []string{"Mary", "John", "Emma"},
			heights:  []int{180, 165, 170},
			expected: []string{"Mary", "Emma", "John"},
		},
		{
			names:    []string{"Alice", "Bob", "Bob"},
			heights:  []int{155, 185, 150},
			expected: []string{"Bob", "Alice", "Bob"},
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, sortPeople(tt.names, tt.heights))
		})
	}
}
