package defanging_an_ip_address

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefangIpAddr(t *testing.T) {
	testCases := []struct {
		address  string
		expected string
	}{
		{
			address:  "1.1.1.1",
			expected: "1[.]1[.]1[.]1",
		},
		{
			address:  "255.100.50.0",
			expected: "255[.]100[.]50[.]0",
		},
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tc.expected, defangIPaddr(tc.address))
		})
	}
}
