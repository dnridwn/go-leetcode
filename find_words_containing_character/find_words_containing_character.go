package find_words_containing_character

func findWordsContaining(words []string, x byte) []int {
	idx := []int{}

	for k, w := range words {
		for _, c := range []byte(w) {
			if c == x {
				idx = append(idx, k)
				break
			}
		}
	}

	return idx
}
