// https://leetcode.com/problems/sorting-the-sentence

package sorting_the_sentence

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	mapper := make(map[string]string)
	for _, word := range strings.Split(s, " ") {
		n := word[len(word)-1:]
		word = word[:len(word)-1]
		mapper[n] = word
	}

	result := make([]string, 0)
	for i := 1; i <= len(mapper); i++ {
		result = append(result, mapper[strconv.Itoa(i)])
	}

	return strings.Join(result, " ")
}
