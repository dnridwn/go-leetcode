package valid_anagram

import (
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss := []rune(s)
	tt := []rune(t)

	sort.Slice(ss, func(i, j int) bool {
		return ss[j] < ss[i]
	})

	sort.Slice(tt, func(i, j int) bool {
		return tt[j] < tt[i]
	})

	for i := 0; i < len(ss); i++ {
		if ss[i] != tt[i] {
			return false
		}
	}

	return true
}
