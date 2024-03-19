package longest_palindrome_substring

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	max := 0
	sp := ""
	rs := reverse(s)

	for i := 0; i < len(s)-1; i++ {
		for j := len(s) - 1; j >= i; j-- {
			if (s[i:i+1] != s[j:j+1]) || (j-i) < max {
				if (j-i) < max && j == len(s)-1 {
					break
				}

				continue
			}

			f := s[i : j+1]
			r := rs[len(s)-j-1 : len(s)-i]

			if f == r {
				sp = f
				max = len(f)
			}
		}
	}

	return sp
}

func reverse(s string) string {
	rs := ""
	for i := 0; i < len(s); i++ {
		rs = s[i:i+1] + rs
	}
	return rs
}
