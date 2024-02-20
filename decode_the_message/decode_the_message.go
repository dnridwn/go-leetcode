package decode_the_message

func removeDuplicatesAndSpaces(chars []rune) []rune {
	unique := make(map[rune]bool)
	result := []rune{}

	for _, c := range chars {
		if _, ok := unique[c]; !ok && c != rune(' ') {
			unique[c] = true
			result = append(result, c)
		}
	}

	return result
}

func mapCharsToKey(chars []rune) map[rune]int {
	result := make(map[rune]int)
	for k, c := range chars {
		result[c] = k
	}

	return result
}

func decodeMessage(key string, message string) string {
	keyMap := mapCharsToKey(removeDuplicatesAndSpaces([]rune(key)))
	alp := "abcdefghijklmnopqrstuvwxyz"
	result := ""

	for _, c := range message {
		if i, ok := keyMap[c]; ok {
			result += string(alp[i])
		} else if c == rune(' ') {
			result += string(c)
		}
	}

	return result
}
