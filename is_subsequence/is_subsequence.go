package is_subsequence

func isSubsequence(s string, t string) bool {
	sb := []rune{}
	sbMap := map[rune]bool{}
	for _, x := range s {
		sb = append(sb, x)
		sbMap[x] = true
	}

	p := 0

	for _, x := range t {
		if p > len(s)-1 {
			break
		}

		if x == sb[p] {
			if _, ok := sbMap[x]; ok {
				p++
			}
		}
	}

	return p == len(s)
}
