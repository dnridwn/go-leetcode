package fizz_buzz

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		n        int
		expected []string
	}{
		{
			n:        3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			n:        5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			n:        15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, fizzBuzz(tt.n))
		})
	}
}
