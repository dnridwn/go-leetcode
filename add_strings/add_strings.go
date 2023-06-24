package add_strings

import "math/big"

func addStrings(num1 string, num2 string) string {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString(num1, 10)
	b.SetString(num2, 10)

	return new(big.Int).Add(a, b).Text(10)
}
