package smallest_even_multiple

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}

	return n * 2
}
