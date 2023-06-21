// https://leetcode.com/problems/reverse-words-in-a-string

package reverse_words_in_string

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	re := regexp.MustCompile(`\s+`)
	s = strings.Trim(s, " ")
	words := re.Split(s, -1)
	newWords := []string{}
	for _, w := range words {
		newWords = append([]string{strings.Trim(w, " ")}, newWords...)
	}
	return strings.Join(newWords, " ")
}
