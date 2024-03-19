package integer_to_roman

func intToRoman(num int) string {
	r := ""
	sM := []struct {
		num    int
		symbol string
	}{
		{
			num:    1000,
			symbol: "M",
		},
		{
			num:    900,
			symbol: "CM",
		},
		{
			num:    500,
			symbol: "D",
		},
		{
			num:    400,
			symbol: "CD",
		},
		{
			num:    100,
			symbol: "C",
		},
		{
			num:    90,
			symbol: "XC",
		},
		{
			num:    50,
			symbol: "L",
		},
		{
			num:    40,
			symbol: "XL",
		},
		{
			num:    10,
			symbol: "X",
		},
		{
			num:    10,
			symbol: "X",
		},
		{
			num:    9,
			symbol: "IX",
		},
		{
			num:    5,
			symbol: "V",
		},
		{
			num:    4,
			symbol: "IV",
		},
		{
			num:    1,
			symbol: "I",
		},
	}

	g := num
	for _, m := range sM {
		x := m.num
		if float64(g/x) < 1 {
			continue
		}

		xx := g / x
		g %= x

		for i := 0; i < xx; i++ {
			r += m.symbol
		}
	}

	return r
}
