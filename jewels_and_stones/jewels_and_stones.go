package jewels_and_stones

func numJewelsInStones(jewels string, stones string) int {
	jewelsM := map[rune]bool{}
	r := 0

	for _, j := range jewels {
		jewelsM[j] = true
	}

	for _, s := range stones {
		if _, ok := jewelsM[s]; ok {
			r++
		}
	}

	return r
}
