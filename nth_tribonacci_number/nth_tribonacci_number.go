package nth_tribonacci_number

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	var (
		v1 = 0
		v2 = 0
		v3 = 1
	)

	for i := 2; i <= n; i++ {
		s := v3 + v2 + v1

		v1 = v2
		v2 = v3
		v3 = s
	}

	return v3
}
