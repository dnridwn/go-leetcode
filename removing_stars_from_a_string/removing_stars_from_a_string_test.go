package removing_stars_from_a_string

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveStars(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{
			s:        "leet**cod*e",
			expected: "lecoe",
		},
		{
			s:        "erase*****",
			expected: "",
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, removeStars(tc.s))
		})
	}
}
