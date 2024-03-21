package reverse_vowels_of_a_string

func reverseVowels(s string) string {
	vowelsPos := []int{}
	vowels := []rune{}

	for k, v := range s {
		if isVowel(v) {
			vowelsPos = append(vowelsPos, k)
			vowels = append([]rune{v}, vowels...)
		}
	}

	sc := []rune(s)
	i := 0
	for _, p := range vowelsPos {
		sc[p] = vowels[i]
		i++
	}

	return string(sc)
}

func isVowel(n rune) bool {
	return n == 65 || n == 69 || n == 73 || n == 79 || n == 85 || n == 97 || n == 101 || n == 105 || n == 111 || n == 117
}
