// https://leetcode.com/problems/merge-strings-alternately/

package merge_strings_alternately

import "math"

func mergeAlternately(word1 string, word2 string) string {
	var (
		result        string
		maxWordLength int = int(math.Max(float64(len(word1)), float64(len(word2))))
	)

	for i := 0; i < maxWordLength; i++ {
		if i < len(word1) {
			result += string(word1[i])
		}

		if i < len(word2) {
			result += string(word2[i])
		}
	}

	return result
}
