package multiply_strings

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	testCases := []struct {
		num1     string
		num2     string
		expected string
	}{
		{
			num1:     "2",
			num2:     "3",
			expected: "6",
		},
		{
			num1:     "123",
			num2:     "456",
			expected: "56088",
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, multiply(tt.num1, tt.num2))
		})
	}
}
