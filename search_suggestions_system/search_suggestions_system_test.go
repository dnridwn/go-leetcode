package search_suggestions_system

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuggestedProducts(t *testing.T) {
	tc := []struct {
		products   []string
		searchWord string
		expected   [][]string
	}{
		{
			products:   []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			searchWord: "mouse",
			expected:   [][]string{{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}},
		},
		{
			products:   []string{"havana"},
			searchWord: "havana",
			expected:   [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			assert.Equal(t, tt.expected, suggestedProducts(tt.products, tt.searchWord))
		})
	}
}
