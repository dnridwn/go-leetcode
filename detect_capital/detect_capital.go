// https://leetcode.com/problems/detect-capital

package detect_capital

import (
	"strings"
)

func detectCapitalUse(word string) bool {
	return (strings.ToUpper(word) == word) ||
		(strings.ToLower(word) == word) ||
		(strings.Title(strings.ToLower(word)) == word)
}
