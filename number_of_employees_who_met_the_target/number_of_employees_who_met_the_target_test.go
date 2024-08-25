package number_of_employees_who_met_the_target

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfEmployeesWhoMetTarget(t *testing.T) {
	testCases := []struct {
		hours    []int
		target   int
		expected int
	}{
		{
			hours:    []int{0, 1, 2, 3, 4},
			target:   2,
			expected: 3,
		},
		{
			hours:    []int{5, 1, 4, 2, 2},
			target:   6,
			expected: 0,
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, numberOfEmployeesWhoMetTarget(tc.hours, tc.target))
		})
	}
}
