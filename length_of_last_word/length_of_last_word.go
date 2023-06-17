// https://leetcode.com/problems/length-of-last-word/

package length_of_last_word

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Split(strings.Trim(s, " "), " ")
	return len(words[len(words)-1])
}
