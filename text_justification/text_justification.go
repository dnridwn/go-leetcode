// https://leetcode.com/problems/text-justification/

package text_justification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var (
		groupedWords   [][]string = [][]string{}
		currGroupIndex int        = 0
		result         []string
	)

	for _, word := range words {
		if len(groupedWords) < currGroupIndex+1 {
			groupedWords = append(groupedWords, []string{})
		}

		tempCurrGroupWords := append(groupedWords[currGroupIndex], word)
		tempCurrGroupWordsLength := len(strings.Join(tempCurrGroupWords, " "))

		if tempCurrGroupWordsLength > maxWidth {
			currGroupIndex += 1
			groupedWords = append(groupedWords, []string{})
		}

		groupedWords[currGroupIndex] = append(groupedWords[currGroupIndex], word)
	}

	for k, wordsInLine := range groupedWords {
		if k+1 == len(groupedWords) {
			line := strings.Join(wordsInLine, " ")
			line += strings.Repeat(" ", maxWidth-len(line))
			result = append(result, line)
			continue
		}

		var (
			totalSpaceLength int = maxWidth - len(strings.Join(wordsInLine, ""))
			i                int = 0
		)

		for totalSpaceLength > 0 {
			wordsInLine[i] = wordsInLine[i] + strings.Repeat(" ", 1)
			totalSpaceLength -= 1
			i += 1
			if i == len(wordsInLine)-1 || (i == 1 && len(wordsInLine) == 1) {
				i = 0
			}
		}

		result = append(result, strings.Join(wordsInLine, ""))
	}

	return result
}
