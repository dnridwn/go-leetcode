package add_strings

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStrings(t *testing.T) {
	testCases := []struct {
		num1     string
		num2     string
		expected string
	}{
		{
			num1:     "11",
			num2:     "123",
			expected: "134",
		},
		{
			num1:     "456",
			num2:     "77",
			expected: "533",
		},
		{
			num1:     "0",
			num2:     "0",
			expected: "0",
		},
		{
			num1:     "3876620623801494171",
			num2:     "6529364523802684779",
			expected: "10405985147604178950",
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, addStrings(tt.num1, tt.num2))
		})
	}
}
