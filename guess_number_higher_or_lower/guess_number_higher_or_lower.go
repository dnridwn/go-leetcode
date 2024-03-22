package guess_number_higher_or_lower

var pickN int

func guess(x int) int {
	if x > pickN {
		return -1
	} else if x < pickN {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	if n == 1 {
		return 1
	}

	max := n
	min := 1

	for {
		x := min + ((max - min) / 2)
		r := guess(x)

		if r == -1 {
			max = x - 1
		} else if r == 1 {
			min = x + 1
		} else {
			return x
		}
	}
}
