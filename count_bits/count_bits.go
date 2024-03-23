package count_bits

func countBits(n int) []int {
	r := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		r = append(r, countOnes(i))
	}
	return r
}

func countOnes(n int) int {
	c := 0
	for n != 0 {
		c += n & 1
		n = n >> 1
	}
	return c
}
