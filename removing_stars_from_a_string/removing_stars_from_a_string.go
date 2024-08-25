package removing_stars_from_a_string

func removeStars(s string) string {
	r := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			r = append(r, s[i])
		} else {
			r = r[0 : len(r)-1]
		}
	}

	return string(r)
}
