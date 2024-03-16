package valid_palindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9\s]+|\s+`)
	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	sR := ""
	for i := len(s) - 1; i >= 0; i-- {
		sR += string(s[i])
	}

	return s == sR
}
