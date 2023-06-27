package convert_the_temperature

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertTemperature(t *testing.T) {
	testCases := []struct {
		celcius  float64
		expected []float64
	}{
		{
			celcius:  36.50,
			expected: []float64{309.65000, 97.70000},
		},
		{
			celcius:  122.11,
			expected: []float64{395.26000, 251.79800},
		},
	}

	for k, tt := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, convertTemperature(tt.celcius))
		})
	}
}
